package oauth

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestNewPasswordGrantRequest(t *testing.T) {
	t.Run("valid_request_success", func(t *testing.T) {
		mockJsonMap := map[string]interface{}{
			"grant_type": "password",
			"username":   "joker88",
			"password":   "BallBagHarley",
		}

		mockJson, err := json.Marshal(mockJsonMap)
		if err != nil {
			t.Fatal(err)
		}

		actual, err := NewPasswordGrantRequest(mockJson)
		if err != nil {
			t.Fatalf("it was not expecting an error but got %s", err.Error())
		}

		if actual.GrantType != GrantTypePassword {
			t.Errorf(
				"it was expecting a grant type %s but got %s",
				GrantTypePassword,
				actual.GrantType,
			)
		}
	})

	t.Run("invalid request grant_type", func(t *testing.T) {
		mockJsonMap := map[string]interface{}{
			"grant_type": "pass",
			"username":   "joker88",
			"password":   "BallBagHarley",
		}

		mockJson, err := json.Marshal(mockJsonMap)
		if err != nil {
			t.Fatal(err)
		}

		_, err = NewPasswordGrantRequest(mockJson)
		if err == nil {
			t.Fatal("it was expecting an error")
		}

		if strings.Contains(err.Error(), "incorrect grant type") {
			t.Errorf("it was expecting error for incorrect grant type but got %s", err.Error())
		}
	})
}
