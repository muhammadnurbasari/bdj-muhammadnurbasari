package hitApiRSRepository

import (
	"bdj-muhammadnurbasari/models/apiRSModel"
	"bdj-muhammadnurbasari/modules/hitApiRS"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type apiRSRepo struct {
	Url string
}

//NewApiRSRepository - will create an object that represent that hitApiRS.HitApiRSRepository interface
func NewApiRSRepository(Url string) hitApiRS.HitApiRSRepository {
	return &apiRSRepo{Url}
}

// GetDataRSFromAPI - hit api from http://api.jakarta.go.id/v1/rumahsakitumum
func (data *apiRSRepo) GetDataRSFromAPI() (*apiRSModel.ResponseGetRS, error) {
	url := data.Url
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, errors.New("GetDataRSFromAPI Err http.NewRequest() = " + err.Error())
	}
	req.Header.Add("Authorization", "LdT23Q9rv8g9bVf8v/fQYsyIcuD14svaYL6Bi8f9uGhLBVlHA3ybTFjjqe+cQO8k")
	req.Header.Add("Cookie", "TS010f22dc=0181e0c80a296432b00ea02ef76beb8095a2dd9604309aa16328e4e2dcff6ae2c9afa403f5289d3a682e79102bb4761dbf0cebeef1; TS518ba6e9_27=086332a3cdab2000ab9bd4bc0f309096ddb3de8aecebe7fb8c18fd75e602872bedf8a253ad84767808b00b15b511200035c5b8ef97e9f023a194bb022df306e53721ccea83a38aa35217bbf3d3a96e92")

	res, err := client.Do(req)
	if err != nil {
		return nil, errors.New("GetDataRSFromAPI Err client.Do() = " + err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("GetDataRSFromAPI Err ioutil.ReadAll() = " + err.Error())
	}

	var resData apiRSModel.ResponseGetRS
	err = json.Unmarshal(body, &resData)
	if err != nil {
		return nil, errors.New("GetDataRSFromAPI Err json.Unmarshal() = " + err.Error())
	}

	return &resData, nil
}
