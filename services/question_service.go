package services

import (
	"database/sql"
	"projet/models"


)


func Rechercher_question(v_id_save int64, v_id_controle_sauvegarde_theme int64) *models.Question {

	// récupération des éléments de la table questions
	// + right join vers la table sauvegardeQuestion pour trier par save et id du thème
	var q models.Question
	resultat_question := Base_donnees().QueryRow("SELECT q.id, q.theme_id, q.type_question, q.libelle, q.option_a, q.option_b, q.option_c, q.option_d, q.est_correcte, q.indice FROM questions q RIGHT JOIN sauvegardeQuestion sq ON q.id = sq.question_id WHERE sq.sauvegarde_id = ? AND sq.sauvegarde_id_theme = ? AND sq.complete = false ORDER BY sq.question_id ASC", v_id_save, v_id_controle_sauvegarde_theme)
	err := resultat_question.Scan(&q.Id, &q.Theme_id, &q.Type, &q.Nom, &q.Reponse_a, &q.Reponse_b, &q.Reponse_c, &q.Reponse_d, &q.Est_correct, &q.Indice)
	if err == sql.ErrNoRows {
		return nil
	}
	return &q
}



func Resultat_question(v_choix_reponse_utilisateur string, q models.Question, v_id_save int64) bool{
	var resultat bool
	// test si résultat = bon ou pas
	if v_choix_reponse_utilisateur == q.Est_correct {
		resultat = true
		Base_donnees().QueryRow("UPDATE sauvegardeQuestion SET complete = 1 WHERE sauvegardeQuestion.question_id = ?", q.Id)
		Base_donnees().QueryRow("UPDATE sauvegardeTheme SET theme_demarre = 1 WHERE sauvegardeTheme.theme_id = ? AND sauvegardeTheme.sauvegarde_id = ?", q.Theme_id, v_id_save)
		Base_donnees().QueryRow("UPDATE sauvegardeQuestion SET score_question = 5 WHERE sauvegardeQuestion.question_id = ?", q.Id)
	} else {
		resultat = false
		Base_donnees().QueryRow("UPDATE sauvegardeQuestion SET complete = 1 WHERE sauvegardeQuestion.question_id = ?", q.Id)
		Base_donnees().QueryRow("UPDATE sauvegardeTheme SET theme_demarre = 1 WHERE sauvegardeTheme.theme_id = ? AND sauvegardeTheme.sauvegarde_id = ?", q.Theme_id, v_id_save)
	}
	return resultat
}

//func Passer_question(q models.Question, v_id_save int64) bool{
	//Base_donnees().QueryRow("UPDATE sauvegardeQuestion SET complete = 1 WHERE sauvegardeQuestion.question_id = ?", q.Id)
	//Base_donnees().QueryRow("UPDATE sauvegardeTheme SET theme_demarre = 1 WHERE sauvegardeTheme.theme_id = ? AND sauvegardeTheme.sauvegarde_id = ?", q.Theme_id, v_id_save)
	//return false
//}
