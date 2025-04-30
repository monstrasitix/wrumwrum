package routing

import (
	"net/http"
	"sim/internal/web"
)

type LoginHandler struct{}

func (h LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getIndex(w)

	case http.MethodPost:
		h.postIndex(w, r)
	}
}

func (l LoginHandler) getIndex(w http.ResponseWriter) {
	web.HTML["login"].ExecuteTemplate(w, "base", nil)
}

func validCreds(email, password string) bool {
	return email == "someone@mail.com" && password == "123"
}

func (l LoginHandler) postIndex(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	if validCreds(email, password) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		web.HTML["login"].ExecuteTemplate(w, "base", nil)
	}
}
