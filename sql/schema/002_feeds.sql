-- +goose Up
CREATE TABLE feeds(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP not null,
    updated_at TIMESTAMP not null,
    name text not null,
    url text UNIQUE not null,
    user_id UUID not null,
    foreign KEY (user_id) references users(id) on delete cascade
);

-- +goose Down
DROP TABLE feeds;