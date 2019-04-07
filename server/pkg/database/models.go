package database

import (
	"time"
)

type Contact struct {
	ID        int
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Address struct {
	ID            int
	ContactID     int
	Line1         string
	Line2         *string
	City          string
	StateProvince string
	PostalCode    string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
