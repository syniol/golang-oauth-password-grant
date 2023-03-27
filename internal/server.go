package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"oauth-password/pkg/oauth"
)

func NewServer() {
	sm := http.NewServeMux()

	sm.HandleFunc("/oauth/clients", func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			rw.WriteHeader(http.StatusNotFound)

			return
		}

		reqBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(err.Error()))

			return
		}

		client, err := oauth.NewClientRequest(reqBody)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(err.Error()))

			return
		}

		cred, err := oauth.NewCredentialPassword(client.Password.String())
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(err.Error()))

			return
		}

		credOut, _ := json.MarshalIndent(cred, "", "\t")

		// Todo persis the data in main database (Postgres)

		rw.Write(credOut)
	})

	sm.HandleFunc("/oauth/token", func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			rw.WriteHeader(http.StatusNotFound)

			return
		}

		err := req.ParseForm()
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			_, _ = rw.Write([]byte(fmt.Sprintf("error parsing form request: %s", err.Error())))

			return
		}

		pgr, err := oauth.NewPasswordGrantRequestWithForm(req.Form)
		if err != nil {
			_, _ = rw.Write([]byte(fmt.Sprintf("error: %s", err.Error())))

			return
		}

		// Todo FindCredentialPasswordByUsername repository method
		// Todo Connect to main database (Postgres to find a CredentialPassword from database)
		userCredentialPassword, err := oauth.FindCredentialPasswordByUsername(pgr.Username.String())
		if err != nil {
			_, _ = rw.Write([]byte(fmt.Sprintf("error: %s", err.Error())))

			return
		}

		userCredentialPassword.VerifyPassword(pgr.Password.String())
		// Todo Create a token and parse it to NewPasswordGrantResponse - No JWT! use the same algorithm used for encryption
		// Todo persist a token with client id as a ref for Cache storage with 1 Hour expiry (Redis)

		resp := oauth.NewPasswordGrantResponse(
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIi",
		)

		_, _ = rw.Write([]byte(fmt.Sprintf("processed entry: %s", resp.String())))
	})

	err := http.ListenAndServe(":8080", sm)
	if err != nil {
		log.Fatal(err.Error())
	}
}
