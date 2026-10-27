package routes

import (
    "github.com/gin-gonic/gin"
    "projet/services"
)


func Restart_theme(c *gin.Context) {
    var donnees struct {
        IdSauvegarde int64 `json:"id_sauvegarde"`
        IdTheme      int64 `json:"id_theme"`
    }

    if err := c.ShouldBindJSON(&donnees); err != nil {
        c.JSON(400, gin.H{
            "message": "Données JSON invalides",
        })
        return
    }

    if donnees.IdSauvegarde <= 0 || donnees.IdTheme <= 0 {
        c.JSON(400, gin.H{
            "message": "Données invalides",
        })
        return
    }

    services.Recommencer_theme(donnees.IdSauvegarde, donnees.IdTheme)

    c.JSON(200, gin.H{
		"message": "Thème réinitialisé",
	})
}