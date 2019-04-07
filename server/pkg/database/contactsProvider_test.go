package database

import (
	"testing"
	"time"
)

func TestContactProvider_Create(t *testing.T) {
	// arrange
	db, err := New(testSettings)
	if err != nil {
		t.Fatal(err)
	}

	contact := Contact{
		FirstName: "John",
		LastName:  "Doe",
	}

	// act
	newContact, err := db.Contacts.Create(contact)
	if err != nil {
		t.Fatal(err)
	}

	// assert
	if contact.FirstName != newContact.FirstName {
		t.Errorf("FirstName, want: %q got: %q", contact.FirstName, newContact.FirstName)
	}
	if contact.LastName != newContact.LastName {
		t.Errorf("LastName, want: %q got: %q", contact.LastName, newContact.LastName)
	}
	if newContact.ID == 0 {
		t.Errorf("ID, want: non-zero got: %d", newContact.ID)
	}
	if newContact.CreatedAt.IsZero() {
		t.Errorf("CreatedAt, want: non-zero got: %v", newContact.CreatedAt)
	}
	if !newContact.UpdatedAt.Equal(newContact.CreatedAt) {
		t.Errorf("UpdatedAt, want: %v got: %v", newContact.CreatedAt, newContact.UpdatedAt)
	}
}

func TestContactProvider_Get(t *testing.T) {
	// arrange
	db, err := New(testSettings)
	if err != nil {
		t.Fatal(err)
	}

	contact := Contact{
		FirstName: "John",
		LastName:  "Doe",
	}

	resp, err := db.Contacts.Create(contact)
	if err != nil {
		t.Fatal(err)
	}

	// act
	newContact, err := db.Contacts.Get(resp.ID)
	if err != nil {
		t.Fatal(err)
	}

	// assert
	if contact.FirstName != newContact.FirstName {
		t.Errorf("FirstName, want: %q got: %q", contact.FirstName, newContact.FirstName)
	}
	if contact.LastName != newContact.LastName {
		t.Errorf("LastName, want: %q got: %q", contact.LastName, newContact.LastName)
	}
	if newContact.ID == 0 {
		t.Errorf("ID, want: non-zero got: %d", newContact.ID)
	}
	if newContact.CreatedAt.IsZero() {
		t.Errorf("CreatedAt, want: non-zero got: %v", newContact.CreatedAt)
	}
	if !newContact.UpdatedAt.Equal(newContact.CreatedAt) {
		t.Errorf("UpdatedAt, want: %v got: %v", newContact.CreatedAt, newContact.UpdatedAt)
	}
}

func TestContactProvider_GetAll(t *testing.T) {
	// arrange
	db, err := New(testSettings)
	if err != nil {
		t.Fatal(err)
	}

	if err = deleteAll(); err != nil {
		t.Fatal(err)
	}

	contacts := []Contact{
		{FirstName: "John", LastName: "Doe"},
		{FirstName: "Jane", LastName: "Doe"},
	}

	for _, contact := range contacts {
		_, err := db.Contacts.Create(contact)
		if err != nil {
			t.Fatal(err)
		}
	}

	// act
	newContacts, err := db.Contacts.GetAll()
	if err != nil {
		t.Fatal(err)
	}

	// assert
	if len(contacts) != len(newContacts) {
		t.Fatalf("len(newContacts) want: %d got: %d", len(contacts), len(newContacts))
	}

	for i := 0; i < len(newContacts); i++ {
		contact := contacts[i]
		newContact := newContacts[i]

		if contact.FirstName != newContact.FirstName {
			t.Errorf("FirstName, want: %q got: %q", contact.FirstName, newContact.FirstName)
		}
		if contact.LastName != newContact.LastName {
			t.Errorf("LastName, want: %q got: %q", contact.LastName, newContact.LastName)
		}
		if newContact.ID == 0 {
			t.Errorf("ID, want: non-zero got: %d", newContact.ID)
		}
		if newContact.CreatedAt.IsZero() {
			t.Errorf("CreatedAt, want: non-zero got: %v", newContact.CreatedAt)
		}
		if !newContact.UpdatedAt.Equal(newContact.CreatedAt) {
			t.Errorf("UpdatedAt, want: %v got: %v", newContact.CreatedAt, newContact.UpdatedAt)
		}
	}
}

