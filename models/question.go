package models


type Question struct {
	Id          int `json:"id"`
	Theme_id    int `json:"theme_id"`
	Type        string `json:"type"`
	Nom         string `json:"nom"`
	Reponse_a   string `json:"reponse_a"`
	Reponse_b   string `json:"reponse_b"`
	Reponse_c   string `json:"reponse_c"`
	Reponse_d   string `json:"reponse_d"`
	Est_correct string `json:"est_correct"`
	Indice 		string `json:"indice"`
}


