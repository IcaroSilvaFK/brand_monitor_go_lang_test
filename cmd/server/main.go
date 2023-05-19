package main

import (
	"log"

	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error in load .env file")
	}

	r := gin.Default()
	// html := template.Must(template.New("index.html").ParseFiles("./public/index.html"))

	r.LoadHTMLGlob("./public/*")
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	routes.InitializeWebRoutes(r)

	r.Run()
}
