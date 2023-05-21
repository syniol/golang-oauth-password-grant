package oauth

import (
	"testing"
)

func TestToken(t *testing.T) {
	token, _ := NewToken()

	if len(token.randomKey) != 12 {
		t.Error("not matching expected length 12")
	}

	t.Log(string(token.randomKey))
}
