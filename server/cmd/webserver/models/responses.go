package models

type ErrorResponse struct {
	Error string `json:"error"`
}

type ContactResponse struct {
	ID        int               `json:"id" example:"1"`
	FirstName string            `json:"firstName" example:"John"`
	LastName  string            `json:"lastName" example:"Doe"`
	Addresses []AddressResponse `json:"addresses"`
	CreatedAt int64             `json:"createdAt" example:"1554441489907"`
	UpdatedAt int64             `json:"updatedAt" example:"1554441489907"`
}

type AddressResponse struct {
	ID            int     `json:"id" example:"1"`
	Line1         string  `json:"line1" example:"1600 Pennsylvania Ave."`
	Line2         *string `json:"line2,omitempty" example:"Ste. 1234"`
	City          string  `json:"city" example:"Washington"`
	StateProvince string  `json:"stateProvince" example:"DC"`
	PostalCode    string  `json:"postalCode" example:"20006"`
	CreatedAt     int64   `json:"createdAt" example:"1554441489907"`
	UpdatedAt     int64   `json:"updatedAt" example:"1554441489907"`
}
