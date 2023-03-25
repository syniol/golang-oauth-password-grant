package internal

import (
	"fmt"
	"net/http"
	"oauth-password/pkg/oauth"
)

func NewServer() {
	sm := http.NewServeMux()

	sm.HandleFunc("/oauth/token", func(rw http.ResponseWriter, req *http.Request) {
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
			rw.Write([]byte(fmt.Sprintf("error: %s", err.Error())))

			return
		}

		rw.Write([]byte(fmt.Sprintf("processed entry: %s", pgr.ToString())))
	})

	http.ListenAndServe(":8080", sm)
}
