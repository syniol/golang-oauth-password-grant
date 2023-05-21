package internal

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"oauth-password/internal/clients"
	"oauth-password/pkg/cache"
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
			rw.Write([]byte("error establishing database connection"))

			log.Println(err.Error())

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
			rw.Write([]byte(fmt.Sprintf("error parsing form request: %s", err.Error())))

			return
		}

		pgr, err := oauth.NewPasswordGrantRequestWithForm(req.Form)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(fmt.Sprintf("error: %s", err.Error())))

			return
		}

		ctx := context.Background()

		repo, _ := clients.NewRepository(ctx)

		userCredentialPassword, err := repo.FindByUsername(pgr.Username)
		if err != nil {
			rw.Write([]byte(fmt.Sprintf("error: %s", err.Error())))

			return
		}

		isVerified := userCredentialPassword.Data.VerifyPassword(pgr.Password.String())
		if !isVerified {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(fmt.Sprintf("password: %s is invalid", pgr.Password.String())))

			log.Println(err.Error())

			return
		}

		tokeniser, err := oauth.NewToken()
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("unexpected error for token creation service"))

			log.Println(err.Error())

			return
		}

		token := string(tokeniser.Sign())

		cacheService, err := cache.NewCache(ctx)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("unexpected error for token storage service"))

			log.Println(err.Error())

			return
		}

		err = cacheService.Persist(
			userCredentialPassword.Data.ClientID,
			token,
		)

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("unexpected error for token storage service"))

			log.Println(err.Error())

			return
		}

		rw.Write(oauth.
			NewPasswordGrantResponse(token).
			Byte(true),
		)
	})

	err := http.ListenAndServe(":8080", sm)
	if err != nil {
		log.Fatal(err.Error())
	}
}
