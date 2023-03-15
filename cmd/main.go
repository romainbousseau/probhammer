package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/romainbousseau/probhammer/internal/server"
	"github.com/romainbousseau/probhammer/internal/storage"
	"github.com/romainbousseau/probhammer/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := utils.LoadConfig(".", ".env")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// set router
	router := gin.Default()

	// init DB
	db, err := utils.OpenDBConnection(config)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	err = utils.Migrate(db)
	if err != nil {
		log.Fatal("cannot migrate tables:", err)
	}

	// build server and run
	s := server.NewServer(storage.NewStorage(db), router)
	err = s.SetRoutesAndRun()
	if err != nil {
		log.Fatal("cannot set routes:", err)
	}
}
