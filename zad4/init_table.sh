#!/usr/bin/env bash

DB_FILE="db.sqlite3"

if [ -f "$DB_FILE" ]; then
    echo "Database file already exists. Exiting..."
    exit 0
fi

# Create the SQLite database file
sqlite3 "$DB_FILE" <<EOF
-- Create the books table
CREATE TABLE IF NOT EXISTS books (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    year INTEGER NOT NULL,
    isbn TEXT NOT NULL UNIQUE,
    pages INTEGER NOT NULL
);

-- Create the authors table
CREATE TABLE IF NOT EXISTS authors (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE
);

-- Create the genres table
CREATE TABLE IF NOT EXISTS genres (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE
);

-- Create the book_genres table
CREATE TABLE IF NOT EXISTS book_genres (
    book_id INTEGER NOT NULL,
    genre_id INTEGER NOT NULL,
    PRIMARY KEY (book_id, genre_id),
    FOREIGN KEY (book_id) REFERENCES books (id),
    FOREIGN KEY (genre_id) REFERENCES genres (id)
);

-- Create the book_authors table
CREATE TABLE IF NOT EXISTS book_authors (
    book_id INTEGER NOT NULL,
    author_id INTEGER NOT NULL,
    PRIMARY KEY (book_id, author_id),
    FOREIGN KEY (book_id) REFERENCES books (id),
    FOREIGN KEY (author_id) REFERENCES authors (id)
);

-- Create the reviews table
CREATE TABLE IF NOT EXISTS reviews (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    book_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    rating INTEGER NOT NULL CHECK (rating >= 1 AND rating <= 5),
    review_text TEXT,
    FOREIGN KEY (book_id) REFERENCES books (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

-- Create the users table
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE
);
EOF

echo "Database and tables created successfully."

# Insert sample data
sqlite3 "$DB_FILE" <<EOF
-- Insert sample data into authors
INSERT INTO authors (name) VALUES
('J.K. Rowling'),
('George R.R. Martin'),
('J.R.R. Tolkien'),
('Friedrich Nietzsche');

-- Insert sample data into genres
INSERT INTO genres (name) VALUES
('Fantasy'),
('Science Fiction'),
('Mystery'),
('Drama');

-- Insert sample data into books
INSERT INTO books (title, year, isbn, pages) VALUES
('Harry Potter and the Philosophers Stone', 1997, '9780747532699', 223),
('A Game of Thrones', 1996, '9780553103540', 694),
('The Hobbit', 1937, '9780345339683', 310);

-- Insert sample data into book_genres
INSERT INTO book_genres (book_id, genre_id) VALUES
(1, 1),
(2, 1),
(3, 1),
(3, 2);

-- Insert sample data into book_authors
INSERT INTO book_authors (book_id, author_id) VALUES
(1, 1),
(2, 2),
(3, 3);

-- Insert sample data into users
INSERT INTO users (name, email) VALUES
('Alice', 'alice@example.com'),
('Bob', 'bob@example.com'),
('Charlie', 'charlie@example.com');

-- Insert sample data into reviews
INSERT INTO reviews (book_id, user_id, rating, review_text) VALUES
(1, 1, 5, 'An amazing start to a magical series!'),
(2, 2, 4, 'A gripping tale of power and betrayal.'),
(3, 3, 5, 'A timeless classic full of adventure.');
EOF

echo "Sample data inserted successfully."
