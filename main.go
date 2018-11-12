package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("pure-css demo")
	if err := os.MkdirAll("files", 0700); err != nil {
		log.Fatalf("could not create files dir: %v", err)
	}
	http.Handle("/dyn/", &dynHdlr{})
	http.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		remember := r.FormValue("remember")
		email := r.FormValue("email")
		passwd := r.FormValue("passwd")
		fmt.Fprintf(w, "hello: email is %q, passwd is %q and remember is %v",
			email, passwd, remember)
	})
	http.HandleFunc("/uploadfile", func(w http.ResponseWriter, r *http.Request) {
		f, fh, err := r.FormFile("uploadfile")
		if err != nil {
			log.Printf("uploadfile: %v", err)
			fmt.Fprintf(w, "file upload error: %v", err)
			return
		}
		fmt.Fprintf(w, "file upload: %q of size %d", fh.Filename, fh.Size)
		of, err := os.Create(filepath.Join("files", fh.Filename))
		if err != nil {
			log.Printf("upload create file: %v", err)
		}
		defer of.Close()
		if _, err := io.Copy(of, f); err != nil {
			log.Printf("upload copy file: %v", err)
		}
	})

	http.Handle("/", gzHandler(http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServeTLS(":8080", "server.pem", "server-key.pem", nil))
}

type zh struct {
	h http.Handler
}

func (z *zh) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	z.h.ServeHTTP(gzRW(w), r)
}
func gzHandler(h http.Handler) http.Handler {
	z := &zh{h: h}
	return z
}

type zrw struct {
	in http.ResponseWriter
}

func (w *zrw) Header() http.Header {
	return w.in.Header()
}
func (w *zrw) Write(b []byte) (int, error) {
	zw := gzip.NewWriter(w.in)
	i, err := zw.Write(b)
	if err != nil {
		return i, err
	}
	defer zw.Close()
	return i, nil
}
func (w *zrw) WriteHeader(status int) {
	w.in.WriteHeader(status)
}
func gzRW(w http.ResponseWriter) http.ResponseWriter {
	zw := &zrw{in: w}
	return zw
}
