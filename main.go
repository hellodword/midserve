package main

import (
	"flag"
	"net/http"
)

func main() {
	flag.Parse()

	http.Handle("/", FileServer(http.Dir(".")))

	http.ListenAndServe(":8000", nil)

}
