package routes

import (
    "github.com/gin-gonic/gin"
    "projet/services"
    "strconv"
)

func AfficherTheme(c *gin.Context) {
    idParam := c.Query("id_sauvegarde")

    idSauvegarde, err := strconv.ParseInt(idParam, 10, 64)
    if err != nil || idSauvegarde <= 0 {
        c.JSON(400, gin.H{
            "message": "Identifiant de sauvegarde invalide",
        })
        return
    }

    theme := services.Afficher_theme(idSauvegarde)
    
	c.JSON(200, theme)
}