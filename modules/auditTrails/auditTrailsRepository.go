package auditTrails

import "bdj-muhammadnurbasari/models/auditTrailsModel"

//AuditTrailRepository - repository interface for module audit trails
type AuditTrailsRepository interface {
	InsertAuditTrails(data *auditTrailsModel.DataAuditTrails) error
}
