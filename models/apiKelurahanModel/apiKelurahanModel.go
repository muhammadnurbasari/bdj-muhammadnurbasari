package apiKelurahanModel

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
