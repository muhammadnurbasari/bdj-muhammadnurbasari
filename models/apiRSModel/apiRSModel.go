package apiRSModel

import "github.com/jinzhu/gorm"

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

// mapping table rumah_sakits
type (
	// RumahSakits - object table rumah_sakits
	RumahSakits struct {
		gorm.Model
		RsId          int     `gorm:"column:rs_id"`
		NamaRsu       string  `gorm:"column:nama_rsu"`
		JenisRsu      string  `gorm:"column:jenis_rsu"`
		KodePos       int     `gorm:"column:kode_pos"`
		Website       string  `gorm:"column:website"`
		Email         string  `gorm:"column:email"`
		KodeKota      int     `gorm:"column:kode_kota"`
		KodeKecamatan int     `gorm:"column:kode_kecamatan"`
		KodeKelurahan int     `gorm:"column:kode_kelurahan"`
		Latitude      float64 `gorm:"column:latitude"`
		Longitude     float64 `gorm:"column:longitude"`
	}

	// DataRumahSakits - parameter to insert data rumah_sakits
	DataRumahSakits struct {
		RsId          int
		NamaRsu       string
		JenisRsu      string
		KodePos       int
		Website       string
		Email         string
		KodeKota      int
		KodeKecamatan int
		KodeKelurahan int
		Latitude      float64
		Longitude     float64
	}
)

// mapping table locations
type (
	// Locations - object table locations
	Locations struct {
		gorm.Model
		Alamat    string  `gorm:"column:alamat"`
		Latitude  float64 `gorm:"column:latitude"`
		Longitude float64 `gorm:"column:longitude"`
		RsId      int     `gorm:"column:rs_id"`
	}

	// DataLocations - parameter to insert data locations
	DataLocations struct {
		Alamat    string
		Latitude  float64
		Longitude float64
		RsId      int
	}
)

// mapping table Faximiles
type (
	// Faximiles - object table Faximiles
	Faximiles struct {
		gorm.Model
		NoFax string `gorm:"column:no_fax"`
		RsId  int    `gorm:"column:rs_id"`
	}

	// DataFaximiles - parameter to insert data Faximiles
	DataFaximiles struct {
		NoFax string
		RsId  int
	}
)

// mapping table Telepons
type (
	// Telepons - object table Telepons
	Telepons struct {
		gorm.Model
		NoTelp string `gorm:"column:no_telp"`
		RsId   int    `gorm:"column:rs_id"`
	}

	// DataTelepons - parameter to insert data Telepons
	DataTelepons struct {
		NoTelp string
		RsId   int
	}
)
