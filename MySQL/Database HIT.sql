ALTER DATABASE hit READ ONLY = 0;
DROP DATABASE hit;
create table students (
students_id INT,
first_name VARCHAR(50),
last_name VARCHAR(50),
email VARCHAR(100),
academic_year INT,
date_of_birth DATE
);

ALTER TABLE students
ADD phone_number VARCHAR(15);

INSERT INTO students
VALUES (1,"OFEK","BITTON","ofek97biton@gmail.com",2,"1997-12-10",0525444803),
       (2,"ANTON","AGOROV","antonagorov@gmail.com",2,"2002-05-05",0587542003);
SELECT * FROM students;
delete from students;






