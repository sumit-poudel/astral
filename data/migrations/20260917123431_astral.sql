-- +goose up
CREATE TABLE users (
    id int NOT NULL PRIMARY KEY,
    name text
);


INSERT INTO "users" ("id", "name") VALUES (1, 'gallant_almeida7');
INSERT INTO "users" ("id", "name") VALUES (2, 'brave_spence8');
INSERT INTO "users" ("id", "name") VALUES (99999, 'jovial_chaum1');
INSERT INTO "users" ("id", "name") VALUES (100000, 'goofy_ptolemy0');

-- +goose down
DROP TABLE users;