func TestContactProvider_GetManyIndividually(t *testing.T) {
	// arrange
	db, err := New(testSettings)
	if err != nil {
		t.Fatal(err)
	}

	if err = deleteAll(); err != nil {
		t.Fatal(err)
	}

	contacts := []Contact{
		{FirstName: "John", LastName: "Doe"},
		{FirstName: "Jane", LastName: "Doe"},
		{FirstName: "Jack", LastName: "Doe"},
	}

	var ids []int
	for _, contact := range contacts {
		newContact, err := db.Contacts.Create(contact)
		if err != nil {
			t.Fatal(err)
		}
		ids = append(ids, newContact.ID)
	}

	// act
	var newContacts []Contact
	for i := len(ids) - 1; i >= 0; i-- {
		// retrieve from DB in reverse order to make sure there isn't some condition for sequential reads
		// that leads to a false pass
		id := ids[i]
		newContact, err := db.Contacts.Get(id)
		if err != nil {
			t.Fatal(err)
		}
		newContacts = append([]Contact{newContact}, newContacts...)
	}

	// assert
	if len(contacts) != len(newContacts) {
		t.Fatalf("len(newContacts) want: %d got: %d", len(contacts), len(newContacts))
	}

	for i := 0; i < len(newContacts); i++ {
		contact := contacts[i]
		newContact := newContacts[i]

		if contact.FirstName != newContact.FirstName {
			t.Errorf("FirstName, want: %q got: %q", contact.FirstName, newContact.FirstName)
		}
		if contact.LastName != newContact.LastName {
			t.Errorf("LastName, want: %q got: %q", contact.LastName, newContact.LastName)
		}
		if newContact.ID == 0 {
			t.Errorf("ID, want: non-zero got: %d", newContact.ID)
		}
		if newContact.CreatedAt.IsZero() {
			t.Errorf("CreatedAt, want: non-zero got: %v", newContact.CreatedAt)
		}
		if !newContact.UpdatedAt.Equal(newContact.CreatedAt) {
			t.Errorf("UpdatedAt, want: %v got: %v", newContact.CreatedAt, newContact.UpdatedAt)
		}
	}
}

