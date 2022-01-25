package main

import (
	"net/http"
	"raffle/router"
)

func main() {
	http.HandleFunc("/", router.Lottery)
	http.HandleFunc("/delete", router.Delete)
	http.ListenAndServe(":8080", nil)
}
