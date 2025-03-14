package token

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidSigningMethod = errors.New("invalid signing method")
	ErrTokenExpired         = errors.New("token expired")
	ErrMissingUserID        = errors.New("missing user_id")
)

type Claims struct {
	UserID    int64  `json:"user_id"`
	Role      string `json:"role"`
	TokenType string `token:"token_type"`
	ExpiresAt int64  `json:"exp"`
}

type TokenManager struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewTokenmanagerRSA(privateKeyPath, publicKeyPath string) (*TokenManager, error) {
	// Читаем приватный ключ
	privData, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key from %s: %v", privateKeyPath, err)
	}
	// Парсим
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(privData)
	if err != nil {
		return nil, err
	}

	// Читаем публичный ключ
	pubData, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key from %s: %v", publicKeyPath, err)
	}

	// Парсим
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubData)
	if err != nil {
		return nil, err
	}

	return &TokenManager{
		privateKey: privKey,
		publicKey:  pubKey,
	}, nil
}

func (tm *TokenManager) GenerateToken(userID int64, role string, ttl time.Duration, tokenType string) (string, error) {
	now := time.Now().Unix()
	exp := now + int64(ttl.Seconds())
	claims := &Claims{
		UserID:    userID,
		Role:      role,
		TokenType: tokenType,
		ExpiresAt: exp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(tm.privateKey)
}

// Проверяем подпись публичным ключом
func (tm *TokenManager) ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, ErrInvalidSigningMethod
		}
		return tm.publicKey, nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		if claims.ExpiresAt.Before(time.Now()) {
			return nil, ErrTokenExpired
		}
		return claims, nil
	}
	return nil, fmt.Errorf("parse token: %w", err)
}
