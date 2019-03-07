package main

import (
	"fmt"
	"lana/settings"
	"net/http"

	"lana/api"
)

func main() {
	fmt.Println("Lana Go challenge")
	a := api.API{}
	a.RegisterRoutes()

	s, h := settings.GetServerHandler(a.Router)
	http.ListenAndServe(s, h)
}
