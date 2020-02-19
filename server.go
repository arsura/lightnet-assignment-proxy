package main

import (
	"github.com/arsura/lightnet-assignment-proxy/router"
)

func main() {
	router := router.Router()
	router.Run(":8082")
}
