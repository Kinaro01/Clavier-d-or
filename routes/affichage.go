package routes

import "github.com/gin-gonic/gin"


func SetupRoutes(router *gin.Engine) {

    router.GET("/", func(c *gin.Context) {
        c.HTML(200, "connexion.html", nil)
    })

    router.GET("/inscription", func(c *gin.Context) {
        c.HTML(200, "inscription.html", nil)
    })

    router.GET("/accueil", func(c *gin.Context) {
        c.HTML(200, "accueil.html", nil)
    })

    router.GET("/question", func(c *gin.Context) {
        c.HTML(200, "question.html", nil)
    })

}