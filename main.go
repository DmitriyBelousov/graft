package main

import (
	"fmt"
	"net/http"
)

func main(){
	fmt.Println("hello kek")

	http.HandleFunc("/", hello)
	http.ListenAndServe(":8090", nil)
}


func hello(w http.ResponseWriter, req *http.Request){
	fmt.Println(req.UserAgent())
	w.WriteHeader(200)
	fmt.Fprintf(w, "hello kek")
}
