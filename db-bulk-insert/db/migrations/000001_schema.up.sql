CREATE TABLE agents (
    id INT PRIMARY KEY,
    name TEXT NOT NULL,
    short_description TEXT,
    create_timestamp TIMESTAMP,
    update_timestamp TIMESTAMP,
    parent INT,
    terms TEXT,
    children INT[],
    spaces INT[],
    events INT[],
    projects INT[]
	-- FOREIGN KEY (parent) REFERENCES agents (id)
);

CREATE TABLE events (
    id INT PRIMARY KEY,
    name TEXT NOT NULL,
    short_description TEXT,
    create_timestamp TIMESTAMP,
    update_timestamp TIMESTAMP,
    classificacao_etaria TEXT,
    q TEXT,
    owner INT, 
    project INT
	-- FOREIGN KEY (project) REFERENCES projects (id)
	-- FOREIGN KEY (owner) REFERENCES agents (id)
);

CREATE TABLE spaces (
    id INT PRIMARY KEY,
    location TEXT NOT NULL,
    name TEXT NOT NULL,
    short_description TEXT,
    create_timestamp TIMESTAMP,
    update_timestamp TIMESTAMP,
    event_occurrences INT[], 
    horario TEXT,
    telefone TEXT,
    email TEXT,
    children INT[], 
    terms TEXT,
    parent INT,
    owner INT
	-- FOREIGN KEY (parent) REFERENCES spaces(id),
	-- FOREIGN KEY (owner) REFERENCES agents(id)
);

-----

CREATE TABLE projects (
    id INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    short_description TEXT,
    create_timestamp TIMESTAMP,
    update_timestamp TIMESTAMP,
    registration_from DATE, 
    registration_to DATE,
    parent INT, 
    children INT[], 
    owner INT, 
    events INT[]
	-- FOREIGN KEY (parent) REFERENCES projects (id),
	-- FOREIGN KEY (owner) REFERENCES agents (id)
);

CREATE TABLE event_occurrences (
    id INT PRIMARY KEY,
    starts_on DATE, 
    starts_at TIME, 
    ends_at TIME, 
    frequency TEXT, 
    separation INT,
    event INT, 
    space INT 
	-- FOREIGN KEY(event) REFERENCES events(id),
	-- FOREIGN KEY(space) REFERENCES spaces(id)
);