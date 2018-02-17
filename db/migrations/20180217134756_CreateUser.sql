
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users (
  id integer primary key,
  name varchar(255),
  age integer
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
