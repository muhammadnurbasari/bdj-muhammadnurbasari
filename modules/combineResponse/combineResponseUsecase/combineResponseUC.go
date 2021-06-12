package combineResponseUsecase

import (
	"bdj-muhammadnurbasari/models/combineModel"
	"bdj-muhammadnurbasari/modules/combineResponse"
)

type combineUsecase struct {
	combineRepo combineResponse.CombineResponseRepository
}

//NewCombineUsecase - will create new an combineUsecase object representation of combineResponse.CombineResponseUsecase interface
func NewCombineUsecase(combineRepo combineResponse.CombineResponseRepository) combineResponse.CombineResponseUsecase {
	return &combineUsecase{
		combineRepo: combineRepo,
	}
}

//
func (combineUC *combineUsecase) GetResponseCombine() (*[]combineModel.ResultData, int, error) {
	// get response from repository
	resp, count, err := combineUC.combineRepo.GetResponseCombine()
	if err != nil {
		return nil, count, err
	}

	return resp, count, nil
}
