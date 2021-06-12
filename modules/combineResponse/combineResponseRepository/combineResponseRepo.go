package combineResponseRepository

import (
	"bdj-muhammadnurbasari/models/combineModel"
	"bdj-muhammadnurbasari/modules/combineResponse"
	"bdj-muhammadnurbasari/modules/hitApiKelurahan"
	"bdj-muhammadnurbasari/modules/hitApiRS"
	"errors"
	"strconv"
)

type combineRepository struct {
	hitApiKelurahanRepo hitApiKelurahan.HitApiKelurahanRepository
	hitApiRSRepo        hitApiRS.HitApiRSRepository
}

//NewCombineRepository - will create an object that represent that combineResponse.CombineResponseRepository interface
func NewCombineRepository(hitApiKelurahanRepo hitApiKelurahan.HitApiKelurahanRepository,
	hitApiRSRepo hitApiRS.HitApiRSRepository) combineResponse.CombineResponseRepository {
	return &combineRepository{
		hitApiKelurahanRepo: hitApiKelurahanRepo,
		hitApiRSRepo:        hitApiRSRepo,
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
		namaKelurahan, errNamaKelurahan := combineRepo.GetNamaKelurahanByKode(v.KodeKelurahan)
		if errNamaKelurahan != nil {
			return nil, countData, errors.New("GetResponseCombine() err nama kelurahan = " + errNamaKelurahan.Error())
		}
		resKelurahan := combineModel.Kelurahan{
			Kode: kodeKelurahanString,
			Nama: namaKelurahan,
		}

		// kecamatan
		kodeKecamatanString := strconv.Itoa(v.KodeKecamatan)
		namaKecamatan, errNamaKecamatan := combineRepo.GetNamaKecamatanByKode(v.KodeKecamatan)
		if errNamaKecamatan != nil {
			return nil, countData, errors.New("GetResponseCombine() err nama kelurahan = " + errNamaKecamatan.Error())
		}
		resKec := combineModel.Kecamatan{
			Kode: kodeKecamatanString,
			Nama: namaKecamatan,
		}

		// kota
		kodeKotaString := strconv.Itoa(v.KodeKota)
		namaKota, errNamaKota := combineRepo.GetNamaKotaByKode(v.KodeKota)
		if errNamaKota != nil {
			return nil, countData, errors.New("GetResponseCombine() err nama kelurahan = " + errNamaKota.Error())
		}
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

// GetNamaKelurahanByKode
func (combineRepo *combineRepository) GetNamaKelurahanByKode(kode int) (string, error) {
	// get resp api Kelurahan
	dataKelurahan, errDataKelurahan := combineRepo.hitApiKelurahanRepo.GetDataKelurahanFromAPI()
	if errDataKelurahan != nil {
		return "", errDataKelurahan
	}
	var resp string
	for _, v := range dataKelurahan.Data {
		if kode == v.KodeKelurahan {
			resp = v.NamaKelurahan
			break
		}
	}

	return resp, nil
}

// GetNamaKecamatanByKode
func (combineRepo *combineRepository) GetNamaKecamatanByKode(kode int) (string, error) {
	// get resp api Kelurahan
	dataKec, errDataKec := combineRepo.hitApiKelurahanRepo.GetDataKelurahanFromAPI()
	if errDataKec != nil {
		return "", errDataKec
	}
	var resp string
	for _, v := range dataKec.Data {
		if kode == v.KodeKecamatan {
			resp = v.NamaKecamatan
			break
		}
	}

	return resp, nil
}

// GetNamaKotaByKode
func (combineRepo *combineRepository) GetNamaKotaByKode(kode int) (string, error) {
	// get resp api Kelurahan
	dataKota, errDataKota := combineRepo.hitApiKelurahanRepo.GetDataKelurahanFromAPI()
	if errDataKota != nil {
		return "", errDataKota
	}
	var resp string
	for _, v := range dataKota.Data {
		if kode == v.KodeKota {
			resp = v.NamaKota
			break
		}
	}

	return resp, nil
}
