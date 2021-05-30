package puzzler

import (
	"net/http"
	"os"
	"path"

	"github.com/PuerkitoBio/goquery"
)

func handler(w http.ResponseWriter, req *http.Request) {
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
	r := &resolver{}

	if err := r.WithContext(ctx, doc); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	html, err := doc.Html()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write([]byte(html))
}

func handleFragment(i int, s *goquery.Selection) {
	_, exists := s.Attr("src")
	if !exists {
		return
	}

	s.SetAttr("src", "http://localhost")
}
