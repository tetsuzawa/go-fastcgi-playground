package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		h := r.Header
		fmt.Printf("%+v", h)
		fmt.Println("stdout log")
		if _, err := fmt.Fprintln(os.Stderr, "stderr log"); err != nil {
			log.Fatalln(err)
		}
		if _, err := w.Write([]byte("response writer log")); err != nil {
			log.Fatalln(err)
		}
	})
	fmt.Println("fast cgi is listening on port 9001")
	l, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalln(err)
	}
	if err := fcgi.Serve(l, nil); err != nil {
		log.Fatalln(err)
	}
}
