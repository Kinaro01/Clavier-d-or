package services


import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
	"projet/models"

	_ "modernc.org/sqlite"
)

func chemin_projet() string {
	// récupèration du chemin du projet dynamiquement
	chemin, err := os.Getwd()
	if err != nil {
		fmt.Println("Impossible de trouver le chemin du projet")
		log.Fatal(err)
	}
	return chemin
}

func Base_donnees() *sql.DB {
	// renvoie la base + complétion du chemin avec le répértoire de la base de données
	db, erreur_db := sql.Open("sqlite", chemin_projet()+"/db/sql.db")
	if erreur_db != nil {
		log.Fatal(erreur_db, "Impossible de se connecter à la base")
	}
	return db
}

func GestionErreur(err error) {
	if err != nil {
		panic(err)
	}
}

func Creer_sauvegarde(v_pseudo string) int64 {

	// démarrer une transaction
	tx, err := Base_donnees().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
		Base_donnees().Close()
	}()

	// initialisation de la structure
	var s models.Sauvegarde
	s.Pseudo = v_pseudo
	s.Date = time.Now().Format(time.DateOnly)
	s.Score_total = 0

	// insertion des éléments dans la base de données (table sauvegarde)
	resultat_sauvegarde, err := tx.Exec("INSERT INTO sauvegarde (pseudo, date, score_total) VALUES (?, ?, ?)", s.Pseudo, s.Date, s.Score_total)
	GestionErreur(err)

	// récupération de l'id de la sauvegarde actuelle + des id de chaque thème
	s.Id, err = resultat_sauvegarde.LastInsertId()

	// initalisation des champs de la sauvegardeTheme
	var saveTheme models.SauvegardeTheme
	saveTheme.Theme_demarre = false

	row_theme, err := tx.Query("SELECT id FROM theme")
	GestionErreur(err)

	// pour chaque ligne de thème, insérer dans la table sauvegardeTheme
	for row_theme.Next() {
		err := row_theme.Scan(&saveTheme.Theme_id)
		GestionErreur(err)

		resultat_theme, err := tx.Exec("INSERT INTO sauvegardeTheme (sauvegarde_id, theme_id, theme_demarre) VALUES (?, ?, ?)", s.Id, saveTheme.Theme_id, saveTheme.Theme_demarre)
		GestionErreur(err)

		// récupération du dernier id automatique
		saveTheme.Id, err = resultat_theme.LastInsertId()
		GestionErreur(err)

		// initialisation de la structure de sauvegarde de questions
		var saveQuestion models.SauvegardeQuestion
		saveQuestion.Complete = false
		saveQuestion.Score_question = 0

		// récupération des id des questions du thème en cours
		row_question, err := tx.Query("SELECT id FROM questions WHERE theme_id = ?", saveTheme.Theme_id)
		GestionErreur(err)

		// pour chaque ligne de question, insérer dans la table sauvegardeQuestion
		for row_question.Next() {
			err := row_question.Scan(&saveQuestion.Question_id)
			GestionErreur(err)
			resultat_question, err := tx.Exec("INSERT INTO sauvegardeQuestion (sauvegarde_id_theme, sauvegarde_id, question_id, complete, score_question) VALUES (?, ?, ?, ?, ?)", saveTheme.Id, s.Id, saveQuestion.Question_id, saveQuestion.Complete, saveQuestion.Score_question)
			GestionErreur(err)

			_, err = resultat_question.LastInsertId()
			GestionErreur(err)
		}
	}
	return s.Id
}

func Charger_sauvegarde(v_pseudo string) int64 {
	// récupération de l'id de la sauvegarde car elle est liée a toutes les tables + return de l'id
	sauvegarde_en_cours := Base_donnees().QueryRow("SELECT id FROM sauvegarde WHERE sauvegarde.pseudo = ?", v_pseudo)

	var id int64
	err := sauvegarde_en_cours.Scan(&id)
	GestionErreur(err)
	return id
}


func Rechercher_sauvegarde(v_pseudo string) (int64, bool) {

    // Recherche du pseudo dans la base de données
    recherche := Base_donnees().QueryRow(
        "SELECT id FROM sauvegarde WHERE pseudo = ?",
        v_pseudo,
    )

    var id int64
    err := recherche.Scan(&id)

    if err == sql.ErrNoRows {
        // Aucun compte correspondant
        return 0, false
    }

    if err != nil {
        GestionErreur(err)
        return 0, false
    }

    // Le joueur existe: on renvoie son ID de sauvegarde
    return id, true
}

func Afficher_score_theme(v_id_save int64, v_id_controle_sauvegarde_theme int64) int64{
    var score_total_theme int64

    err := Base_donnees().QueryRow("SELECT COALESCE(SUM(sq.score_question), 0) FROM sauvegardeQuestion sq WHERE sq.sauvegarde_id = ? AND sq.sauvegarde_id_theme = ?", v_id_save, v_id_controle_sauvegarde_theme).Scan(&score_total_theme)
    GestionErreur(err)
	//fmt.Printf("DEBUG: ID Save = %d | ID Theme = %d\n", v_id_save, v_id_controle_sauvegarde_theme)
	//fmt.Println(score_total_theme)
	return score_total_theme
}

func Afficher_score_total(v_id_save int64) int64 {
	var score_total int64

	err := Base_donnees().QueryRow("SELECT COALESCE(SUM(sq.score_question), 0) FROM sauvegardeQuestion sq WHERE sq.sauvegarde_id = ?", v_id_save).Scan(&score_total)
	Base_donnees().QueryRow("UPDATE sauvegarde SET score_total = ? WHERE sauvegarde.id = ?", score_total, v_id_save)
	GestionErreur(err)
	return score_total
}

func Fin_de_partie(v_score_theme int) {
	fmt.Println()
	fmt.Println("Partie terminée !")
	fmt.Println()
}

func Recommencer_theme(v_id_save int64, v_id_controle_sauvegarde_theme int64) {
	Base_donnees().Exec("UPDATE sauvegardeQuestion SET complete = 0, score_question = 0 WHERE sauvegarde_id = ? AND sauvegarde_id_theme = ?", v_id_save, v_id_controle_sauvegarde_theme)
	Base_donnees().Exec("UPDATE sauvegardeTheme SET theme_demarre = 0 WHERE sauvegarde_id = ? AND theme_id = ?", v_id_save, v_id_controle_sauvegarde_theme)
}
