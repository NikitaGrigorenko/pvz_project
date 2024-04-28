-- +goose Up
-- +goose StatementBegin
create table pvz(
    id BIGSERIAL PRIMARY KEY NOT NULL ,
    name TEXT NOT NULL DEFAULT '',
    address TEXT NOT NULL DEFAULT '',
    contact TEXT NOT NULL DEFAULT ''
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table pvz;
-- +goose StatementEnd
