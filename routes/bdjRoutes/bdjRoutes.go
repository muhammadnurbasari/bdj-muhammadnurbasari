package bdjRoutes

import (
	"bdj-muhammadnurbasari/modules/auditTrails/auditTrailsRepository"
	"bdj-muhammadnurbasari/modules/auditTrails/auditTrailsUsecase"
	"bdj-muhammadnurbasari/modules/combineResponse/combineResponseHandler"
	"bdj-muhammadnurbasari/modules/combineResponse/combineResponseRepository"
	"bdj-muhammadnurbasari/modules/combineResponse/combineResponseUsecase"
	"bdj-muhammadnurbasari/modules/hitApiKelurahan/hitApiKelurahanRepository"
	"bdj-muhammadnurbasari/modules/hitApiRS/hitApiRSRepository"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// BdjRoutes - list of welcome route
func BdjRoutes(r *gin.Engine, db *gorm.DB) {
	auditTrailRepo := auditTrailsRepository.NewAuditTrailsRepository(db)
	auditTrailUC := auditTrailsUsecase.NewAuditTrailsUsecase(auditTrailRepo)

	apiKelurahanRepo := hitApiKelurahanRepository.NewApiKelurahanRepository("http://api.jakarta.go.id/v1/kelurahan")
	apiRSRepo := hitApiRSRepository.NewApiRSRepository("http://api.jakarta.go.id/v1/rumahsakitumum")

	combineRepo := combineResponseRepository.NewCombineRepository(apiKelurahanRepo, apiRSRepo)
	combineUC := combineResponseUsecase.NewCombineUsecase(combineRepo)
	combineResponseHandler.NewCombineHTTPHandler(r, auditTrailUC, combineUC)
}
