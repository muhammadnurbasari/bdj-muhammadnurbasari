package combineResponseHandler

import (
	"bdj-muhammadnurbasari/helpers/constanta"
	"bdj-muhammadnurbasari/helpers/validation"
	"bdj-muhammadnurbasari/modules/auditTrails"
	"bdj-muhammadnurbasari/modules/combineResponse"
	"net/http"

	"github.com/gin-gonic/gin"
)

type combineHandler struct {
	auditTrailsUsecase auditTrails.AuditTrailsUsecase
	combineUC          combineResponse.CombineResponseUsecase
}

func NewCombineHTTPHandler(r *gin.Engine,
	auditTrailsUC auditTrails.AuditTrailsUsecase, combineUC combineResponse.CombineResponseUsecase) {
	handler := combineHandler{
		auditTrailsUsecase: auditTrailsUC,
		combineUC:          combineUC,
	}

	bdj := r.Group("/bdj")
	bdj.Use(auditTrailsUC.MiddlewareAuditTrail)
	{
		bdj.GET("/data/combine", handler.GetDataCombine)
	}
}

// GetDataCombine godoc
// @Summary get data Combine RS DATA & Kelurahan DATA
// @Description get all data Combine RS DATA & Kelurahan DATA
// @Tags Combine RS DATA & Kelurahan DATA
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization"
// @Success 200 {object} combineModel.Response
// @security JWTToken
// @Router /bdj/data/combine [get]
func (handler *combineHandler) GetDataCombine(c *gin.Context) {
	Authorization := c.Request.Header.Get("Authorization")

	if Authorization == "" {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":   http.StatusBadRequest,
				"messages": "PLEASE, Authorization IN HEADER CANT EMPTY",
			},
		)
		return
	}

	// CHECK Authorization
	isTrue := validation.CheckAuthHash(Authorization, constanta.AuthorizationHashing)

	if isTrue == false {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":   http.StatusBadRequest,
				"messages": "SORY, YOU DONT HAVE ACCESS",
			},
		)
		return
	}

	res, count, err := handler.combineUC.GetResponseCombine()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":   http.StatusInternalServerError,
				"messages": err.Error(),
			},
		)
		return
	}

	if count == 0 {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": "success",
				"count":  count,
				"data":   &[]string{},
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status": "success",
			"count":  count,
			"data":   &res,
		},
	)
}
