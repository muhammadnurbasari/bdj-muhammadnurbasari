package appServerModel

type (
	SetConnDb struct {
		DbHost     string
		DbPort     string
		DbUser     string
		DbPass     string
		DbName     string
		DbDriver   string
		DbSSL      string
		DbTimezone string
		MaxIdle    int
		MaxConn    int
	}
)
