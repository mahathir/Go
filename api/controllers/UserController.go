package controllers

import (
	"fmt"
	"net/http"

	"three_tiers/business/interfacebusiness"
	"three_tiers/utilities"

	"github.com/julienschmidt/httprouter"
)

type (
	UserController struct {
		UserService interfacebusiness.IUserService
	}
)

// GetUser retrieves an individual user resource
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	u, err := uc.UserService.GetUser(id)

	// Fetch user
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	} else if u == nil {
		http.Error(w, "User not found", 404)
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := utilities.JsonEncode(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}
