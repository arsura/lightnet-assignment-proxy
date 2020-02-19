package main

import (
	"github.com/arsura/lightnet-assignment-proxy/router"
)

func main() {
	router := routes.Router()
	router.Run(":8082")
}
