package database

import (
	"github.com/jinzhu/gorm"
)

type ContactProvider struct {
	db     *gorm.DB
	parent DB
}

func (c ContactProvider) Create(contact Contact) (Contact, error) {
	if contact.ID != 0 {
		return Contact{}, &invalidRequest{"create contact", "id must be 0"}
	}
	if db := c.db.Create(&contact); db.Error != nil {
		return Contact{}, db.Error
	}
	return contact, nil
}

func (c ContactProvider) Get(id int) (Contact, error) {
	contact := Contact{ID: id}
	if db := c.db.Take(&contact); db.Error != nil {
		if IsNotFound(db.Error) {
			return Contact{}, &recordNotFound{"get contact", contact.ID}
		}
		return Contact{}, db.Error
	}
	return contact, nil
}

func (c ContactProvider) GetAll() ([]Contact, error) {
	contacts := make([]Contact, 0)
	if db := c.db.Order("id").Find(&contacts); db.Error != nil {
		return nil, db.Error
	}
	return contacts, nil
}

func (c ContactProvider) Update(contact Contact) (Contact, error) {
	existing, err := c.Get(contact.ID)
	if err != nil {
		if IsNotFound(err) {
			return Contact{}, &recordNotFound{"update contact", contact.ID}
		}
		return Contact{}, err
	}

	if contact.CreatedAt.IsZero() {
		contact.CreatedAt = existing.CreatedAt
	}

	if db := c.db.Save(&contact); db.Error != nil {
		return Contact{}, db.Error
	}
	return contact, nil
}

func (c ContactProvider) Delete(id int) error {
	if err := c.parent.Addresses.DeleteAllByContactID(id); err != nil {
		return err
	}

	db := c.db.Delete(&Contact{ID: id})
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return &recordNotFound{"delete contact", id}
	}
	return nil
}
