package models

import "github.com/vicesoftware/vice-go-boilerplate/pkg/database"

func MapContactResponse(contact database.Contact, addresses []database.Address) ContactResponse {
	return ContactResponse{
		ID:        contact.ID,
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Addresses: MapContactAddresses(addresses),
		CreatedAt: toMS(contact.CreatedAt),
		UpdatedAt: toMS(contact.UpdatedAt),
	}
}

func MapContactAddresses(addresses []database.Address) []AddressResponse {
	resp := make([]AddressResponse, 0, len(addresses))
	for _, address := range addresses {
		resp = append(resp, MapContactAddress(address))
	}
	return resp
}

func MapContactAddress(address database.Address) AddressResponse {
	return AddressResponse{
		ID:            address.ID,
		Line1:         address.Line1,
		Line2:         address.Line2,
		City:          address.City,
		StateProvince: address.StateProvince,
		PostalCode:    address.PostalCode,
		CreatedAt:     toMS(address.CreatedAt),
		UpdatedAt:     toMS(address.UpdatedAt),
	}
}

func MapCreateContactRequest(request ContactRequest) database.Contact {
	return database.Contact{
		FirstName: request.FirstName,
		LastName:  request.LastName,
	}
}

func MapUpdateContactRequest(contactID int, request ContactRequest) database.Contact {
	return database.Contact{
		ID:        contactID,
		FirstName: request.FirstName,
		LastName:  request.LastName,
	}
}

func MapCreateAddressRequest(contactID int, request AddressRequest) database.Address {
	return database.Address{
		ContactID:     contactID,
		Line1:         request.Line1,
		Line2:         request.Line2,
		City:          request.City,
		StateProvince: request.StateProvince,
		PostalCode:    request.PostalCode,
	}
}

func MapUpdateAddressRequest(contactID, addressID int, request AddressRequest) database.Address {
	return database.Address{
		ID:            addressID,
		ContactID:     contactID,
		Line1:         request.Line1,
		Line2:         request.Line2,
		City:          request.City,
		StateProvince: request.StateProvince,
		PostalCode:    request.PostalCode,
	}
}
