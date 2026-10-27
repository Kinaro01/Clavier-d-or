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

func VerifierQuestion(c *gin.Context) {
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

	var requete struct {
		ChoixReponse string `json:"choix_reponse"`
	}

	if err := c.ShouldBindJSON(&requete); err != nil {
        c.JSON(400, gin.H{
            "message": "Données invalides",
        })
        return
    }

	if requete.ChoixReponse == "" {
        c.JSON(400, gin.H{
            "message": "Réponse manquante",
        })
        return
    }

	question := services.Rechercher_question(idSauvegarde, idTheme)

	if question == nil {
        c.JSON(404, gin.H{
            "message": "Aucune question disponible",
        })
        return
    }

	resultat := services.Resultat_question(requete.ChoixReponse, *question, idSauvegarde,)

	c.JSON(200, gin.H{
        "correct": resultat,
    })
}