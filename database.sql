CREATE EXTENSION IF NOT EXISTS "uuid-ossp";      -- enable uuid extension

CREATE TABLE users (
    user_id CHAR(36) PRIMARY KEY,                -- UUID v7 as CHAR(36)
    username VARCHAR(50) NOT NULL,               -- Username, max 50 characters
    email VARCHAR(320) NOT NULL UNIQUE,          -- Email, unique with a maximum of 320 characters
    password VARCHAR(255) NOT NULL,              -- Password, max 255 characters
    phone_number VARCHAR(20) NOT NULL UNIQUE,    -- Phone number, unique with max 20 characters
    picture_url VARCHAR(255),                    -- Nullable picture URL, max 255 characters
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Auto-generated creation time
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Auto-updated modification time
    deleted_at TIMESTAMP,                        -- Soft deletion timestamp
    UNIQUE (email),                              -- Index for unique constraint on email
    UNIQUE (phone_number)                        -- Index for unique constraint on phone number
);
