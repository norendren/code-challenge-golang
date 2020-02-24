package main

// http server that logs requests and allows clients to define response duration and returned status code.
//
// Examples:
//   # return a response after 10 minutes with a 500 status code
//   curl <host>/?duration=10m&statusCode=500
//
//   # return a response after 20 seconds with a 201 status code
//   curl <host>/?duration=10m&statusCode=500
//
//   # return a response immediately with a 200 status code (defaults)
//   curl <host>/

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//log.Print("echo server received a request")

	query := r.URL.Query()
	statusCode := query.Get("statusCode")
	duration := query.Get("duration")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//log.Printf("error reading body: %v", err)
		return
	}

	if len(body) > 0 {
		//log.Printf("body received: %v", string(body))
	}

	if duration != "" {
		d, err := time.ParseDuration(duration)
		if err != nil {
			msg := fmt.Sprintf("invalid duration: %v", duration)
			//log.Print(msg)
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(msg))
			return
		}
		//log.Printf("sleeping requested duration: %v", d)
		time.Sleep(d)
	}

	if statusCode != "" {
		s, err := strconv.Atoi(statusCode)
		if err != nil {
			msg := fmt.Sprintf("invalid statusCode: %v", statusCode)
			//log.Print(msg)
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(msg))
			return
		}
		w.WriteHeader(s)
        if s == 200 {
            log.Println("printing image")
            w.Write([]byte{137, 80, 78, 71, 13, 10, 26, 10, 0, 0, 0, 13, 73, 72, 68, 82, 0, 0, 0, 1, 0, 0, 0, 1, 8, 2, 0, 0, 0, 144, 119, 83, 222, 0, 0, 0, 12, 73, 68, 65, 84, 8, 215, 99, 216, 122, 225, 2, 0, 4, 147, 2, 86, 234, 93, 183, 214, 0, 0, 0, 0, 73, 69, 78, 68, 174, 66, 96, 130})
        }
		return
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	//log.Print("echo server started")

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
