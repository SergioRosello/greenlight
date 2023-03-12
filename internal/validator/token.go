package validator

import ()

// Check that the plaintext token has been provided and is exactly 26 bytes long.
func (v *Validator) ValidateTokenPlaintext(tokenPlaintext string) {
	v.Check(tokenPlaintext != "", "token", "must be provided")
	v.Check(len(tokenPlaintext) == 26, "token", "must be 26 bytes long")
}
