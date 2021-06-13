CREATE TABLE IF NOT EXISTS info_locations(
    id BIGINT NOT NULL AUTO_INCREMENT,
    kode_provinsi BIGINT NOT NULL,
    nama_provinsi varchar(100) not null,
    kode_kota BIGINT NOT NULL,
    nama_kota varchar(100) not null,
    kode_kecamatan BIGINT NOT NULL,
    nama_kecamatan varchar(100) not null,
    kode_kelurahan BIGINT NOT NULL,
    nama_kelurahan varchar(100) not null,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id)
);