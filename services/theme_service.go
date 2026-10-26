package services


import (
	"projet/models"
)


func Afficher_theme(v_id_save int64) []models.Theme{
	// récupération de l'id du theme de la sauvegarde et du nom par un left join

	liste_theme, err := Base_donnees().Query("SELECT st.id, t.nom_theme, st.theme_demarre FROM sauvegardeTheme st LEFT JOIN theme t ON st.theme_id = t.id WHERE st.sauvegarde_id = ?", v_id_save)
	GestionErreur(err)
	defer Base_donnees().Close()

	var v_return []models.Theme

	// affichage de chaque thème
	for liste_theme.Next() {
		var t models.Theme
		err := liste_theme.Scan(&t.Id, &t.Nom, &t.Demarre)
		GestionErreur(err)
		v_return = append(v_return, t)
	}
	return v_return
}
