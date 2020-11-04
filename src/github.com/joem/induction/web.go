package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {

	srv := startHttpServer()

	time.Sleep(time.Second * 10)

	srv.Shutdown(nil)
}

func startHttpServer() *http.Server {
	srv := &http.Server{Addr: ":8080"}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "hello world\n")
	})

	go func() {
		if err:=srv.ListenAndServe();err != nil {
			log.Printf("http.Server ListenAndServer() error : %s", err)
		}
	} ()
	return srv
}
