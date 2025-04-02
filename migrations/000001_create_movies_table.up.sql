CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS movies (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    director VARCHAR(255),
    year INT,
    plot TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO movies (title, director, year, plot)
VALUES
('Inception', 'Christopher Nolan', 2010, 'A thief who steals corporate secrets through the use of dream-sharing technology.'),
('The Matrix', 'Lana Wachowski, Lilly Wachowski', 1999, 'A computer hacker learns about the true nature of his reality.'),
('Interstellar', 'Christopher Nolan', 2014, 'A team of explorers travel through a wormhole in space in an attempt to ensure humanity''s survival.');