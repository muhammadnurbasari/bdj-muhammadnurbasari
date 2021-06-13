package apiKelurahanModel

import "github.com/jinzhu/gorm"

// ResponseGetKelurahan - object for get response Kelurahan
type ResponseGetKelurahan struct {
	Status string                `json:"status"`
	Count  int                   `json:"count"`
	Data   []ResultDataKelurahan `json:"data"`
}

// ResultDataKelurahan - object manage array data from response Kelurahan
type ResultDataKelurahan struct {
	KodeProvinsi  int    `json:"kode_provinsi"`
	NamaProvinsi  string `json:"nama_provinsi"`
	KodeKota      int    `json:"kode_kota"`
	NamaKota      string `json:"nama_kota"`
	KodeKecamatan int    `json:"kode_kecamatan"`
	NamaKecamatan string `json:"nama_kecamatan"`
	KodeKelurahan int    `json:"kode_kelurahan"`
	NamaKelurahan string `json:"nama_kelurahan"`
}

// ErrorResponseGetApi - object for unmarshal json if error
type ErrorResponseGetApi struct {
	Result bool   `json:"result"`
	Data   string `json:"data"`
}

// mapping table info_locations
type (
	// InfoLocations - object table info_locations
	InfoLocations struct {
		gorm.Model
		KodeProvinsi  int    `gorm:"column:kode_provinsi"`
		NamaProvinsi  string `gorm:"column:nama_provinsi"`
		KodeKota      int    `gorm:"column:kode_kota"`
		NamaKota      string `gorm:"column:nama_kota"`
		KodeKecamatan int    `gorm:"column:kode_kecamatan"`
		NamaKecamatan string `gorm:"column:nama_kecamatan"`
		KodeKelurahan int    `gorm:"column:kode_kelurahan"`
		NamaKelurahan string `gorm:"column:nama_kelurahan"`
	}
)
