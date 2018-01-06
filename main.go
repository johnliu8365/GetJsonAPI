package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	res, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintln(w, string("welcome"))
	fmt.Fprintln(w, string(res))
}

func main() {
	fmt.Println("HI")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
