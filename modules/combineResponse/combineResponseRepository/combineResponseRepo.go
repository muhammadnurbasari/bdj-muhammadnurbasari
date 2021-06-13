package combineResponseRepository

import (
	"bdj-muhammadnurbasari/models/apiKelurahanModel"
	"bdj-muhammadnurbasari/models/combineModel"
	"bdj-muhammadnurbasari/modules/combineResponse"
	"bdj-muhammadnurbasari/modules/hitApiKelurahan"
	"bdj-muhammadnurbasari/modules/hitApiRS"
	"errors"
	"strconv"

	"github.com/jinzhu/gorm"
)

type combineRepository struct {
	hitApiKelurahanRepo hitApiKelurahan.HitApiKelurahanRepository
	hitApiRSRepo        hitApiRS.HitApiRSRepository
	Conn                *gorm.DB
}

//NewCombineRepository - will create an object that represent that combineResponse.CombineResponseRepository interface
func NewCombineRepository(hitApiKelurahanRepo hitApiKelurahan.HitApiKelurahanRepository,
	hitApiRSRepo hitApiRS.HitApiRSRepository, Conn *gorm.DB) combineResponse.CombineResponseRepository {
	return &combineRepository{
		hitApiKelurahanRepo: hitApiKelurahanRepo,
		hitApiRSRepo:        hitApiRSRepo,
		Conn:                Conn,
	}
}

// GetResponseCombine - get response combine
func (combineRepo *combineRepository) GetResponseCombine() (*[]combineModel.ResultData, int, error) {
	countData := 0

	// get resp api RS
	dataRS, errDataRS := combineRepo.hitApiRSRepo.GetDataRSFromAPI()

	if errDataRS != nil {
		return nil, countData, errors.New("GetResponseCombine() errDataRS = " + errDataRS.Error())
	}

	if dataRS.Status != "success" {
		return nil, countData, errors.New("GetResponseCombine() status get api RS = " + dataRS.Status)
	}

	countData = len(dataRS.Data)

	var response []combineModel.ResultData
	for _, v := range dataRS.Data {
		// location
		resLocation := combineModel.Location{
			Latitude:  v.Location.Latitude,
			Longitude: v.Location.Longitude,
		}

		// kelurahan
		kodeKelurahanString := strconv.Itoa(v.KodeKelurahan)
		namaKelurahan, namaKec, namaKota, errNamaKelurahan := combineRepo.GetNamaByKode(v.KodeKelurahan, v.KodeKecamatan, v.KodeKota)
		if errNamaKelurahan != nil {
			return nil, countData, errors.New(errNamaKelurahan.Error())
		}
		resKelurahan := combineModel.Kelurahan{
			Kode: kodeKelurahanString,
			Nama: namaKelurahan,
		}

		// kecamatan
		kodeKecamatanString := strconv.Itoa(v.KodeKecamatan)
		resKec := combineModel.Kecamatan{
			Kode: kodeKecamatanString,
			Nama: namaKec,
		}

		// kota
		kodeKotaString := strconv.Itoa(v.KodeKota)
		resKota := combineModel.Kota{
			Kode: kodeKotaString,
			Nama: namaKota,
		}

		result := combineModel.ResultData{
			ID:        v.ID,
			NamaRsu:   v.NamaRsu,
			JenisRsu:  v.JenisRsu,
			Location:  resLocation,
			Alamat:    v.Location.Alamat,
			KodePos:   v.KodePos,
			Telepon:   v.Telepon,
			Faximile:  v.Faximile,
			Website:   v.Website,
			Email:     v.Email,
			Kelurahan: resKelurahan,
			Kecamatan: resKec,
			Kota:      resKota,
		}

		response = append(response, result)
	}

	return &response, countData, nil

}

// GetNamaByKode
func (combineRepo *combineRepository) GetNamaByKode(kodeKel, kodeKec, kodeKota int) (string, string, string, error) {
	var data apiKelurahanModel.InfoLocations
	errKelurahan := combineRepo.Conn.Select("*").Where("kode_kelurahan = ? AND kode_kecamatan = ? AND kode_kota = ?", kodeKel, kodeKec, kodeKota).First(&data).Error
	if errKelurahan != nil {
		if errKelurahan.Error() == "record not found" {
			return "", "", "", errKelurahan
		}
		return "", "", "", errors.New("FindByID err = " + errKelurahan.Error())
	}
	var namaKel, namaKec, namaKota string
	namaKel = data.NamaKelurahan
	namaKec = data.NamaKecamatan
	namaKota = data.NamaKota

	return namaKel, namaKec, namaKota, nil
}
