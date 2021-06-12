package main

import (
	"flag"
	"fmt"
	"os"

	"bdj-muhammadnurbasari/database/migrates"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	err := godotenv.Load("config/sample.env")
	if err != nil {
		log.Error().Msg("Failed read configuration database")
		return
	}

	pathMigration := os.Getenv("PATH_MIGRATION")
	migrationDir := flag.String("migration-dir", pathMigration, "migration directory")

	upMigration := flag.Bool("up", false, "Up migration flag")
	downMigration := flag.Bool("down", false, "Down migration flag")

	newMigrationFile := flag.Bool("create", false, "Create new migration file")
	newMigrationFileName := flag.String("filename", "", "New migration filename")
	flag.Parse()

	if *newMigrationFile {
		if *newMigrationFileName == "" {
			fmt.Println("Please specify migration filename with --filename")
			migrates.ShowHelp()
			return
		}

		// create a new migration file
		err := migrates.CreateNewMigrationFile(*migrationDir, *newMigrationFileName)
		if err != nil {
			fmt.Println("Failed to create migration file " + err.Error())
		}
		return
	}

	//check if at least up or down flag is specified
	if !(*upMigration || *downMigration) {
		log.Error().Msg("please specify --up or --down for migration")
		migrates.ShowHelp()
		return
	}

	//check migration direction up/down
	if *upMigration && *downMigration {
		log.Warn().Msg("use --up or --down at once only")
		migrates.ShowHelp()
		return
	}

	//setting db config
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbDriver := os.Getenv("DB_DRIVER")
	migrationConf, errMigrationConf := migrates.NewMigrationConfig(*migrationDir, dbHost, dbPort, dbUser, dbPass, dbName, dbDriver)
	if errMigrationConf != nil {
		log.Error().Msg(errMigrationConf.Error())
		return
	}

	defer func() {
		errConnClose := migrationConf.Db.Close()
		if errConnClose != nil {
			log.Error().Msg("errConnClose : " + errConnClose.Error())
		}
	}()

	if *upMigration {
		err = migrates.MigrateUp(migrationConf)
		if err != nil {
			log.Error().Msg(err.Error())
			return
		}
	} else if *downMigration {
		err = migrates.MigrateDown(migrationConf)
		if err != nil {
			log.Error().Msg(err.Error())
			return
		}
	}
}
