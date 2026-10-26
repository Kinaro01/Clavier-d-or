package routes

import (
    "github.com/gin-gonic/gin"
    "strings"
    "projet/services"
)


func Connexion(c *gin.Context) {
    var requete struct {
        Pseudo string `json:"pseudo"`
    }

    if err := c.ShouldBindJSON(&requete); err != nil {
        c.JSON(400, gin.H{
            "message": "Requête invalide",
        })
        return
    }

    if strings.TrimSpace(requete.Pseudo) == "" {
        c.JSON(400, gin.H{
            "message": "Le pseudo est obligatoire",
        })
        return
    }

    id, connexion := services.Rechercher_sauvegarde(requete.Pseudo)

    if !connexion {
        c.JSON(404, gin.H{
            "message": "Pseudo introuvable",
        })
        return
    }

    c.JSON(200, gin.H{
        "id_sauvegarde": id,
        "pseudo":        requete.Pseudo,
    })
}