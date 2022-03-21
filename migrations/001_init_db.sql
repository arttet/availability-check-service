-- +goose Up

CREATE TABLE checks (
    id           INT                NOT NULL AUTO_INCREMENT,
    host         VARCHAR(2048)      NOT NULL,
    port         SMALLINT UNSIGNED  NOT NULL,
    status       ENUM ('ok','fail') NOT NULL,
    timeout      INT UNSIGNED       NOT NULL,
    fail_message TEXT               NOT NULL,

    CONSTRAINT checks_pk PRIMARY KEY (id)
);

-- +goose Down

DROP TABLE checks;
