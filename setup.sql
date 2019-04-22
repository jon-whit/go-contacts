CREATE TABLE contacts (
    user_id NOT NULL INTEGER,
    id NOT NULL INTEGER,
    first_name TEXT,
    last_name TEXT,
    phone TEXT,
    email TEXT,
    PRIMARY KEY(user_id, id)
);

INSERT INTO contacts (user_id, id, first_name, last_name, phone, email) VALUES (1, 1, "James", "Butt", "504-621-8927", "jbutt@gmail.com");
INSERT INTO contacts (user_id, id, first_name, last_name, phone, email) VALUES (1, 2, "Kris", "Marrier", "810-292-9388", "kris@gmail.com");
INSERT INTO contacts (user_id, id, first_name, last_name, phone, email) VALUES (2, 1, "Willard", "Kolmetz", "972-303-9197", "willard@hotmail.com");