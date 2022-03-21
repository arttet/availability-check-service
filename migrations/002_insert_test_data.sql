-- +goose Up

INSERT INTO checks
  (host, port, status, timeout, fail_message)
VALUES
  ('google.com', 443, 'fail', 1000 , 'no route to host'),
  ('12.12.12.12', 80, 'fail', 5000, 'timeout'),
  ('localhost', 22, 'ok', 100, '');

-- +goose Down

DELETE FROM checks;
