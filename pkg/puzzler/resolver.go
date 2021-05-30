package puzzler

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/sync/errgroup"
)

// Resovler ...
type Resolver interface{}

type resolver struct {
}

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

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return err
			}

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
