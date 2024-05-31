-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS debug_posts (
  text text NOT NULL 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE debug_posts;
-- +goose StatementEnd