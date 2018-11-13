package main

import (
	"compress/gzip"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"
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
<body>
	<div class="flex-row">
		<nav class="pure-menu">
			<ul class="pure-menu-list">
				<li><a href="/" class="pure-menu-link">Home</a></li>
				<li><a href="/grid.html" class="pure-menu-link">Grid Layout</a></li>
				<li><a href="/flex.html" class="pure-menu-link">Flex Layout</a></li>
				<li class="pure-menu-item pure-menu-disabled">Dynamic</li>
			</ul>
		</nav>
		<section>
			<header>The Header</header>
`

	dynBody = `	<h1>Hello</h1>
			<p>received: {{ .Path }}
			<p>headers: {{ .Headers }}
`
	dynFoot = `
			<footer>Copyright &copy {{ .Year }} Beyond Broadcast LLP.
				All rights reserved.
			</footer>
		</section>
	</div>
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

func (h *dynHdlr) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	hdrs := strings.Join(r.Header["Accept-Encoding"], "|")
	var (
		w  io.Writer
		zw *gzip.Writer
	)
	w = wr
	if strings.Index(hdrs, "gzip") != -1 {
		zw = gzip.NewWriter(wr)
		defer zw.Close()
		wr.Header().Set("Content-Type", "text/html; charset=utf-8")
		wr.Header().Set("Content-Encoding", "gzip")
		w = zw
		//log.Println("dyn compr")
	}
	dHeadT.Execute(w, struct {
		Title string
	}{"dyn1"})

	dBodyT.Execute(w, struct {
		Path    string
		Headers string
	}{r.URL.Path, fmt.Sprintf("%s", r.Header)})

	dFootT.Execute(w, struct {
		Year string
	}{time.Now().Format("2006")})
}
