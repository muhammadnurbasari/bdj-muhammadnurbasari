package registry

import (
	"bdj-muhammadnurbasari/app/appServer"
	"bdj-muhammadnurbasari/database/connect-db/mysql"
	"bdj-muhammadnurbasari/models/appServerModel"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type AppRegistry struct {
	ConnPublic *gorm.DB
	ConnSql    *sql.DB
	serverHttp *appServer.HttpHandler
}

//NewAppRegistry will return new object for App Registry
func NewAppRegistry() *AppRegistry {
	return &AppRegistry{}
}

func initializeEnvPublic() (*appServerModel.SetConnDb, error) {
	moduleName := "registry.initializeEnvPublic"
	log.Debug().Msg("Read file env for DB PUBLIC. . .")
	err := godotenv.Load("config/sample.env")
	if err != nil {
		return nil, errors.New(moduleName + ".err : " + err.Error())
	}

	port := os.Getenv("PORT")
	if port == "" {
		return nil, errors.New(moduleName + ".error : port cant empty")
	}

	log.Debug().Msg("mapping config database . . .")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbTimezone := os.Getenv("DB_TIME_ZONE")
	// dbSSL := os.Getenv("DB_SSL")
	maxIdle := os.Getenv("MAX_IDLE")
	maxConn := os.Getenv("MAX_CONN")

	myMaxIdle, errMaxIdle := strconv.Atoi(maxIdle)
	if errMaxIdle != nil {
		return nil, errors.New(moduleName + ".errMaxIdle : " + errMaxIdle.Error())
	}

	myMaxConn, errMaxConn := strconv.Atoi(maxConn)
	if errMaxConn != nil {
		return nil, errors.New(moduleName + ".errMaxConn : " + errMaxConn.Error())
	}

	log.Debug().Msg("set config database . . .")
	setConnDb := appServerModel.SetConnDb{
		DbHost:     dbHost,
		DbPort:     dbPort,
		DbUser:     dbUser,
		DbPass:     dbPass,
		DbName:     dbName,
		DbSSL:      "disable",
		DbTimezone: dbTimezone,
		MaxIdle:    myMaxIdle,
		MaxConn:    myMaxConn,
	}

	return &setConnDb, nil
}

//StartServer will do the server initialization
func (reg *AppRegistry) StartServer() {
	printAsciiArt()
	log.Debug().Msg("prepare start server . . .")
	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)
	log.Debug().Msg("initial database . . .")
	setConnDbPublic, errInitialEnvPublic := initializeEnvPublic()
	if errInitialEnvPublic != nil {
		log.Error().Msg(errInitialEnvPublic.Error())
		return
	}

	//initial database Public
	connPublic, errConnPublic := getDBConnection(setConnDbPublic)
	if errConnPublic != nil {
		log.Error().Msg(errConnPublic.Error())
		return
	}
	connSQL, errConnSQL := getDBConnectionSQL(setConnDbPublic)
	if errConnSQL != nil {
		log.Error().Msg(errConnSQL.Error())
		return
	}

	reg.ConnPublic = connPublic
	reg.ConnSql = connSQL

	//close connection
	defer func() {
		log.Info().Msg("Close connection . . .")
		errConnClose := reg.ConnPublic.Close()
		if errConnClose != nil {
			log.Error().Msg("Service bdj.errConnClose : " + errConnClose.Error())
		}
	}()

	errApp := reg.initializeAppRegistry()
	if errApp != nil {
		log.Error().Msg(errApp.Error())
		return
	}

	//Run Swagger
	log.Info().Msg("Swagger run on /docs/swagger/index.html")
	reg.serverHttp.RunSwaggerMiddleware()

	//Run HTTP Server
	appPort := os.Getenv("PORT")
	log.Info().Msg("Last Update : " + time.Now().Format("2006-01-02 15:04:05"))
	log.Info().Msg("REST API BDJ Running version 0.1.5 at port : " + appPort)

	if errHTTP := reg.serverHttp.RunHttpServer(); errHTTP != nil {
		log.Error().Msg(errHTTP.Error())
	}
}

func (reg *AppRegistry) initializeAppRegistry() error {
	//initial read file env
	log.Debug().Msg("prepare initial env . . .")
	port := os.Getenv("PORT")
	//initial handler
	errHandler := reg.initializeHandler(port)
	if errHandler != nil {
		return errHandler
	}

	//initial modules
	reg.initializeDomainModules()
	reg.insertData()

	return nil
}

func (reg *AppRegistry) initializeHandler(appPort string) error {
	//Register HTTP Server Handler
	if appPort == "" {
		return errors.New("registry.error : port cant empty")
	}
	reg.serverHttp = appServer.NewHTTPHandler(":" + appPort)
	return nil
}

func getDBConnection(data *appServerModel.SetConnDb) (*gorm.DB, error) {
	log.Debug().Msg("data.DbHost : " + data.DbHost)
	log.Debug().Msg("data.DbUser : " + data.DbUser)
	log.Debug().Msg("data.DbPass : " + data.DbPass)
	log.Debug().Msg("data.DbName : " + data.DbName)
	log.Debug().Msg("data.DbPort : " + data.DbPort)
	// log.Debug().Msg("data.DbSSL : " + data.DbSSL)
	// log.Debug().Msg("data.DbTimezone : " + data.DbTimezone)
	conn, err := mysql.ConnMySQLORM(data.DbHost, data.DbPort, data.DbUser, data.DbPass,
		data.DbName, data.MaxIdle, data.MaxConn)
	if err != nil {
		return nil, errors.New("registry.getDBConnection.err : " + err.Error())
	}

	return conn, nil
}

func getDBConnectionSQL(data *appServerModel.SetConnDb) (*sql.DB, error) {
	log.Debug().Msg("data.DbHost : " + data.DbHost)
	log.Debug().Msg("data.DbUser : " + data.DbUser)
	log.Debug().Msg("data.DbPass : " + data.DbPass)
	log.Debug().Msg("data.DbName : " + data.DbName)
	log.Debug().Msg("data.DbPort : " + data.DbPort)
	log.Debug().Msg("data.DbSSL : " + data.DbSSL)
	log.Debug().Msg("data.DbTimezone : " + data.DbTimezone)
	// conn, err := mysql.ConnMySQLORM(data.DbHost, data.DbPort, data.DbUser, data.DbPass,
	// 	data.DbName, data.MaxIdle, data.MaxConn)
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", data.DbUser, data.DbPass, data.DbHost, data.DbPort, data.DbName))

	if err != nil {
		return nil, errors.New("registry.getDBConnection.err : " + err.Error())
	}

	return db, nil
}
