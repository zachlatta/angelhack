
-- +goose Up
ALTER TABLE entries
  DROP COLUMN user_id;

CREATE TABLE journals (
  id serial not null primary key,
  user_id integer not null references users(id),
  created date not null,
  updated date not null,
  name text not null
);

ALTER TABLE entries
  ADD COLUMN journal_id integer not null references journals(id) default 0;

-- +goose Down
ALTER TABLE entries
  ADD COLUMN user_id integer not null references users(id) default 0;

DROP TABLE journals;

ALTER TABLE entries
  DROP COLUMN journal_id;
