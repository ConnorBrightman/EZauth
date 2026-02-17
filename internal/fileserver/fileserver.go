package fileserver

import (
	"log"
	"net/http"
)

// ServePublic serves all files from ./public
func ServePublic() http.Handler {
	fs := http.FileServer(http.Dir("./public"))
	log.Println("✅ Serving static files from ./public")
	return fs
}

// ServePage maps a clean URL to a specific HTML file
func ServePage(path string, htmlFile string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/"+htmlFile)
	}
}
