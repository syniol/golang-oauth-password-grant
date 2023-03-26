package internal

import (
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

		client, err := oauth.NewClient(reqBody)

		rw.Write([]byte(client.String()))
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

		_, err = oauth.NewPasswordGrantRequestWithForm(req.Form)
		if err != nil {
			_, _ = rw.Write([]byte(fmt.Sprintf("error: %s", err.Error())))

			return
		}

		// TODO create a repository service to verify client credentials

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
