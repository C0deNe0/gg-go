package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// this will help us through out to write the json for the ERRORS or SUCCESS or to send any type of Response
func WriteJSON(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Add("Content-Type", "application-json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// Creating a struct for the structured show case of the errors in the file
// this just have the error in the string format
type ApiError struct {
	Error string
}

// ANY function which will match this signature( => we are taking that as a TYPE.)
type apiFunc func(w http.ResponseWriter, r *http.Request) error

/*
IDEALOGY := basicallay we don't want to solve the error in the same function we want to return it to the user,
RESONS: 1) It makes the codebase clutter
		2) to keep it clean
*/

//	GENERATOR FUNCTION
//
// we are creating this because we want to send back the http.HandleFunc from the function which was sending the error( so we just wrapped it). using the signature as the context
func MakeHTTPHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			//handle the error
			//We didn't return error here because the basic httpHandleFunc doesn't return it
			//we created the ApiError struct for this specific reason as we cannot return the error here so we just showed the error back in the string format
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()}) //this err.Error() just changed the error into a string format
		}
	}
}

// this is the structure of the server we are creating
type APIServer struct {
	ListenAddr string
	store      Storage
}

// this is like a constructor to the struct we cerated above  which is returning the new server it will create and also the listen address with it.
func NewAPIServer(ListenAddr string, store Storage) *APIServer {
	return &APIServer{
		ListenAddr: ListenAddr,
		store:      store,
	}
}

// This is Run function which will be having all the other component to run the server such as the routes and other data base initialization
func (s *APIServer) Run() {
	//New Router is a function of mux which will give us the router that we are storing in to a variable..
	router := mux.NewRouter()
	//this is a route that mathes the URL and passes it to the function(handlere ) which takes care of the activity to be done on it
	//here we are taking it as "s." because it is taking the APIServer as COntext and will have all the access to the memebers of the struct

	//and we wrapped th s.HandleAccount up because it is returning error and we want to return http.HandleFunc()
	router.HandleFunc("/account", MakeHTTPHandlerFunc(s.HandleAccount))
	router.HandleFunc("/account/{id}", MakeHTTPHandlerFunc(s.HandleGetAccount))

	log.Println("JSON Api running on port: ", s.ListenAddr)

	http.ListenAndServe(s.ListenAddr, router)

}

// Account route which will handle the account section when hit with the above Endpoint
func (s *APIServer) HandleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.HandleGetAccount(w, r)
	}

	if r.Method == "DELETE" {
		return s.HandleDeleteAccount(w, r)
	}
	if r.Method == "POST" {
		return s.HandleCreateAccount(w, r)
	}

	return fmt.Errorf("invalid type of method %s", r.Method)
}

/// VARIOUS HANDLERS WHICH WILL BE CONTAINING THE LOGIC OF THE PROJECT


//we are making a create account request which will take the two parameteres of firstname and the lastname and we will Decode that from the r.Body
//how the decoder works is on main.go
func (s *APIServer) HandleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	CreateAccountReq := new(CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(CreateAccountReq); err != nil {
		return err
	}
	//now creating a new account by passing the necessary data all other are assigned to zero or respective values
	account := NewAccount(CreateAccountReq.FirstName, CreateAccountReq.LastName)
	if err := s.store.CreateAccount(account); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, account)

}

func (s *APIServer) HandleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) HandleGetAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]

	fmt.Println(id)
	return WriteJSON(w, http.StatusOK, &Account{})
}
