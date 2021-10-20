package main

import (
	"flag"
	"fmt"
	"github.com/DmitriyBelousov/kek/processor"
	"net/http"
)

var (
	node = "NBZY3YSA7BVLOUH4C5H4RXBQCSVARWW4QV6R55WXRXZPM4MJGCNU4JSF"
	url = "0.0.0.0:4222"
	replicas = 3
)

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "set up port")
	flag.Parse()

	fmt.Println("hello kek")

	proc, err := processor.NewProcessor(node,url,replicas)
	if err != nil {
		panic(err)
	}

	proc.Start()

	http.HandleFunc("/", hello)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.UserAgent())
	w.WriteHeader(200)
	_, err :=fmt.Fprintf(w, "hello kek")
	if err != nil {
		fmt.Println(err)
	}
}

