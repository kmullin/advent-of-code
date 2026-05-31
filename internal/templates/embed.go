package templates

import "embed"

type Data struct {
	Year, Day int
}

//go:embed *.tmpl
var FS embed.FS
