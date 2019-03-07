package settings

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func GetServerHandler(r *mux.Router) (string, http.Handler) {
	port := os.Getenv("LANA_PORT")
	fmt.Println("Starting server at port", port)
	return ":" + port, r
}
