CREATE TABLE IF NOT EXISTS users 
(
id SERIAL PRIMARY KEY, 
user_name VARCHAR[255], 
created_time  TIMESTAMP NOT NULL, 
begin_of_the_course  TIMESTAMP NOT NULL,
end_of_the_course  TIMESTAMP NOT NULL, 
end_time TIMESTAMP NOT NULL 
);