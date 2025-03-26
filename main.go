package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	endpoint := flag.String("endpoint", "", "endpoint to call")
	flag.Parse()
	if err := http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request")
		if *endpoint != "" {
			r, _ := http.NewRequest("GET", *endpoint, nil)
			rs, _ := http.DefaultClient.Do(r)
			b, _ := io.ReadAll(rs.Body)
			fmt.Println(string(b))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello world"))
	})); err != nil {
		panic(err)
	}
	time.Sleep(time.Hour)
}