func TestContactProvider_Update(t *testing.T) {
	// arrange
	db, err := New(testSettings)
	if err != nil {
		t.Fatal(err)
	}

	if err = deleteAll(); err != nil {
		t.Fatal(err)
	}

	contacts := []Contact{
		{FirstName: "John", LastName: "Doe"},
		{FirstName: "Jane", LastName: "Doe"},
	}

	for i := range contacts {
		contacts[i], err = db.Contacts.Create(contacts[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	const newFirstName = "Jack"
	const newLastName = "Johnson"

	// act
	contacts[0].FirstName = newFirstName
	contacts[0].LastName = newLastName
	update, err := db.Contacts.Update(contacts[0])
	if err != nil {
		t.Fatal(err)
	}

	// assert
	if newFirstName != update.FirstName {
		t.Errorf("update.FirstName, want: %q got: %q", newFirstName, update.FirstName)
	}
	if newLastName != update.LastName {
		t.Errorf("update.LastName, want: %q got: %q", newLastName, update.LastName)
	}
	if contacts[0].ID != update.ID {
		t.Errorf("update.ID, want: %v got: %v", contacts[0].ID, update.ID)
	}
	if !update.CreatedAt.Equal(contacts[0].CreatedAt) {
		t.Errorf("update.CreatedAt, want: %v got: %v", contacts[0].CreatedAt, update.CreatedAt)
	}
	if !update.UpdatedAt.After(contacts[0].UpdatedAt) {
		t.Errorf("update.UpdatedAt, want: time after %v got: %v", contacts[0].UpdatedAt, update.UpdatedAt)
	}

	// assert that querying returns the same results and no other records were modified
	contacts[0] = update

	newContacts, err := db.Contacts.GetAll()
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(newContacts); i++ {
		contact := contacts[i]
		newContact := newContacts[i]

		if contact.FirstName != newContact.FirstName {
			t.Errorf("FirstName, want: %q got: %q", contact.FirstName, newContact.FirstName)
		}
		if contact.LastName != newContact.LastName {
			t.Errorf("LastName, want: %q got: %q", contact.LastName, newContact.LastName)
		}
		if contact.ID != newContact.ID {
			t.Errorf("ID, want: %d got: %d", contact.ID, newContact.ID)
		}

		// due to how timestamps are stored we may not get the precision back which was sent to the database

		newCreatedAt := newContact.CreatedAt.Round(time.Millisecond)
		newUpdatedAt := newContact.UpdatedAt.Round(time.Millisecond)
		oldCreatedAt := contact.CreatedAt.Round(time.Millisecond)
		oldUpdatedAt := contact.UpdatedAt.Round(time.Millisecond)

		if !newCreatedAt.Equal(oldCreatedAt) {
			t.Errorf("CreatedAt, want: %v got: %v", oldCreatedAt, newCreatedAt)
		}
		if !newUpdatedAt.Equal(oldUpdatedAt) {
			t.Errorf("UpdatedAt, want: %v got: %v", oldUpdatedAt, newUpdatedAt)
		}
	}
}

func TestContactProvider_Delete(t *testing.T) {
	// arrange
	db, err := New(testSettings)
	if err != nil {
		t.Fatal(err)
	}

	if err = deleteAll(); err != nil {
		t.Fatal(err)
	}

	contacts := []Contact{
		{FirstName: "John", LastName: "Doe"},
		{FirstName: "Jane", LastName: "Doe"},
	}

	for i := range contacts {
		contacts[i], err = db.Contacts.Create(contacts[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	// act
	for _, contact := range contacts {
		// make sure current contact exists
		if _, err = db.Contacts.Get(contact.ID); err != nil {
			t.Fatal(err)
		}

		// delete
		if err = db.Contacts.Delete(contact.ID); err != nil {
			t.Fatal(err)
		}

		// assert contact is deleted
		_, err = db.Contacts.Get(contact.ID)
		if err == nil {
			t.Errorf("wanted contact ID '%d' to be deleted", contact.ID)
		} else if !IsNotFound(err) {
			t.Fatal(err)
		}
	}
}

func TestContactProvider_UpdateDeletedRecordShouldReturnIsNotFound(t *testing.T) {
	// arrange
	db, err := New(testSettings)
	if err != nil {
		t.Fatal(err)
	}

	contact := Contact{FirstName: "John", LastName: "Doe"}

	newContact, err := db.Contacts.Create(contact)
	if err != nil {
		t.Fatal(err)
	}

	if err = db.Contacts.Delete(newContact.ID); err != nil {
		t.Fatal(err)
	}

	// make sure contact is deleted...
	_, err = db.Contacts.Get(newContact.ID)
	if err == nil {
		t.Errorf("wanted contact ID '%d' to be deleted", newContact.ID)
	} else if !IsNotFound(err) {
		t.Fatal(err)
	}

	// act
	got, err := db.Contacts.Update(newContact)

	// assert
	if err == nil {
		t.Errorf("expected contact ID '%d' to be deleted, got %#v", newContact.ID, got)
	} else if !IsNotFound(err) {
		t.Fatal(err)
	}
}

func TestContactProvider_GetAllWhenNoRecordsReturnsEmptySlice(t *testing.T) {
	// arrange
	db, err := New(testSettings)
	if err != nil {
		t.Fatal(err)
	}

	if err = deleteAll(); err != nil {
		t.Fatal(err)
	}

	// act
	records, err := db.Contacts.GetAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(records) != 0 {
		t.Errorf("len(records), want: %d got: %d", 0, len(records))
	}
}

func TestContactProvider_CreateReturnsErrorIfIDNotZero(t *testing.T) {
	// arrange
	db, err := New(testSettings)
	if err != nil {
		t.Fatal(err)
	}

	// act
	_, err = db.Contacts.Create(Contact{ID: 99999})

	// assert
	if !IsInvalidRequest(err) {
		t.Fatal("expected error, got <nil>")
	}
}

func TestContactProvider_UpdateReturnsErrorIfIDZero(t *testing.T) {
	// arrange
	db, err := New(testSettings)
	if err != nil {
		t.Fatal(err)
	}

	// act
	_, err = db.Contacts.Update(Contact{FirstName: "John", LastName: "Doe"})

	// assert
	if !IsNotFound(err) {
		t.Fatal("expected error, got <nil>")
	}
}

func TestContactProvider_DeleteReturnsErrorIfIDZero(t *testing.T) {
	// arrange
	db, err := New(testSettings)
	if err != nil {
		t.Fatal(err)
	}

	// act
	err = db.Contacts.Delete(0)

	// assert
	if !IsNotFound(err) {
		t.Fatal("expected error, got <nil>")
	}
}

func TestContactProvider_DeleteReturnsErrorIfNotFound(t *testing.T) {
	// arrange
	db, err := New(testSettings)
	if err != nil {
		t.Fatal(err)
	}

	if err = deleteAll(); err != nil {
		t.Fatal(err)
	}

	// act
	err = db.Contacts.Delete(1)

	// assert
	if !IsNotFound(err) {
		t.Fatal("expected error, got <nil>")
	}
}
