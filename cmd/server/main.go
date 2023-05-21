package main

import (
	"log"

	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/database"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error in load .env file")
	}

	database.NewDatabaseConnection()

	r := gin.Default()

	r.Static("/styles", "./public/styles")
	r.LoadHTMLGlob("./public/pages/*")
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	routes.NewApiRoutes(r)
	routes.NewWebRoutes(r)

	r.Run()
}
