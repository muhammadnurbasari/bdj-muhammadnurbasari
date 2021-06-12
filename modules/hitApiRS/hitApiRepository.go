package hitApiRS

import "bdj-muhammadnurbasari/models/apiRSModel"

type HitApiRSRepository interface {
	GetDataRSFromAPI() (*apiRSModel.ResponseGetRS, error)
}
