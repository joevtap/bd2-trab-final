CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    short_description VARCHAR(255) NOT NULL,
    classificacao_etaria INTEGER NOT NULL
);

-- Grants

GRANT ALL ON SCHEMA public TO application;
GRANT ALL ON ALL SEQUENCES IN SCHEMA public TO application;
GRANT ALL ON ALL TABLES IN SCHEMA public TO application;