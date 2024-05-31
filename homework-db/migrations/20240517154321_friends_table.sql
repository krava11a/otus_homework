-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS friends (
 id UUID NOT NULL,
 id_friend UUID NOT NULL,
 PRIMARY KEY (id,id_friend)
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE friends;
-- +goose StatementEnd
