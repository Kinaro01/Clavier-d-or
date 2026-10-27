package routes

import (
    "github.com/gin-gonic/gin"
    "projet/services"
    "strconv"
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

func ExportPDF(c *gin.Context) {
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

    if err != nil || idTheme <= 0 {
        c.JSON(400, gin.H{
            "message": "Identifiant de theme invalide",
        })
        return
    }

    contenuPDF, err := services.Exporter_PDF(idSauvegarde, idTheme)
    if err != nil {
        c.JSON(500, gin.H{"erreur": "Impossible de générer le PDF"})
        return
    }

    c.Header("Content-Disposition", `attachment; filename="resultats_quizz.pdf"`)
    c.Data(200, "application/pdf", contenuPDF)
}