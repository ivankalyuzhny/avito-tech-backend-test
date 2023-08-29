CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS segments
(
    id   SERIAL PRIMARY KEY,
    slug TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS user_segments
(
    user_id    INTEGER NOT NULL,
    segment_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (segment_id) REFERENCES segments (id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, segment_id)
);

INSERT INTO users DEFAULT VALUES;
INSERT INTO users DEFAULT VALUES;
INSERT INTO users DEFAULT VALUES;
INSERT INTO users DEFAULT VALUES;
INSERT INTO users DEFAULT VALUES;
INSERT INTO users DEFAULT VALUES;
INSERT INTO users DEFAULT VALUES;
INSERT INTO users DEFAULT VALUES;
INSERT INTO users DEFAULT VALUES;
INSERT INTO users DEFAULT VALUES;