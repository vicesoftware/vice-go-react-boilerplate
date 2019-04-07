package database

import "github.com/jinzhu/gorm"

type AddressProvider struct {
	db     *gorm.DB
	parent DB
}

func (a AddressProvider) Create(address Address) (Address, error) {
	if address.ID != 0 {
		return Address{}, &invalidRequest{"create address", "id must be 0"}
	}
	if db := a.db.Create(&address); db.Error != nil {
		return Address{}, db.Error
	}
	return address, nil
}

func (a AddressProvider) Get(id int) (Address, error) {
	address := Address{ID: id}
	if db := a.db.Take(&address); db.Error != nil {
		if IsNotFound(db.Error) {
			return Address{}, &recordNotFound{"get address", address.ID}
		}
		return Address{}, db.Error
	}
	return address, nil
}

func (a AddressProvider) GetAll() ([]Address, error) {
	addresses := make([]Address, 0)
	if db := a.db.Order("id").Find(&addresses); db.Error != nil {
		return nil, db.Error
	}
	return addresses, nil
}

func (a AddressProvider) GetAllByContactID(contactID int) ([]Address, error) {
	_, err := a.parent.Contacts.Get(contactID)
	if err != nil {
		return nil, err
	}

	addresses := make([]Address, 0)
	if db := a.db.Order("id").Where(Address{ContactID: contactID}).Find(&addresses); db.Error != nil {
		return nil, db.Error
	}
	return addresses, nil
}

func (a AddressProvider) Update(address Address) (Address, error) {
	existing, err := a.Get(address.ID)
	if err != nil {
		if IsNotFound(err) {
			return Address{}, &recordNotFound{"update address", address.ID}
		}
		return Address{}, err
	}

	if address.CreatedAt.IsZero() {
		address.CreatedAt = existing.CreatedAt
	}

	if db := a.db.Save(&address); db.Error != nil {
		return Address{}, db.Error
	}
	return address, nil
}

func (a AddressProvider) Delete(id int) error {
	db := a.db.Delete(&Address{ID: id})
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return &recordNotFound{"delete address", id}
	}
	return nil
}

func (a AddressProvider) DeleteAllByContactID(contactID int) error {
	_, err := a.parent.Contacts.Get(contactID)
	if err != nil {
		return err
	}

	db := a.db.Delete(&Address{ContactID: contactID})
	if db.Error != nil {
		return db.Error
	}
	return nil
}
