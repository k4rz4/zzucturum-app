BEGIN;
CREATE TABLE IF NOT EXISTS counter_data (
    id SERIAL PRIMARY KEY,
    domain_name varchar(250) NOT NULL,
    number_requests INT NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE IF not EXISTS social_media (
    id SERIAL primary key,
    url varchar(250) not NULL,
    name varchar(250) int NOT null,
    is_active BOOLEAN
);

CREATE TABLE IF not EXISTS jobs (
    id SERIAL primary key,
    task_name varchar(250) not null,
    started_at timestamptz NOT NULL DEFAULT now(),
    finished_at timestamptz NOT NULL DEFAULT now()
)
COMMIT;

