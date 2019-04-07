package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/vicesoftware/vice-go-boilerplate/cmd/webserver/docs"
	"github.com/vicesoftware/vice-go-boilerplate/cmd/webserver/models"
	"github.com/vicesoftware/vice-go-boilerplate/pkg/database"
	"github.com/vicesoftware/vice-go-boilerplate/pkg/log"
	"go.uber.org/zap"
)

type webserver struct {
	addr string
	db   database.DB
}

func (ws *webserver) Start() {
	r := ws.router()

	log.Info("starting http server", zap.String("addr", ws.addr))
	log.Fatal(http.ListenAndServe(ws.addr, r))
}

func (ws *webserver) router() *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	r.NotFoundHandler = handler(notFoundHandler)

	apiv1 := r.PathPrefix("/api/v1").Subrouter()

	apiv1.HandleFunc("/ping", handler(ws.handlePing)).Methods("GET")

	apiv1.HandleFunc("/contacts", handler(ws.handleGetContacts)).Methods("GET")
	apiv1.HandleFunc("/contacts/{contactID}", handler(ws.handleGetContact)).Methods("GET")
	apiv1.HandleFunc("/contacts", handler(ws.handlePostContact)).Methods("POST")
	apiv1.HandleFunc("/contacts/{contactID}", handler(ws.handlePutContact)).Methods("PUT")
	apiv1.HandleFunc("/contacts/{contactID}", handler(ws.handleDeleteContact)).Methods("DELETE")

	apiv1.HandleFunc("/contacts/{contactID}/addresses", handler(ws.handleGetContactAddresses)).Methods("GET")
	apiv1.HandleFunc("/contacts/{contactID}/addresses/{addressID}", handler(ws.handleGetContactAddress)).Methods("GET")
	apiv1.HandleFunc("/contacts/{contactID}/addresses", handler(ws.handlePostContactAddresses)).Methods("POST")
	apiv1.HandleFunc("/contacts/{contactID}/addresses/{addressID}", handler(ws.handlePutContactAddress)).Methods("PUT")
	apiv1.HandleFunc("/contacts/{contactID}/addresses/{addressID}", handler(ws.handleDeleteContactAddress)).Methods("DELETE")

	return r
}

// @Summary Get all contacts
// @Produce json
// @Success 200 {array} models.ContactResponse
// @Router /contacts [get]
func (ws *webserver) handleGetContacts(w http.ResponseWriter, r *http.Request) error {
	// get all contacts
	contacts, err := ws.db.Contacts.GetAll()
	if err != nil {
		return err
	}

	// get all addresses for each contact
	response := make([]models.ContactResponse, 0)
	for _, contact := range contacts {
		addresses, err := ws.db.Addresses.GetAllByContactID(contact.ID)
		if err != nil {
			return err
		}
		response = append(response, models.MapContactResponse(contact, addresses))
	}

	return Ok(w, response)
}

// @Summary Get a contact
// @Produce json
// @Success 200 {object} models.ContactResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Param contactID path int true "Contact ID"
// @Router /contacts/{contactID} [get]
func (ws *webserver) handleGetContact(w http.ResponseWriter, r *http.Request) error {
	// get url params
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["contactID"])
	if err != nil {
		return &invalidRequest{}
	}

	// get contact
	contact, err := ws.db.Contacts.Get(id)
	if err != nil {
		return err
	}

	// get contact addresses
	addresses, err := ws.db.Addresses.GetAllByContactID(contact.ID)
	if err != nil {
		return err
	}

	// create response
	response := models.MapContactResponse(contact, addresses)

	return Ok(w, response)
}

// @Summary Create a contact
// @Param contact body models.ContactRequest true "Create contact"
// @Accept json
// @Produce json
// @Success 200 {object} models.ContactResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /contacts [post]
func (ws *webserver) handlePostContact(w http.ResponseWriter, r *http.Request) error {
	// create var ready to hold decoded json from body
	var request models.ContactRequest

	// decode body
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&request); err != nil {
		return &invalidRequest{}
	}

	// create contact
	create := models.MapCreateContactRequest(request)
	contact, err := ws.db.Contacts.Create(create)
	if err != nil {
		return err
	}

	// create response
	response := models.MapContactResponse(contact, nil)

	return Ok(w, response)
}

// @Summary Update a contact
// @Param contact body models.ContactRequest true "Update contact"
// @Accept json
// @Produce json
// @Success 200 {object} models.ContactResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Param contactID path int true "Contact ID"
// @Router /contacts/{contactID} [put]
func (ws *webserver) handlePutContact(w http.ResponseWriter, r *http.Request) error {
	// get url params
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["contactID"])
	if err != nil {
		return &invalidRequest{}
	}

	// create var ready to hold decoded json from body
	var request models.ContactRequest

	// decode body
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&request); err != nil {
		return &invalidRequest{}
	}

	// update contact
	update := models.MapUpdateContactRequest(id, request)
	contact, err := ws.db.Contacts.Update(update)
	if err != nil {
		return err
	}

	// create response
	response := models.MapContactResponse(contact, nil)

	return Ok(w, response)
}

