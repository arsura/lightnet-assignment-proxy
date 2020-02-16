package main

import (
	"github.com/arsura/lightnet-assignment-proxy/routes"
)

func main() {
	router := routes.Router()
	router.Run(":8082")
}
