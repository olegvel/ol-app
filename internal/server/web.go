package server

import (
	"github/olegvel/ol-app/web"
	"io"
	"io/fs"
	"net/http"
)

func NewWebServer() http.Handler {
	mux := http.NewServeMux()

	webFS := web.GetWebFS()
	staticFS, err := fs.Sub(webFS, "static")
	if err != nil {
		panic(err)
	}

	mux.Handle("/static/", http.StripPrefix("/static", http.FileServerFS(staticFS)))

	mux.HandleFunc("/", ServeEmbeddedHTMLFile(webFS, "index.html"))
	mux.HandleFunc("/media", ServeEmbeddedHTMLFile(webFS, "pages/media.html"))

	return mux
}

func ServeEmbeddedHTMLFile(fileSystem fs.FS, filename string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		file, err := fileSystem.Open(filename)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		stat, err := file.Stat()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		http.ServeContent(w, r, stat.Name(), stat.ModTime(), file.(io.ReadSeeker))
	}
}
