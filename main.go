/**
*  Muxs
*
 */

package main

import "net/http"
import "log"

func main() {

	/*
	 * Using own mux instead of default
	 */
	myMux := http.NewServeMux()
	myMux.HandleFunc("/", someFunc)
	myMux.HandleFunc("/test", testFunc)
	http.ListenAndServe(":8080", myMux)
	log.Fatal(http.ListenAndServe(":8080", myMux))
}

func someFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello Golang World!! :)"))
}

func testFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("This is s sample test route...."))
}
