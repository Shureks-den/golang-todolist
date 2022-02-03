DROP TABLE IF EXISTS tasks CASCADE;

CREATE TABLE tasks (
    id              SERIAL PRIMARY KEY,
    title           TEXT NOT NULL UNIQUE,
    description     TEXT DEFAULT '',
    is_finished      BOOLEAN DEFAULT false,
    created         TIMESTAMP WITH TIME ZONE DEFAULT now()
);