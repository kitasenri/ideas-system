package controllers_common

import (
	"encoding/base64"
	"net/http"
	"strings"

	"src/consts"

	"github.com/zenazn/goji/web"
)

//---------------------------------------------------------------------
// Property
//---------------------------------------------------------------------

//---------------------------------------------------------------------
// Method
//---------------------------------------------------------------------
// Check Basic Auth
func CheckBasicAuth(c *web.C, h http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {

		auth := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Basic ") {
			pleaseAuth(w)
			return
		}

		password, error := base64.StdEncoding.DecodeString(auth[6:])
		if error != nil || string(password) != consts.BASIC_AUTH_PASS {
			pleaseAuth(w)
			return
		}

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// Please auth response
func pleaseAuth(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", `Basic realm="Gritter"`)
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("(´・ω・｀)"))
}
