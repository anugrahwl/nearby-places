package main

import (
	"log"

	"github.com/anugrahwl/nearby-places/models"
	"github.com/anugrahwl/nearby-places/router"
)

func main() {
	batchPlace := models.LoadAll()
	r := router.SetupRouter(batchPlace)

	log.Fatalln(r.Run())
}
