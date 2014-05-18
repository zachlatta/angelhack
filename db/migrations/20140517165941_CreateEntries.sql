
-- +goose Up
CREATE TABLE entries (
  id serial not null primary key,
  user_id integer not null references users(id),
  created date not null,
  updated date not null,
  rating integer not null,
  message text not null
);


-- +goose Down
DROP TABLE entries;

