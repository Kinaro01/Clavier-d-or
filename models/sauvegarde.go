package models


type Sauvegarde struct {
	Id          int64 `json:"id"`
	Pseudo      string `json:"pseudo"`
	Date        string `json:"date"`
	Score_total int `json:"score_total"`
}

type SauvegardeTheme struct {
	Id            int64 `json:"id"`
	Sauvegarde_id int `json:"sauvegarde_id"`
	Theme_id      int `json:"theme_id"`
	Theme_demarre bool `json:"theme_demarre"`
}

type SauvegardeQuestion struct {
	Id                  int64 `json:"id"`
	Sauvegarde_id_theme int `json:"sauvegarde_id_theme"`
	Sauvegarde_id       int `json:"sauvegarde_id"`
	Question_id         int `json:"question_id"`
	Complete            bool `json:"complete"`
	Score_question      int `json:"score_question"`
}

