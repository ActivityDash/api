
-- +migrate Up
CREATE TABLE users (
  id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
  email character varying NOT NULL UNIQUE
);

-- +migrate Down
DROP TABLE users;
