-- +goose Up
-- +goose StatementBegin
create table if not exists someusers(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    username VARCHAR (255) NOT NULL,
    email VARCHAR (255) NOT NULL,
    password VARCHAR (255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table someusers;
-- +goose StatementEnd
