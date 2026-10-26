package routes

import (
    "github.com/gin-gonic/gin"
    "strings"
    "projet/services"
)

func Inscription(c *gin.Context) {
    var requete struct {
        Pseudo string `json:"pseudo"`
    }

    if err := c.ShouldBindJSON(&requete); err != nil {
        c.JSON(400, gin.H{
            "message": "Requête invalide",
        })
        return
    }

    requete.Pseudo = strings.TrimSpace(requete.Pseudo)

    if requete.Pseudo == "" {
        c.JSON(400, gin.H{
            "message": "Le pseudo est obligatoire",
        })
        return
    }

    _, existe := services.Rechercher_sauvegarde(requete.Pseudo)

    if existe {
        c.JSON(409, gin.H{
            "message": "Ce pseudo est déjà utilisé",
        })
        return
    }

    id := services.Creer_sauvegarde(requete.Pseudo)

    c.JSON(201, gin.H{
        "id_sauvegarde": id,
        "pseudo":        requete.Pseudo,
    })
}