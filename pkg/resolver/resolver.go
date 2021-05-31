package resolver

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/katallaxie/puzzler/pkg/header"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/sync/errgroup"
)

// Resovler ...
type Resolver interface {
	WithContext(context.Context, *goquery.Document) error
}

// ResolverFunc ...
type ResolverFunc = func(*goquery.Document) func() error

// New ...
func New() Resolver {
	return &resolver{}
}

type resolver struct {
}

// WithContext ...
func (r *resolver) WithContext(ctx context.Context, doc *goquery.Document) error {
	g, gctx := errgroup.WithContext(ctx)

	doc.Find("fragment").Each(func(i int, s *goquery.Selection) {
		src, exists := s.Attr("src")
		if !exists {
			return
		}

		g.Go(func() error {
			req, err := http.NewRequest("GET", src, nil)
			if err != nil {
				return err
			}

			req = req.WithContext(gctx)
			client := http.DefaultClient
			res, err := client.Do(req)
			if err != nil {
				return err
			}
			defer res.Body.Close()

			if res.StatusCode != http.StatusOK {
				return nil
			}

			// demo header
			res.Header.Set("Link", "<https://unpkg.com/react-dom@17/umd/react-dom.development.js>; rel=\"script\"; crossorigin=\"\"\"")

			l := header.Header(res.Header.Get("Link"))
			nodes := header.CreateNodes(l.Links())
			head := doc.Find("head")
			head.AppendNodes(nodes...)

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return err
			}

			// do not replace when not resolved
			s.ReplaceWithHtml(string(b))

			return nil
		})
	})

	// this is sync for now
	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}
