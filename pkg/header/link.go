package header

import (
	"strings"
)

// Link ...
type Link struct {
	URL    string
	Rel    string
	Params map[string]string
}

// Header ...
type Header string

// Link ...
func (s Header) Links() []Link {
	links := make([]Link, 0)

	for _, chunk := range strings.Split(string(s), ",") {

		l := Link{URL: "", Rel: "", Params: make(map[string]string)}

		for _, part := range strings.Split(chunk, ";") {
			part = strings.Trim(part, " ")
			if part == "" {
				continue
			}
			if part[0] == '<' && part[len(part)-1] == '>' {
				l.URL = strings.Trim(part, "<>")
				continue
			}

			key, val := parseParam(part)
			if key == "" {
				continue
			}

			if strings.ToLower(key) == "rel" {
				l.Rel = val

				continue
			}
			l.Params[key] = val
		}

		if l.URL != "" {
			links = append(links, l)
		}
	}

	return links
}

// FilterByStylesheet ...
func FilterByStylesheet(links ...Link) []Link {
	return FilterByRel(links, "stylesheet")
}

// FilterByStylesheet ...
func FilterByScript(links ...Link) []Link {
	return FilterByRel(links, "script")
}

// FilterByRel ...
func FilterByRel(links []Link, rel string) []Link {
	ll := make([]Link, 0)

	for _, l := range links {
		if l.Rel != rel {
			continue
		}

		ll = append(ll, l)
	}

	return ll
}

func parseParam(raw string) (key, val string) {

	parts := strings.SplitN(raw, "=", 2)
	if len(parts) == 1 {
		return parts[0], ""
	}
	if len(parts) != 2 {
		return "", ""
	}

	key = parts[0]
	val = strings.Trim(parts[1], "\"")

	return key, val

}
