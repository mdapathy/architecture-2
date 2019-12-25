-- Drop tables if exist.--
DROP TABLE IF EXISTS "forums";
DROP TABLE IF EXISTS "users";
DROP TABLE IF EXISTS "userforum";

-- Create tables--
CREATE TABLE "forums"
(
    "id"   SERIAL PRIMARY KEY,
    "name" VARCHAR(100) NOT NULL UNIQUE,
    "topic_keyword" VARCHAR(50) NOT NULL
);

CREATE TABLE "users"
(
    "id"   SERIAL PRIMARY KEY,
    "users_name" VARCHAR(40) UNIQUE NOT NULL, 
    "users_interests" text[] NOT NULL

);

CREATE TABLE userforum (
    forumid integer REFERENCES forums(id),
    userid integer  REFERENCES users(id)
);

-- Insert demo data.
INSERT INTO users (users_name, users_interests) VALUES ('user1', '{"politics","music"}');
INSERT INTO users (users_name, users_interests) VALUES ('user2', '{"tv-series","music"}');
INSERT INTO users (users_name, users_interests) VALUES ('user3', '{"politics","tv-series"}');

INSERT INTO forums (name, topic_keyword) VALUES ('Politics in USA','politics');
INSERT INTO forums (name, topic_keyword) VALUES ('Game of Thrones S8','tv-series');
INSERT INTO forums (name, topic_keyword) VALUES ('Rock bands', 'music');

INSERT INTO userforum (forumid, userid) VALUES (1, 1);
INSERT INTO userforum (forumid, userid) VALUES (1, 3);
INSERT INTO userforum (forumid, userid) VALUES (2, 2);
INSERT INTO userforum (forumid, userid) VALUES (2, 3);
INSERT INTO userforum (forumid, userid) VALUES (3, 1);
INSERT INTO userforum (forumid, userid) VALUES (3, 2);
