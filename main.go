package main

import (
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
	http.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		remember := r.FormValue("remember")
		email := r.FormValue("email")
		passwd := r.FormValue("passwd")
		fmt.Fprintf(w, "hello: email is %q, passwd is %q and remember is %v",
			email, passwd, remember)
	})
	http.Handle("/", http.FileServer(http.Dir("static")))
	log.Fatal(http.ListenAndServeTLS(":8080", "server.pem", "server-key.pem", nil))
}
