CREATE TABLE IF NOT EXISTS projects (
    id BIGINT PRIMARY KEY,
    chain INT NOT NULL,
    owner VARCHAR(42) NOT NULL,
    detailsLocationType INT NOT NULL,
    detailsLocation TEXT NOT NULL,
    shortName VARCHAR(250) NOT NULL,
    tipJarAddress VARCHAR(42) NOT NULL,
    dbAdded TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now()
    -- TODO: add: created time in smart contract, isScam
);
CREATE INDEX IF NOT EXISTS projects_by_owner ON projects (owner);

CREATE TABLE IF NOT EXISTS updates (
    tx TEXT PRIMARY KEY,
    projectId BIGINT REFERENCES projects(id),
    chain INT NOT NULL,
    content TEXT NOT NULL,
    dbAdded TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS comments (
    tx TEXT PRIMARY KEY,
    projectId BIGINT REFERENCES projects(id),
    chain INT NOT NULL,
    author VARCHAR(42) NOT NULL,
    content TEXT NOT NULL
    -- created TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX IF NOT EXISTS comments_by_author ON comments (author);

CREATE TABLE IF NOT EXISTS proposals (
    id SERIAL PRIMARY KEY,
    projectId BIGINT REFERENCES projects(id),
    chain INT NOT NULL,
    address VARCHAR(42) NOT NULL,
    message VARCHAR(1000) NOT NULL,
    created TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX IF NOT EXISTS proposals_unique ON proposals (projectId, address);

CREATE TABLE IF NOT EXISTS team_members (
    tx TEXT PRIMARY KEY,
    projectId BIGINT REFERENCES projects(id),
    chain INT NOT NULL,
    address VARCHAR(42) NOT NULL,
    accepted BOOLEAN DEFAULT FALSE,
    memberLeft BOOLEAN DEFAULT FALSE
);
CREATE INDEX IF NOT EXISTS team_members_by_project ON team_members (projectId);
CREATE INDEX IF NOT EXISTS team_members_by_address ON team_members (address);

CREATE TABLE IF NOT EXISTS progress (
    chain INT PRIMARY KEY,
    lastBlock BIGINT NOT NULL
);