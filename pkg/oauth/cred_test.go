package oauth

import "testing"

func TestNewCredentialPassword(t *testing.T) {
	t.Run("Given password is correct", func(t *testing.T) {
		actual, err := NewCredentialPassword("Johnspa$$word")
		if err != nil {
			t.Fatal("it was not expecting an error but got", err.Error())
		}

		isVerified := actual.PasswordVerify("Johnspa$$word")
		if isVerified != true {
			t.Error("it was expecting to be verified")
		}
	})

	t.Run("Given password is incorrect", func(t *testing.T) {
		actual, err := NewCredentialPassword("Johnspa$$word")
		if err != nil {
			t.Fatal("it was not expecting an error but got", err.Error())
		}

		isVerified := actual.PasswordVerify("Johnspa$$wordNotMatching")
		if isVerified == true {
			t.Error("it was expecting to not verify the incorrect password")
		}
	})

}
