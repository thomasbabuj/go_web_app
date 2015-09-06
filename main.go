/**
*  Basic Server
*
 */

package main

import "net/http"
import "log"

func main() {
	http.HandleFunc("/", someFunc)
	http.HandleFunc("/test", testFunc)
	http.ListenAndServe(":8080", nil)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func someFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello Golang World!! :)"))
}

func testFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("This is s sample test route...."))
}

/*
http.HandleFunc("/", someFunction)

    Go matches requests to the most specific route registered

    http://golang.org/pkg/net/http/#ServeMux

    everything matches "/"

    nil

     - http.ListenAndServe(":8080", nil)
     - meaning: use the DefaultServeMux

      http://golang.org/pkg/net/http/#pkg-variables

    behind the scenes:

      - request comes in
      - received on primary thread
      - goroutine created

        -- runs concurrently to main thread

            -- lightweight

    - request passed to goroutine

    - handling multiple requests at same time: "multiplexing"
*/
