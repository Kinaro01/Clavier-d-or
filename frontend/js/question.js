// parse = conversion de json a javascript
const theme = JSON.parse(sessionStorage.getItem("theme_selectionne"));

async function afficherScoreTheme() {
    try {
        // récupération id de la save
        const idSauvegarde = sessionStorage.getItem("id_sauvegarde");

        const idTheme = theme.id;

        if (!idSauvegarde || !idTheme) {
            console.error("Aucune sauvegarde trouvée.");
            return;
        }

        const response = await fetch(`/api/scoretheme?id_sauvegarde=${encodeURIComponent(idSauvegarde)}&id_theme=${encodeURIComponent(idTheme)}`);

        if (!response.ok) {
        const erreur = await response.text();

        console.error("Statut HTTP :", response.status);
        console.error("Réponse Go :", erreur);

        return;
    }

    const data = await response.json();

    document.getElementById("ScoreTheme").textContent = data.score_theme;

    } catch (error) {
        console.error(error);
    }
}

afficherScoreTheme();


async function AfficherQuestion() {
    try {
        // récupération id de la save
        const idSauvegarde = sessionStorage.getItem("id_sauvegarde");

        const idTheme = theme.id;

        if (!idSauvegarde || !idTheme) {
            console.error("Aucune sauvegarde trouvée.");
            return;
        }

        const response = await fetch(`/api/afficherquestion?id_sauvegarde=${encodeURIComponent(idSauvegarde)}&id_theme=${encodeURIComponent(idTheme)}`);

        if (!response.ok) {
        const erreur = await response.text();

        console.error("Statut HTTP :", response.status);
        console.error("Réponse Go :", erreur);

        return;
    }

    const data = await response.json();

    document.getElementById("Libelle").textContent = data.nom
    document.getElementById("Indice").textContent = "Indice : "  + data.indice
    document.getElementById("ReponseA").textContent = data.reponse_a
    document.getElementById("ReponseB").textContent = data.reponse_b
    document.getElementById("ReponseC").textContent = data.reponse_c
    document.getElementById("ReponseD").textContent = data.reponse_d

    } catch {
        console.error(error);
    }
}

AfficherQuestion();


// affichage nom theme
if (theme) {
    document.getElementById("TitreTheme").textContent = `Thème : ${theme.nom}`;
}

// bouton retour menu
const boutonMenu = document.getElementById("btn_menu");

boutonMenu.addEventListener("click", () => {
    sessionStorage.removeItem("theme_selectionne")
    window.location.href = "/accueil"
});


