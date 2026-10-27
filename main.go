package main

import (
	"projet/routes"

	"github.com/gin-gonic/gin"
)

func main() {

    router := gin.Default()

    router.Static("/css", "./frontend/css")
    router.Static("/js", "./frontend/js")

    router.LoadHTMLGlob("frontend/templates/*")

	// Chargement des pages dont on a besoin
	routes.SetupRoutes(router)

	// Déclaration routes api
	router.POST("/api/connexion", routes.Connexion)
	router.POST("/api/inscription", routes.Inscription)
	router.GET("/api/scoretotal", routes.ScoreTotal)
	router.GET("/api/affichertheme", routes.AfficherTheme)
	router.GET("/api/scoretheme", routes.ScoreTheme)
	router.GET("/api/afficherquestion", routes.AfficherQuestion)
	

    router.Run(":8080")
}
