package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
	db        *gorm.DB
	Contacts  *ContactProvider
	Addresses *AddressProvider
}

type Settings struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string // disable, require, verify-ca, verify-full. see https://godoc.org/github.com/lib/pq
}

func New(settings Settings) (DB, error) {
	connectionString := getConnectionString(settings)
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return DB{}, err
	}

	d := DB{db: db}
	d.Contacts = &ContactProvider{db: db, parent: d}
	d.Addresses = &AddressProvider{db: db, parent: d}

	return d, nil
}

func getConnectionString(settings Settings) string {
	sslMode := settings.SSLMode
	if sslMode == "" {
		sslMode = "disable"
	}
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		settings.Host, settings.Port, settings.User, settings.DBName, settings.Password, sslMode)
}
