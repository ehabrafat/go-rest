-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS todos(
    id bigserial primary key,
    title varchar(500) not null,
    description text,
    createdAt timestamp not null,
    updatedAt timestamp not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS todos;
-- +goose StatementEnd
