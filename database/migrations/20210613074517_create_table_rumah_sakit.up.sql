CREATE TABLE IF NOT EXISTS rumah_sakits(
    id BIGINT NOT NULL AUTO_INCREMENT,
    rs_id int NOT NULL,
    nama_rsu VARCHAR(100) NOT NULL,
    jenis_rsu VARCHAR(255) NOT NULL,
    kode_pos int NOT NULL,
    website VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    kode_kota int NOT NULL,
    kode_kecamatan int NOT NULL,
    kode_kelurahan int NOT NULL,
    latitude FLOAT NOT NULL,
    longitude FLOAT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id)
);