package oauth

import "testing"

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

	act, _ := NewClientCredential(creds, "johndoe1")
	ok := act.VerifyPassword("johnspassword1!")

	if !ok {
		t.Fatal("sssssssss")
	}
}

func TestCredentialVerifyPasswordFromDatabase(t *testing.T) {
	var sut *Credential

	sut = &Credential{
		PublicKey:      "MmQyZDJkMmQyZDQyNDU0NzQ5NGUyMDUwNTU0MjRjNDk0MzIwNGI0NTU5MmQyZDJkMmQyZDBhNGQ0MzZmNzc0MjUxNTk0NDRiMzI1Njc3NDE3OTQ1NDE2NTJmNGM3NDc4Mzc3NTRlN2E0ZTQ4NTQ1MTZhMzM0NTM0NzczMzRiNTA0ODY3NzA0NTU1MzM2YTZlNzI2MTQ2NTA2NjcxNjQ0NjUzNTY1NTM2NzczODNkMGEyZDJkMmQyZDJkNDU0ZTQ0MjA1MDU1NDI0YzQ5NDMyMDRiNDU1OTJkMmQyZDJkMmQwYQ==",
		PrivateKey:     "MmQyZDJkMmQyZDQyNDU0NzQ5NGUyMDUwNTI0OTU2NDE1NDQ1MjA0YjQ1NTkyZDJkMmQyZDJkMGE0ZDQzMzQ0MzQxNTE0MTc3NDI1MTU5NDQ0YjMyNTY3NzQyNDM0OTQ1NDk0MjYxNDk0YTU5NDkzNjZjNGU2NjMwMzU1OTY2NjE0ODM2NzE3Nzc3MzU2MTMwNDk1OTU5NDU1YTc5NDY1OTc4NDkzNjRkNzY1NTczNzA2ZTRkNmEzMjBhMmQyZDJkMmQyZDQ1NGU0NDIwNTA1MjQ5NTY0MTU0NDUyMDRiNDU1OTJkMmQyZDJkMmQwYQ==",
		HashedPassword: "MjhmZjY0M2I1ZjQzNDlmZjg5NTVkNjE5NDlmOWRlYTJhZDU5M2M5OGEyZThhNzY0ZDNkNTdjMjczZjZiNGRlNGFhNDJkMTY3ZjU5YzVkODZhNTEzNDQ0YTY1OGJjMTNiZDkwM2FjNTYxOGNjMzFhZmFmMTNlYzZlNGFlMWQ2MDg=",
	}

	ok := sut.VerifyPassword("johnspassword1!")

	if !ok {
		t.Error("password is not matching")
	}
}
