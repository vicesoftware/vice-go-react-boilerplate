package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func IsNotFound(err error) bool {
	if _, ok := err.(*recordNotFound); ok {
		return true
	}
	return gorm.IsRecordNotFoundError(err)
}

func IsInvalidRequest(err error) bool {
	_, ok := err.(*invalidRequest)
	return ok
}

type recordNotFound struct {
	action string
	id     int
}

func (e *recordNotFound) Error() string {
	return fmt.Sprintf("%s ID %d: record not found", e.action, e.id)
}

type invalidRequest struct {
	action  string
	message string
}

func (e *invalidRequest) Error() string {
	return fmt.Sprintf("%s: %s", e.action, e.message)
}
