package validator

import "github.com/SergioRosello/greenlight/internal/data"

func (v *Validator) ValidateEmail(email string) {
	v.Check(email != "", "email", "must be provided")
	v.Check(Matches(email, EmailRX), "email", "must be a valid email address")
}

func (v *Validator) ValidatePasswordPlaintext(password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 8, "password", "must be at least 8 bytes long")
	v.Check(len(password) <= 72, "password", "must not be more than 72 bytes long")
}

// The drawback of requiring all fields we want to validate, is that we have to
// allow the validator to access all the fields we have to validate.
// In this instance, we are forced to make accessible the user plaintext password
// and hashed version of such.
func (v *Validator) ValidateUser(user *data.User) {
	v.Check(user.Name != "", "name", "must be provided")
	v.Check(len(user.Name) <= 500, "name", "must not be more than 500 bytes long")

	// Call the standalone ValidateEmail() helper.
	v.ValidateEmail(user.Email)

	// If the plaintext password is not nil, call the standalone
	// ValidatePasswordPlaintext() helper.
	if user.Password.Plaintext != nil {
		v.ValidatePasswordPlaintext(*user.Password.Plaintext)
	}

	// If the password hash is ever nil, this will be due to a logic error in our
	// codebase (probably because we forgot to set a password for the user). It's a
	// useful sanity check to include here, but it's not a problem with the data
	// provided by the client. So rather than adding an error to the validation map we
	// raise a panic instead.
	if user.Password.Hash == nil {
		panic("missing password hash for user")
	}
}
