CREATE TABLE agents (
    id INT PRIMARY KEY,
    _type INT NOT NULL, 
    name VARCHAR(255) NOT NULL,
    shortDescription TEXT ,
    longDescription TEXT,
    status INT NOT NULL, 
    parent INT, 
    _children INT[], 
    _spaces INT[], 
    _projects INT[], 
    _events INT[]
	-- FOREIGN KEY(parent) REFERENCES agents(id)
);


CREATE TABLE spaces (
    id INT PRIMARY KEY,
    location VARCHAR(255) NOT NULL,
    _geoLocation VARCHAR(255) NOT NULL, 
    name VARCHAR(255) NOT NULL,
    public BOOLEAN NOT NULL,
    shortDescription TEXT,
    longDescription TEXT,
    status INT NOT NULL,
    _type INT NOT NULL,
    eventOccurrences INT[], 
    parent INT,
    _children INT[], 
    owner INT
	-- FOREIGN KEY(parent) REFERENCES spaces(id),
	-- FOREIGN KEY (owner) REFERENCES agents(id)
);


CREATE TABLE projects (
    id INT PRIMARY KEY,
    _type INT  NOT NULL,
    name VARCHAR(255) NOT NULL,
    shortDescription TEXT,
    longDescription TEXT,
    updateTimestamp DATE, 
    registrationFrom DATE, 
    registrationTo DATE,
    createTimestamp DATE, 
    status INT NOT NULL, 
    parent INT, 
    _children INT[], 
    owner INT, 
    _events INT[]
	-- FOREIGN KEY(parent) REFERENCES projects(id),
	-- FOREIGN KEY(parent) REFERENCES agents(id)
);


CREATE TABLE events (
    id INT PRIMARY KEY,
    _type INT NOT NULL, 
    name VARCHAR(255) NOT NULL,
    shortDescription TEXT NOT NULL,
    longDescription TEXT,
    rules TEXT,
    createTimestamp DATE,
    status INT NOT NULL,
    occurrences INT[],
    owner INT, 
    project INT
	-- FOREIGN KEY(project) REFERENCES projects(id)
	-- FOREIGN KEY(owner) REFERENCES agents(id)
);



CREATE TABLE event_occurrences (
    id INT PRIMARY KEY,
    _startsOn DATE, 
    _endsOn DATE, 
    _startsAt DATE, 
    _endsAt DATE, 
    frequency VARCHAR(255), 
    separation INT NOT NULL ,
    _count INT,
    _until DATE,
    timezoneName VARCHAR(255) NOT NULL,
    event INT, 
    eventId INT NOT NULL,
    space INT, 
    spaceId INT NOT NULL, 
    rule TEXT NOT NULL,
    status INT NOT NULL
	-- FOREIGN KEY(event) REFERENCES events(id),
	-- FOREIGN KEY(space) REFERENCES spaces(id)
);