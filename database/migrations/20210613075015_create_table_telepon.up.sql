CREATE TABLE IF NOT EXISTS telepons(
    id BIGINT NOT NULL AUTO_INCREMENT,
    no_telp VARCHAR(255) NOT NULL,
    rs_id int NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id)
);