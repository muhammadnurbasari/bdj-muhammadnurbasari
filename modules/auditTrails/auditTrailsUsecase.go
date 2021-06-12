package auditTrails

import "github.com/gin-gonic/gin"

//AuditTrailsUsecase - usecase interface for module audit trails
type AuditTrailsUsecase interface {
	MiddlewareAuditTrail(context *gin.Context)
}
