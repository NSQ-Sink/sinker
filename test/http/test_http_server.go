package http

import "net/http"

func main() {
	http.HandleFunc("/retrieve-post-json", getRoot)
	http.HandleFunc("/retrieve-post-json", getHello)

	err := http.ListenAndServe(":8007", nil)
}
