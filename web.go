package main

import (
	"fmt"
	"net/http"
	"os"
)

var dat string

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/data", data)
	fmt.Println("listening")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func hello(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, dat)
}

func data(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	dat = req.Form.Get("data")
	fmt.Fprintln(rw, "OK")
}
