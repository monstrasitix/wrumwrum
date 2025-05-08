package routing

import (
	"net/http"

	"sim/internal/web"
)

type HomeHandler struct{}

func (h HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getIndex(w)
	}
}

func (h HomeHandler) getIndex(w http.ResponseWriter) {
	web.HTML["home"].ExecuteTemplate(w, "base", nil)
}
