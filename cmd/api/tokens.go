package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/SergioRosello/greenlight/internal/data"
	"github.com/SergioRosello/greenlight/internal/data/models"
	"github.com/SergioRosello/greenlight/internal/validator"
)

func (app *application) createAuthenticationTokenHandler(w http.ResponseWriter, r *http.Request) {

	// Parse the email and password from the request body
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Validate the email and password provided by the client

	v := validator.New()
	v.ValidateEmail(input.Email)
	v.ValidatePasswordPlaintext(input.Password)

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
	}

	// Lokup the user record based on the e-mail address. If no mathcing user
	// was found, then we call the app.imvalidCredentialsResponse() helper to
	// send a 401 Unauthorized response to the client.
	user, err := app.data.Users.GetByEmail(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrRecordNotFound):
			app.invalidCredentialsResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// Check if the provided password matches the actual password for the user.
	match, err := user.Password.Matches(input.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// If the passwords don't match, then we call the app.invalidCredentialsResponse()
	// helper again and return
	if !match {
		app.invalidCredentialsResponse(w, r)
		return
	}

	// Otherwise, if the password is correct, we generate a new token with a 24-hour
	// expiry time and scope 'authentication'.
	token, err := app.data.Tokens.New(user.ID, 24*time.Hour, data.ScopeAuthentication)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// Encode the token to JSON and send it in the reponse along with a 201 created
	// status code.
	err = app.writeJSON(w, http.StatusCreated, envelope{"authentication_token": token}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
