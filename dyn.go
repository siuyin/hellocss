package main

import (
	"html/template"
	"net/http"
	"time"
)

const (
	dynHead = `<!DOCTYPE html>
<html>
<head>
<title>{{ .Title }}</title>
<link href="/css/pure-min.css" rel="stylesheet">
<link href="/css/app.css" rel="stylesheet">
</head>
<header>The Header</header>
<body>
`

	dynBody = `<h1>Hello</h1>
<p>received: {{ .Path }}
`
	dynFoot = `
<footer>Copyright &copy {{ .Year }} Beyond Broadcast LLP.
	All rights reserved.
</footer>
</body>
</html>`
)

var dHeadT, dBodyT, dFootT *template.Template

func init() {
	dHeadT = template.Must(template.New("dynHead").Parse(dynHead))
	dBodyT = template.Must(template.New("dynBody").Parse(dynBody))
	dFootT = template.Must(template.New("dynFoot").Parse(dynFoot))
}

type dynHdlr struct{}

func (h *dynHdlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dHeadT.Execute(w, struct {
		Title string
	}{"dyn1"})

	dBodyT.Execute(w, struct {
		Path string
	}{r.URL.Path})

	dFootT.Execute(w, struct {
		Year string
	}{time.Now().Format("2006")})
}
