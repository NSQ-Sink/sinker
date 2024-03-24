package http

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/retrieve-post-json", getRoot)
	http.HandleFunc("/retrieve-post-json", getHello)

	http.ListenAndServe(":8007", nil)
}

func getRoot(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "root\n")
}

func getHello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}
