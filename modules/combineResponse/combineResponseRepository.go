package combineResponse

import "bdj-muhammadnurbasari/models/combineModel"

type CombineResponseRepository interface {
	GetResponseCombine() (*[]combineModel.ResultData, int, error)
}
