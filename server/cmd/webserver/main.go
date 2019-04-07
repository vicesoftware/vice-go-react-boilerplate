package main

import (
	"os"

	"github.com/vicesoftware/vice-go-boilerplate/pkg/database"
	"github.com/vicesoftware/vice-go-boilerplate/pkg/log"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("skeleton", "A skeleton REST API that uses Postgres.")

	flagListen     = app.Flag("listen", "The HTTP listen address.").Default("127.0.0.1:8423").String()
	flagDBHost     = app.Flag("db-host", "The database host.").Default("127.0.0.1").String()
	flagDBPort     = app.Flag("db-port", "The database port.").Default("5432").Int()
	flagDBUser     = app.Flag("db-user", "The database user.").Default("vice_boilerplate_user").String()
	flagDBPassword = app.Flag("db-password", "The database user's password.").Default("vicesoftware").String()
	flagDBName     = app.Flag("db-name", "The database name.").Default("vice_boilerplate").String()
	flagDBSSL      = app.Flag("db-ssl", "The database SSL mode.").Default("disable").String()
)

// @title Vice Software Example API
// @version 1
// @BasePath /api/v1

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	dbSettings := database.Settings{
		Host:     *flagDBHost,
		Port:     *flagDBPort,
		User:     *flagDBUser,
		Password: *flagDBPassword,
		DBName:   *flagDBName,
		SSLMode:  *flagDBSSL,
	}

	log.Info("connecting to the database...")

	db, err := database.New(dbSettings)
	if err != nil {
		log.Fatal(err)
	}

	ws := webserver{addr: *flagListen, db: db}
	ws.Start()
}
