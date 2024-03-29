package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var solved bool = false

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("flag{1234}"))
}

func execHandler(w http.ResponseWriter, r *http.Request) {
	solved = true
	w.Write([]byte("OK"))
}

func execGetterHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(strconv.FormatBool(solved)))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/exec", execHandler)
	http.HandleFunc("/solved", execGetterHandler)
	fmt.Println("Server running...")
	http.ListenAndServe(":12345", nil)
}
