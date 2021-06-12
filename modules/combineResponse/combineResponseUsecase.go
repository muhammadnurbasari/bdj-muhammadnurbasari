package combineResponse

import "bdj-muhammadnurbasari/models/combineModel"

type CombineResponseUsecase interface {
	GetResponseCombine() (*[]combineModel.ResultData, int, error)
}
