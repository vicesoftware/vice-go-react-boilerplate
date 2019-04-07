package models

type ContactRequest struct {
	FirstName string `json:"firstName" example:"John"`
	LastName  string `json:"lastName" example:"Doe"`
}

type AddressRequest struct {
	Line1         string  `json:"line1" example:"1600 Pennsylvania Ave."`
	Line2         *string `json:"line2,omitempty" example:"Ste. 1234"`
	City          string  `json:"city" example:"Washington"`
	StateProvince string  `json:"stateProvince" example:"DC"`
	PostalCode    string  `json:"postalCode" example:"20006"`
}
