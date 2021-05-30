package handler

import (
	"net/http"
	"os"
	"path"

	"github.com/katallaxie/puzzler/pkg/resolver"

	"github.com/PuerkitoBio/goquery"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	cwd, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	f, err := os.Open(path.Join(cwd, "examples", "example.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer f.Close()

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	ctx := req.Context()
	r := resolver.New()

	if err := r.WithContext(ctx, doc); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	html, err := doc.Html()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write([]byte(html))
}
