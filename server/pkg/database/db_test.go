package database

import (
	"fmt"
	"testing"
)

var testSettings = Settings{
	Host:     "127.0.0.1",
	Port:     5434,
	User:     "postgres",
	Password: "password",
	DBName:   "vicetestdb",
}

func TestGetConnectionString(t *testing.T) {
	// arrange
	// taken care of by initializing 'testSettings' variable

	// act
	got := getConnectionString(testSettings)

	// assert
	want := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		testSettings.Host, testSettings.Port, testSettings.User, testSettings.DBName, testSettings.Password, "disable")

	if got != want {
		t.Errorf("got: %q want: %q", got, want)
	}
}

func TestNew(t *testing.T) {
	// arrange
	// taken care of by initializing 'testSettings' variable

	// act
	_, err := New(testSettings)

	// assert
	if err != nil {
		t.Fatal(err)
	}
}

func TestNew_BadSettingsReturnsError(t *testing.T) {
	// arrange
	settings := Settings{}

	// act
	_, err := New(settings)

	// assert
	if err == nil {
		t.Errorf("expected error, got <nil>")
	}
}

func deleteAll() error {
	db, err := New(testSettings)
	if err != nil {
		return err
	}

	if clone := db.db.Delete(Address{}); clone.Error != nil {
		return clone.Error
	}
	if clone := db.db.Delete(Contact{}); clone.Error != nil {
		return clone.Error
	}

	return nil
}
