-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS posts (
 id_post UUID PRIMARY KEY DEFAULT gen_random_uuid(),
 id_user UUID,
 text text NOT NULL,
 timestamp BIGINT not NULL 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE posts;
-- +goose StatementEnd
