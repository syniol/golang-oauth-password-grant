package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"oauth-password/internal/clients"
	"oauth-password/pkg/cache"
	"oauth-password/pkg/oauth"
)

func NewServer() {
	sm := http.NewServeMux()

	sm.HandleFunc("/healthz", func(wr http.ResponseWriter, req *http.Request) {
		_, _ = wr.Write([]byte("ok"))
	})

	sm.HandleFunc("/oauth2/clients", func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			rw.WriteHeader(http.StatusNotFound)

			return
		}

		var clientRequest *oauth.ClientRequest
		err := json.NewDecoder(req.Body).Decode(clientRequest)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			log.Println(err)

			_, _ = rw.Write([]byte("Unable to parse request body, please ensure request body is not malformed"))

			return
		}

		if err = clientRequest.Password.Validate(); err != nil {
			rw.WriteHeader(http.StatusBadRequest)

			_, _ = rw.Write([]byte(err.Error()))

			return
		}

		if err = clientRequest.Username.Validate(); err != nil {
			rw.WriteHeader(http.StatusBadRequest)

			_, _ = rw.Write([]byte(err.Error()))

			return
		}

		cred, err := oauth.NewCredentialPassword(clientRequest.Password.String())
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			_, _ = rw.Write([]byte(err.Error()))

			return
		}

		clientCred, err := oauth.NewClientCredential(cred, clientRequest.Username.String())
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			_, _ = rw.Write([]byte(err.Error()))

			return
		}

		repo, err := clients.NewRepository(nil)
		if err != nil {
			log.Println(err.Error())

			rw.WriteHeader(http.StatusInternalServerError)
			_, _ = rw.Write([]byte("error establishing database connection"))

			return
		}

		_, err = repo.InsertSingle(*clientCred)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)

			_, _ = rw.Write([]byte(err.Error()))

			return
		}

		resp := &oauth.ClientResponse{
			ClientID: clientCred.ClientID,
		}

		_, _ = rw.Write(resp.Bytes(true))
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
			rw.WriteHeader(http.StatusBadRequest)
			_, _ = rw.Write([]byte(fmt.Sprintf("error: %s", err.Error())))

			return
		}

		ctx := context.Background()

		repo, _ := clients.NewRepository(ctx)

		userCredentialPassword, err := repo.FindByUsername(pgr.Username)
		if err != nil {
			_, _ = rw.Write([]byte(fmt.Sprintf("error: %s", err.Error())))

			return
		}

		isVerified := userCredentialPassword.Data.VerifyPassword(pgr.Password.String())
		if !isVerified {
			rw.WriteHeader(http.StatusBadRequest)
			_, _ = rw.Write([]byte(fmt.Sprintf("password: %s is invalid", pgr.Password.String())))

			log.Println(err.Error())

			return
		}

		tokeniser, err := oauth.NewToken()
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			_, _ = rw.Write([]byte("unexpected error for token creation service"))

			log.Println(err.Error())

			return
		}

		token := string(tokeniser.Sign())

		cacheService, err := cache.NewCache(ctx)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			_, _ = rw.Write([]byte("unexpected error for token storage service"))

			log.Println(err.Error())

			return
		}

		err = cacheService.Persist(
			userCredentialPassword.Data.ClientID,
			token,
		)

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			_, _ = rw.Write([]byte("unexpected error for token storage service"))

			log.Println(err.Error())

			return
		}

		_, _ = rw.Write(oauth.
			NewPasswordGrantResponse(token).
			Byte(true),
		)
	})

	err := http.ListenAndServe(":8080", sm)
	//err := http.ListenAndServeTLS(":80", "server.crt", "server.key", sm)
	if err != nil {
		log.Fatal(err.Error())
	}
}
