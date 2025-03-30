CREATE TABLE IF NOT EXISTS mentors (
    id SERIAL PRIMARY KEY,
    mentor_email TEXT UNIQUE NOT NULL,
    contact TEXT,
    count_reviews INTEGER NOT NULL DEFAULT 0,
    sum_rating FLOAT NOT NULL DEFAULT 0,
    average_rating FLOAT GENERATED ALWAYS AS (sum_rating::FLOAT / NULLIF(count_reviews, 0)) STORED
);
