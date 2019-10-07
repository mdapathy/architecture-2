-- Drop tables if exist.--
DROP TABLE IF EXISTS "forums";
DROP TABLE IF EXISTS "users";

-- Create tables--
CREATE TABLE "forums"
(
    "id"   SERIAL PRIMARY KEY,
    "name" VARCHAR(100) NOT NULL UNIQUE,
    "topic_keyword" VARCHAR(50) NOT NULL,
    "users" int[]
);

CREATE TABLE "users"
(
    "id"   SERIAL PRIMARY KEY,
    "users_name" VARCHAR(40) UNIQUE NOT NULL, 
    "users_interests" text[] NOT NULL

);

-- Insert demo data.
INSERT INTO users (users_name, users_interests) VALUES ('user1', '{"politics","music"}');
INSERT INTO users (users_name, users_interests) VALUES ('user2', '{"tv-series","music"}');
INSERT INTO users (users_name, users_interests) VALUES ('user3', '{"politics","tv-series"}');

INSERT INTO forums (name, topic_keyword, users) VALUES ('Politics in USA','politics', '{1,3}');
INSERT INTO forums (name, topic_keyword, users) VALUES ('Game of Thrones S8','tv-series','{2,3}');
INSERT INTO forums (name, topic_keyword, users) VALUES ('Rock bands', 'music', '{1,2}');




