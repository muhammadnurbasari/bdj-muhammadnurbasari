package apiRSModel

// ResponseGetRS - object for get response RS
type ResponseGetRS struct {
	Status string       `json:"status"`
	Count  int          `json:"count"`
	Data   []ResultData `json:"data"`
}

// ResultData - object manage array data from response RS
type ResultData struct {
	ID            int      `json:"id"`
	NamaRsu       string   `json:"nama_rsu"`
	JenisRsu      string   `json:"jenis_rsu"`
	Location      Location `json:"location"`
	KodePos       int      `json:"kode_pos"`
	Telepon       []string `json:"telepon"`
	Faximile      []string `json:"faximile"`
	Website       string   `json:"website"`
	Email         string   `json:"email"`
	KodeKota      int      `json:"kode_kota"`
	KodeKecamatan int      `json:"kode_kecamatan"`
	KodeKelurahan int      `json:"kode_kelurahan"`
	Latitude      float64  `json:"latitude"`
	Longitude     float64  `json:"longitude"`
}

// Location - object manage get Location from data
type Location struct {
	Alamat    string  `json:"alamat"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
