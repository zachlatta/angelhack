
-- +goose Up
ALTER TABLE users ALTER COLUMN created TYPE timestamp;
ALTER TABLE users ALTER COLUMN updated TYPE timestamp;
ALTER TABLE journals ALTER COLUMN created TYPE timestamp;
ALTER TABLE journals ALTER COLUMN updated TYPE timestamp;
ALTER TABLE entries ALTER COLUMN created TYPE timestamp;
ALTER TABLE entries ALTER COLUMN updated TYPE timestamp;


-- +goose Down
ALTER TABLE users ALTER COLUMN created TYPE date;
ALTER TABLE users ALTER COLUMN updated TYPE date;
ALTER TABLE journals ALTER COLUMN created TYPE date;
ALTER TABLE journals ALTER COLUMN updated TYPE date;
ALTER TABLE entries ALTER COLUMN created TYPE date;
ALTER TABLE entries ALTER COLUMN updated TYPE date;
