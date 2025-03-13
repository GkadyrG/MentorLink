package token

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
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

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"user_id":    userID,
		"role":       role,
		"token_type": tokenType,
		"exp":        exp,
	})

	// Подписываем приватный ключ
	return token.SignedString(tm.privateKey)
}

// Проверяем подпись публичным ключом
func (tm *TokenManager) ParseToken(tokenStr string) (*Claims, error) {
	parsed, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		// Проверяем, что алгоритм действительно RS256
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return tm.publicKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !parsed.Valid {
		return nil, errors.New("Invalid token")
	}

	mapClaims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	c := &Claims{}

	if uid, ok := mapClaims["user_id"].(float64); ok {
		c.UserID = int64(uid)
	} else {
		return nil, errors.New("missing user_id")
	}

	if role, ok := mapClaims["role"].(string); ok {
		c.Role = role
	} else {
		return nil, errors.New("missing role")
	}

	if tt, ok := mapClaims["token_type"].(string); ok {
		c.TokenType = tt
	} else {
		return nil, errors.New("missing token_type")
	}

	if exp, ok := mapClaims["exp"].(float64); ok {
		c.ExpiresAt = int64(exp)
	} else {
		return nil, errors.New("missing exp")
	}
	return c, nil
}
