CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    name VARCHAR(70) NOT NULL,
    short_description VARCHAR(255) NOT NULL,
    classificacao_etaria VARCHAR(50) NOT NULL
);