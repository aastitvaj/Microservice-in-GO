package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello")
		d, err := ioutil.ReadAll(r.Body)
		com := bytes.Compare(d, []byte("aas"))
		if err != nil || com == 0 {
			http.Error(rw, "Oops", http.StatusBadRequest)

			//rw.WriteHeader(http.StatusBadRequest)
			//rw.Write([]byte("Oops"))
			return
		}

		fmt.Fprintf(rw, "hello %s", d)
	})

	http.HandleFunc("/favicon.ico", func(rw http.ResponseWriter, r *http.Request) {
		http.Error(rw, "Oops", http.StatusBadRequest)
	})

	http.HandleFunc("/Goodbye", func(http.ResponseWriter, *http.Request) {
		fmt.Println("Goodbye")
	})
	http.ListenAndServe(":8080", nil)
}
