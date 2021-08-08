package usersRoutes

import (
	"bdj-muhammadnurbasari/modules/auditTrails/auditTrailsRepository"
	"bdj-muhammadnurbasari/modules/auditTrails/auditTrailsUsecase"
	"bdj-muhammadnurbasari/modules/users/usersDelivery"
	"bdj-muhammadnurbasari/modules/users/usersRepository"
	"bdj-muhammadnurbasari/modules/users/usersUsecase"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// UsersRoutes - list of welcome route
func UsersRoutes(r *gin.Engine, db *gorm.DB, dbSQLl *sql.DB) {
	auditTrailRepo := auditTrailsRepository.NewAuditTrailsRepository(db)
	auditTrailUC := auditTrailsUsecase.NewAuditTrailsUsecase(auditTrailRepo)

	usersRepo := usersRepository.NewUsersRepository(dbSQLl)
	usersUC := usersUsecase.NewUsersUsecase(usersRepo)
	usersDelivery.NewUsersHandler(r, auditTrailUC, usersUC)

}
