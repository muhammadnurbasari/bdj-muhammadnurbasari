package registry

import (
	"bdj-muhammadnurbasari/modules/hitApiKelurahan/hitApiKelurahanRepository"
	"bdj-muhammadnurbasari/modules/hitApiRS/hitApiRSRepository"
)

func (reg *AppRegistry) insertData() error {
	apiRSRepo := hitApiRSRepository.NewApiRSRepository("http://api.jakarta.go.id/v1/rumahsakitumum", reg.ConnPublic)
	err := apiRSRepo.InsertDataRS()
	if err != nil {
		return err
	}

	apiKelurahanRepo := hitApiKelurahanRepository.NewApiKelurahanRepository("http://api.jakarta.go.id/v1/kelurahan", reg.ConnPublic)
	err = apiKelurahanRepo.InsertDataLocationInfo()
	if err != nil {
		return err
	}
	return nil
}
