package usersDelivery

import (
	"bdj-muhammadnurbasari/modules/auditTrails"
	"bdj-muhammadnurbasari/modules/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

type usersHandler struct {
	auditTrailsUsecase auditTrails.AuditTrailsUsecase
	usersuc            users.UsersUsecase
}

func NewUsersHandler(r *gin.Engine,
	auditTrailsUC auditTrails.AuditTrailsUsecase, usersuc users.UsersUsecase) {
	handler := usersHandler{
		auditTrailsUsecase: auditTrailsUC,
		usersuc:            usersuc,
	}

	users := r.Group("/users")
	users.Use(auditTrailsUC.MiddlewareAuditTrail)
	{
		users.GET("/data/all", handler.GetDataAll)
	}
}

func (handler *usersHandler) GetDataAll(c *gin.Context) {
	res, err := handler.usersuc.GetAllUsers()

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

	c.JSON(
		http.StatusOK,
		gin.H{
			"status": "success",
			"data":   &res,
		},
	)
}
