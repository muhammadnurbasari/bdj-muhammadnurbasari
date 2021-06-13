package hitApiKelurahanRepository

import (
	"bdj-muhammadnurbasari/models/apiKelurahanModel"
	"bdj-muhammadnurbasari/modules/hitApiKelurahan"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/gorm"
)

type apiKelurahanRepo struct {
	Url  string
	Conn *gorm.DB
}

//NewApiKelurahanRepository - will create an object that represent that hitApiRS.HitApiRSRepository interface
func NewApiKelurahanRepository(Url string, Conn *gorm.DB) hitApiKelurahan.HitApiKelurahanRepository {
	return &apiKelurahanRepo{Url: Url, Conn: Conn}
}

// GetDataKelurahanFromAPI - hit api from http://api.jakarta.go.id/v1/kelurahan
func (data *apiKelurahanRepo) GetDataKelurahanFromAPI() (*apiKelurahanModel.ResponseGetKelurahan, error) {
	url := data.Url
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, errors.New("GetDataKelurahanFromAPI Err http.NewRequest() = " + err.Error())
	}
	req.Header.Add("Authorization", "LdT23Q9rv8g9bVf8v/fQYsyIcuD14svaYL6Bi8f9uGhLBVlHA3ybTFjjqe+cQO8k")
	req.Header.Add("Cookie", "TS010f22dc=0181e0c80a3a50803a5a695aea22aa313885dba6a4fefc48e5cf46fb4685db6999ffd23feeb6fbb0b37634292b662aa9b4f5529b90; TS518ba6e9_27=086332a3cdab2000865d54c5279d81bd4c525228c021f141cf4d9d6f022925e27ea5edba665123ab083bdbd136112000689887718d1bf19de2e64654f55e977cdb941251c166641d7225367fceaa7218")

	res, err := client.Do(req)
	if err != nil {
		return nil, errors.New("GetDataKelurahanFromAPI Err client.Do() = " + err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("GetDataKelurahanFromAPI Err ioutil.ReadAll() = " + err.Error())
	}

	var resData apiKelurahanModel.ResponseGetKelurahan
	err = json.Unmarshal(body, &resData)
	if err != nil {
		var ErrorData apiKelurahanModel.ErrorResponseGetApi
		err = json.Unmarshal(body, &ErrorData)
		if err != nil {
			return nil, errors.New("GetDataKelurahanFromAPI Err json.Unmarshal() ErrorData = " + err.Error())
		}
		return nil, errors.New(ErrorData.Data)
	}

	return &resData, nil
}

func (data *apiKelurahanRepo) InsertDataLocationInfo() error {
	var kelurahan []apiKelurahanModel.InfoLocations

	errGet := data.Conn.Find(&kelurahan).Error
	if errGet != nil {
		return errGet
	}

	if len(kelurahan) == 0 {
		// get data rs
		dataKel, err := data.GetDataKelurahanFromAPI()
		if err != nil {
			return err
		}
		// var rsData apiRSModel.DataRumahSakits
		// var locationData apiRSModel.DataLocations
		// var faximileData apiRSModel.DataFaximiles
		// var teleponData apiRSModel.DataTelepons
		for _, v := range dataKel.Data {
			// insert rumah_sakits
			kelData := apiKelurahanModel.InfoLocations{
				KodeProvinsi:  v.KodeProvinsi,
				NamaProvinsi:  v.NamaProvinsi,
				KodeKota:      v.KodeKota,
				NamaKota:      v.NamaKota,
				KodeKecamatan: v.KodeKecamatan,
				NamaKecamatan: v.NamaKecamatan,
				KodeKelurahan: v.KodeKelurahan,
				NamaKelurahan: v.NamaKelurahan,
			}

			err := data.Conn.Create(&kelData).Error

			if err != nil {
				return err
			}

		}
	}

	return nil
}
