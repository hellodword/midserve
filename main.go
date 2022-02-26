package main

import (
	"flag"
	"net/http"
	"regexp"
)

func main() {
	flag.Parse()

	http.Handle("/", FileServer(Dir("."), []*regexp.Regexp{
		regexp.MustCompile(`^\.git`),
		regexp.MustCompile(`^\.vscode`),
		regexp.MustCompile(`^\.idea`),
	}))

	http.ListenAndServe(":8000", nil)

}
