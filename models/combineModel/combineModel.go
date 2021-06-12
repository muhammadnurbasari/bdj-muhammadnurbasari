package combineModel

// ResultData - object for response data
type ResultData struct {
	ID        int       `json:"id"`
	NamaRsu   string    `json:"nama_rsu"`
	JenisRsu  string    `json:"jenis_rsu"`
	Location  Location  `json:"location"`
	Alamat    string    `json:"alamat"`
	KodePos   int       `json:"kode_pos"`
	Telepon   []string  `json:"telepon"`
	Faximile  []string  `json:"faximile"`
	Website   string    `json:"website"`
	Email     string    `json:"email"`
	Kelurahan Kelurahan `json:"kelurahan"`
	Kecamatan Kecamatan `json:"kecamatan"`
	Kota      Kota      `json:"kota"`
}

// Location - object manage get Location from data
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Kelurahan - object manage get Kelurahan from data
type Kelurahan struct {
	Kode string `json:"kode"`
	Nama string `json:"nama"`
}

// Kecamatan - object manage get Kecamatan from data
type Kecamatan struct {
	Kode string `json:"kode"`
	Nama string `json:"nama"`
}

// Kota - object manage get Kota from data
type Kota struct {
	Kode string `json:"kode"`
	Nama string `json:"nama"`
}

type Response struct {
	Status string       `json:"status" example:"success"`
	Count  int          `json:"count" example:"5"`
	Data   []ResultData `json:"data"`
}
