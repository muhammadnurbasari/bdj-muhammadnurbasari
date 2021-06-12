package auditTrailsModel

import "github.com/jinzhu/gorm"

// mapping table
type (
	// AuditTrails - object table audit_trails
	AuditTrails struct {
		gorm.Model
		UrlApi       string `gorm:"column:url_api"`
		FunctionName string `gorm:"column:function_name"`
		IpAddress    string `gorm:"column:ip_address"`
		MethodApi    string `gorm:"column:method_api"`
		ResponseCode int16  `gorm:"column:response_code"`
		RequestBody  string `gorm:"column:request_body"`
		ResponseBody string `gorm:"column:response_body"`
	}

	// DataAuditTrails - parameter to insert data audit_trails
	DataAuditTrails struct {
		UrlApi       string
		FunctionName string
		IpAddress    string
		MethodApi    string
		ResponseCode int16
		RequestBody  string
		ResponseBody string
	}
)
