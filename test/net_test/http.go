package main

import "net/http"

func main2() {

	http.HandleFunc("/hello", helloHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic("error:" + err.Error())
	}
}

var helloHandler = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
	return
}
