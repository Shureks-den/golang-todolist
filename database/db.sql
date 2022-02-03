DROP TABLE IF EXISTS tasks CASCADE;

CREATE TABLE tasks (
    id              SERIAL PRIMARY KEY,
    title           TEXT NOT NULL,
    description     TEXT DEFAULT "",
    isFinished      BOOLEAN DEFAULT false,
    created         TIMESTAMP WITH TIME ZONE DEFAULT now(),
    finished        TIMESTAMP WITH TIME ZONE DEFAULT NULL,
);