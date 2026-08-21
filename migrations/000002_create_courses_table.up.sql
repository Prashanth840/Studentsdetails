CREATE TABLE IF NOT EXISTS courses (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    student_id INTEGER REFERENCES students (id)
);

CREATE INDEX IF NOT EXISTS idx_courses_student_id ON courses (student_id);
