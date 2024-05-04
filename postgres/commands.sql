-- \? help
-- \l list dbs
-- \c {db name} connect to db by name

-- CREATE DATABASE {db name};

-- CREATE TABLE - www.postgresql.org/docs/current/sql-createtable.html 

-- \d list tables
-- \d {table name} show table
-- ALTER TABLEL -  www.postgresql.org/docs/current/sql-altertable.html

CREATE DATABASE edu_hit_zone; 
-- DROP DATABASE edu_hit_zone;
-- \c edu_hit_zone

CREATE TABLE course (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    thumbnail TEXT NOT NULL,
    page TEXT NOT NULL
);
-- MUST USE '' NOT ""
INSERT INTO course(name, thumbnail, page) 
VALUES (
    'Infi 2', 
    'https://upload.wikimedia.org/wikipedia/commons/thumb/9/9f/Integral_example.svg/300px-Integral_example.svg.png',
    'https://he.wikipedia.org/wiki/%D7%90%D7%99%D7%A0%D7%98%D7%92%D7%A8%D7%9C'
);

INSERT INTO course(name, thumbnail, page) 
VALUES (
    'C lang', 
    'https://upload.wikimedia.org/wikipedia/commons/thumb/1/18/C_Programming_Language.svg/380px-C_Programming_Language.svg.png?20201031132917',
    'https://he.wikipedia.org/wiki/C_(%D7%A9%D7%A4%D7%AA_%D7%AA%D7%9B%D7%A0%D7%95%D7%AA)'
);

SELECT * FROM course;