package routing

import "net/http"

func StaticHandler() http.Handler {
	dir := http.Dir("./static")
	handler := http.FileServer(dir)

	return http.StripPrefix("/static/", handler)
}
