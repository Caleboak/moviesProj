package main

import (
	"MovieDatabases/handler"
)

func main() {
	svr := handler.NewServer()
	svr.ListenAndServe()

}
