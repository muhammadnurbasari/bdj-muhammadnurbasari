CREATE TABLE IF NOT EXISTS locations(
    id BIGINT NOT NULL AUTO_INCREMENT,
    alamat VARCHAR(255) NOT NULL,
    latitude FLOAT NOT NULL,
    longitude FLOAT NOT NULL,
    rs_id int NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id)
);