package hitApiKelurahan

import "bdj-muhammadnurbasari/models/apiKelurahanModel"

type HitApiKelurahanRepository interface {
	GetDataKelurahanFromAPI() (*apiKelurahanModel.ResponseGetKelurahan, error)
	InsertDataLocationInfo() error
}
