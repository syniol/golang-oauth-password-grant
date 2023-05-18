package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"oauth-password/pkg/cache"
	"strconv"

	_ "github.com/lib/pq"

	"oauth-password/internal/clients"
	"oauth-password/pkg/oauth"
)

func NewServer() {
	sm := http.NewServeMux()

	sm.HandleFunc("/oauth2/clients", func(rw http.ResponseWriter, req *http.Request) {
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
		defer req.Body.Close()

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

		clientCred, err := oauth.NewClientCredential(cred, client.Username.String())
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(err.Error()))

			return
		}

		repo, err := clients.NewRepository(nil)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)

			// todo give generic error handler
			rw.Write([]byte(err.Error()))

			return
		}

		_, err = repo.InsertSingle(*clientCred)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)

			// todo generic error message
			rw.Write([]byte(err.Error()))

			return
		}

		resp := &oauth.ClientResponse{
			ClientID: clientCred.ClientID,
		}

		rw.Write(resp.Bytes(true))
	})

	sm.HandleFunc("/oauth2/token", func(rw http.ResponseWriter, req *http.Request) {
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

		repo, _ := clients.NewRepository(nil)

		userCredentialPassword, err := repo.FindByUsername(pgr.Username)
		if err != nil {
			_, _ = rw.Write([]byte(fmt.Sprintf("error: %s", err.Error())))

			return
		}

		isVerified := userCredentialPassword.Data.VerifyPassword(pgr.Password.String())
		if !isVerified {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(fmt.Sprintf("password: %s is invalid", pgr.Password.String())))

			return
		}

		// todo: Create a token and parse it to NewPasswordGrantResponse - No JWT! use the same algorithm used for encryption
		tokenMock := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIi"
		cacheService, _ := cache.NewCache(nil)
		_ = cacheService.Persist(strconv.Itoa(int(userCredentialPassword.ID)), tokenMock)

		resp := oauth.NewPasswordGrantResponse(tokenMock)

		_, _ = rw.Write([]byte(fmt.Sprintf("processed entry: %s", resp.String())))
	})

	err := http.ListenAndServe(":8080", sm)
	if err != nil {
		log.Fatal(err.Error())
	}
}
