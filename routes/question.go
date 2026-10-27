package routes

import (
    "github.com/gin-gonic/gin"
    "projet/services"
    "strconv"
)


func AfficherQuestion(c *gin.Context) {
    idParam := c.Query("id_sauvegarde")
    idSauvegarde, err := strconv.ParseInt(idParam, 10, 64)

    if err != nil || idSauvegarde <= 0 {
        c.JSON(400, gin.H{
            "message": "Identifiant de sauvegarde invalide",
        })
        return
    }

    idThemeParam := c.Query("id_theme")
    idTheme, err := strconv.ParseInt(idThemeParam, 10, 64)

    if err != nil || idSauvegarde <= 0 {
        c.JSON(400, gin.H{
            "message": "Identifiant de theme invalide",
        })
        return
    }

    question := services.Rechercher_question(idSauvegarde, idTheme)

    c.JSON(200, question)
}