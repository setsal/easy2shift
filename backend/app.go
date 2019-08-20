package main

import (
	"net/http"

	"./routes"
)

func main() {
	r := routes.NewRouter()
	print("[INFO] Start serving at localhost:3000\n")
	http.ListenAndServe(":3000", r)
}
