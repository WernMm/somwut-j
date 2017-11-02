// +build !appengine

import (
	"net/http"
	"somwut-j/myserver"
)

func main() {
	myserver.RegisterHandlers()
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":8080", nil)
}
