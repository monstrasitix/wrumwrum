package routing

import (
	"net/http"

	"sim/internal/web"
)

type ChartsHandler struct{}

func (h ChartsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getIndex(w)
	}
}

func (h ChartsHandler) getIndex(w http.ResponseWriter) {
	web.HTML["charts"].ExecuteTemplate(w, "base", nil)
}
