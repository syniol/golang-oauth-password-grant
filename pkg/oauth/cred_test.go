package oauth

import (
	"testing"
)

func TestNewCredentialPassword(t *testing.T) {
	t.Run("Given password is correct", func(t *testing.T) {
		actual, err := NewCredentialPassword("Johnspa$$word")
		if err != nil {
			t.Fatal("it was not expecting an error but got", err.Error())
		}

		isVerified := actual.VerifyPassword("Johnspa$$word")
		if isVerified != true {
			t.Error("it was expecting to be verified")
		}
	})

	t.Run("Given password is incorrect", func(t *testing.T) {
		actual, err := NewCredentialPassword("Johnspa$$word")
		if err != nil {
			t.Fatal("it was not expecting an error but got", err.Error())
		}

		isVerified := actual.VerifyPassword("Johnspa$$wordNotMatching")
		if isVerified == true {
			t.Error("it was expecting to not verify the incorrect password")
		}
	})

}

func TestCredentialVerifyPassword(t *testing.T) {
	creds, _ := NewCredentialPassword("johnspassword1!")

	t.Run("correct", func(t *testing.T) {
		ok := creds.VerifyPassword("johnspassword1!")
		if !ok {
			t.Error("it should have verified the password")
		}
	})

	t.Run("incorrect", func(t *testing.T) {
		ok := creds.VerifyPassword("johnspassword1!IncorrecT")
		if ok {
			t.Error("it should have not verified the password")
		}
	})
}

func TestCredentialVerifyPasswordFromRawValues(t *testing.T) {
	var sut *Credential

	sut = &Credential{
		PublicKey:      "MmQyZDJkMmQyZDQyNDU0NzQ5NGUyMDUwNTU0MjRjNDk0MzIwNGI0NTU5MmQyZDJkMmQyZDBhNGQ0MzZmNzc0MjUxNTk0NDRiMzI1Njc3NDE3OTQ1NDE2MTVhNTc2NjZlNmUzNzQxNmU0YzQ0NDY2MTRmMzM0NzMwNDQ2OTdhNzA1NzRhNDY2MjdhNzI3NjcyNDIzMTRhNjE3YTc1NzM0NjY5MzQ2ODY4Mzg2NzNkMGEyZDJkMmQyZDJkNDU0ZTQ0MjA1MDU1NDI0YzQ5NDMyMDRiNDU1OTJkMmQyZDJkMmQwYQ==",
		PrivateKey:     "MmQyZDJkMmQyZDQyNDU0NzQ5NGUyMDUwNTI0OTU2NDE1NDQ1MjA0YjQ1NTkyZDJkMmQyZDJkMGE0ZDQzMzQ0MzQxNTE0MTc3NDI1MTU5NDQ0YjMyNTY3NzQyNDM0OTQ1NDk0ZTMzNDY2NTU3NzM0ZjMzNTg1MDM2MzkzMzZmMzk0ZTY0NDE0NDM5NDMzODMwMzM2YzY1Njk3MTZjMmI2MTU4NTk3ODU0NDM2MjQyNDg1MDJmNzQ2YjBhMmQyZDJkMmQyZDQ1NGU0NDIwNTA1MjQ5NTY0MTU0NDUyMDRiNDU1OTJkMmQyZDJkMmQwYQ==",
		HashedPassword: "ZDRkM2QzMWQyN2JjZTYyZTRjODI2MTFkYmZjMzk0YmIzNTI4MmRhODMwYTBhMWI3NjBiZjhkZjQzOGZjZDViOTViMGI4ZDBjMTY5ZjlhMzAxNGIwMGY4ZDVlYTMyMWE5MDAzNzVhNGE0MWZhMTFhZDViNjEwYTg0YTk2ZTAyMDI=",
	}

	ok := sut.VerifyPassword("johnspassword1")
	if !ok {
		t.Error("password is not matching")
	}
}
