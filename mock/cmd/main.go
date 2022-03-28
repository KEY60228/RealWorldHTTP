package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>Hello</body></html>\n")
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	log.Println("start http listening :50510")
	httpServer.Addr = ":50510"
	log.Println(httpServer.ListenAndServe())
}
