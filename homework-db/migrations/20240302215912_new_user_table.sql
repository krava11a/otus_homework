-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
 id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
 first_name VARCHAR(255) NOT NULL,
 second_name VARCHAR(255) NOT NULL,
 birthdate  DATE NOT NULL,
 biography VARCHAR(255) NOT NULL,
 city VARCHAR(200) NOT NULL,
 hP VARCHAR(255) NOT NULL
);
INSERT INTO users (
    first_name,
    second_name,
    birthdate,
    biography,
    city,
    hP)
    VALUES(
        'scot',
        'tiger',
        '1978-03-03',
        'go,java',
        'SP',
        'askdjflkajsdlfkja;lfd'

    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