// @Summary Delete a contact
// @Produce json
// @Success 200 {string} string "{}"
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Param contactID path int true "Contact ID"
// @Router /contacts/{contactID} [delete]
func (ws *webserver) handleDeleteContact(w http.ResponseWriter, r *http.Request) error {
	// get url params
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["contactID"])
	if err != nil {
		return &invalidRequest{}
	}

	// delete contact
	if err = ws.db.Contacts.Delete(id); err != nil {
		return err
	}

	// struct{}{} is an empty object, returns "{}" to the client
	return Ok(w, struct{}{})
}

// @Summary Get all of a contact's addresses
// @Produce json
// @Success 200 {array} models.AddressResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Param contactID path int true "Contact ID"
// @Router /contacts/{contactID}/addresses [get]
func (ws *webserver) handleGetContactAddresses(w http.ResponseWriter, r *http.Request) error {
	// get url params
	vars := mux.Vars(r)
	contactID, err := strconv.Atoi(vars["contactID"])
	if err != nil {
		return &invalidRequest{}
	}

	// get contact addresses
	addresses, err := ws.db.Addresses.GetAllByContactID(contactID)
	if err != nil {
		return err
	}

	// create response
	response := models.MapContactAddresses(addresses)

	return Ok(w, response)
}

// @Summary Get a contact address
// @Produce json
// @Success 200 {object} models.AddressResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Param contactID path int true "Contact ID"
// @Param addressID path int true "Address ID"
// @Router /contacts/{contactID}/addresses/{addressID} [get]
func (ws *webserver) handleGetContactAddress(w http.ResponseWriter, r *http.Request) error {
	// get url params
	vars := mux.Vars(r)
	contactID, err := strconv.Atoi(vars["contactID"])
	if err != nil {
		return &invalidRequest{}
	}
	addressID, err := strconv.Atoi(vars["addressID"])
	if err != nil {
		return &invalidRequest{}
	}

	// get address by ID
	address, err := ws.db.Addresses.Get(addressID)
	if err != nil {
		return err
	}
	// ensure address belongs to contact
	if address.ContactID != contactID {
		return &notFound{}
	}

	// create response
	response := models.MapContactAddress(address)

	return Ok(w, response)
}

// @Summary Create a contact address
// @Param address body models.AddressRequest true "Create address"
// @Accept json
// @Produce json
// @Success 200 {object} models.AddressResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Param contactID path int true "Contact ID"
// @Router /contacts/{contactID}/addresses [post]
func (ws *webserver) handlePostContactAddresses(w http.ResponseWriter, r *http.Request) error {
	// get url params
	vars := mux.Vars(r)
	contactID, err := strconv.Atoi(vars["contactID"])
	if err != nil {
		return &invalidRequest{}
	}

	// create var ready to hold decoded json from body
	var request models.AddressRequest

	// decode body
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&request); err != nil {
		return &invalidRequest{}
	}

	// create contact
	create := models.MapCreateAddressRequest(contactID, request)
	address, err := ws.db.Addresses.Create(create)
	if err != nil {
		return err
	}

	// create response
	response := models.MapContactAddress(address)

	return Ok(w, response)
}

// @Summary Update a contact address
// @Param address body models.AddressRequest true "Update address"
// @Accept json
// @Produce json
// @Success 200 {object} models.AddressResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Param contactID path int true "Contact ID"
// @Param addressID path int true "Address ID"
// @Router /contacts/{contactID}/addresses/{addressID} [put]
func (ws *webserver) handlePutContactAddress(w http.ResponseWriter, r *http.Request) error {
	// get url params
	vars := mux.Vars(r)
	contactID, err := strconv.Atoi(vars["contactID"])
	if err != nil {
		return &invalidRequest{}
	}
	addressID, err := strconv.Atoi(vars["addressID"])
	if err != nil {
		return &invalidRequest{}
	}

	// get address
	address, err := ws.db.Addresses.Get(addressID)
	if err != nil {
		return err
	}
	// ensure address belongs to contact
	if address.ContactID != contactID {
		return &notFound{}
	}

	// create var ready to hold decoded json from body
	var request models.AddressRequest

	// decode body
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&request); err != nil {
		return &invalidRequest{}
	}

	// update address
	update := models.MapUpdateAddressRequest(contactID, addressID, request)
	newAddress, err := ws.db.Addresses.Update(update)
	if err != nil {
		return err
	}

	// create response
	response := models.MapContactAddress(newAddress)

	return Ok(w, response)
}

// @Summary Delete a contact address
// @Produce json
// @Success 200 {string} string "{}"
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Param contactID path int true "Contact ID"
// @Param addressID path int true "Address ID"
// @Router /contacts/{contactID}/addresses/{addressID} [delete]
func (ws *webserver) handleDeleteContactAddress(w http.ResponseWriter, r *http.Request) error {
	// get url params
	vars := mux.Vars(r)
	contactID, err := strconv.Atoi(vars["contactID"])
	if err != nil {
		return &invalidRequest{}
	}
	addressID, err := strconv.Atoi(vars["addressID"])
	if err != nil {
		return &invalidRequest{}
	}

	// get address
	address, err := ws.db.Addresses.Get(addressID)
	if err != nil {
		return err
	}
	// ensure address belongs to contact
	if address.ContactID != contactID {
		return &notFound{}
	}

	// delete address
	if err = ws.db.Addresses.Delete(addressID); err != nil {
		return err
	}

	// struct{}{} is an empty object, returns "{}" to the client
	return Ok(w, struct{}{})
}
