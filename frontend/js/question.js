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

    if (data === null || data === undefined) {
        window.location.href = "/resultat";
        return;
    }

    document.getElementById("Libelle").textContent = data.nom ?? "";
    document.getElementById("Indice").textContent = "Indice : "  + data.indice ?? "";
    document.getElementById("ReponseA").textContent = data.reponse_a ?? "";
    document.getElementById("ReponseB").textContent = data.reponse_b ?? "";
    document.getElementById("ReponseC").textContent = data.reponse_c ?? "";
    document.getElementById("ReponseD").textContent = data.reponse_d ?? "";

    } catch(error) {
        console.error(error);
    }
}

AfficherQuestion();


async function VerifierQuestion(choixReponse) {
    try {
         // récupération id de la save
        const idSauvegarde = sessionStorage.getItem("id_sauvegarde");

        const idTheme = theme.id;

        if (!idSauvegarde || !idTheme) {
            console.error("Aucune sauvegarde trouvée.");
            return;
        }

        const response = await fetch(`/api/verifierquestion?id_sauvegarde=${idSauvegarde}&id_theme=${theme.id}`,
            {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({
                    choix_reponse: choixReponse
                })
            }
        );

        if (!response.ok) {
            const erreur = await response.text();

            console.error("Statut HTTP :", response.status);
            console.error("Réponse Go :", erreur);

        return;
    }

    const resultat = await response.json();

    if (resultat.correct) {
        document.getElementById("Reponse" + choixReponse)
            .classList.add("reponse-correcte");

    } else {
        document.getElementById("Reponse" + choixReponse)
            .classList.add("reponse-incorrecte");
    }

    } catch {
        console.error(error);
    }
}

function DesactiverReponses() {
    document.getElementById("ReponseA").disabled = true;
    document.getElementById("ReponseB").disabled = true;
    document.getElementById("ReponseC").disabled = true;
    document.getElementById("ReponseD").disabled = true;
}


function ActiverReponses() {
    document.getElementById("ReponseA").disabled = false;
    document.getElementById("ReponseB").disabled = false;
    document.getElementById("ReponseC").disabled = false;
    document.getElementById("ReponseD").disabled = false;
}

function ReinitialiserCouleurs() {
    const reponses = ["A", "B", "C", "D"];

    reponses.forEach(function(lettre) {
        const bouton = document.getElementById("Reponse" + lettre);

        bouton.classList.remove("reponse-correcte");
        bouton.classList.remove("reponse-incorrecte");
    });
}

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

// bouton indice
const boutonIndice = document.getElementById("btn_indice");
const texteIndice = document.getElementById("Indice");

boutonIndice.addEventListener("click", () => {
    texteIndice.style.visibility = "visible";
})

// verification quand on choisit une reponse
const boutonSuivant = document.getElementById("QuestionSuivante");

document.getElementById("ReponseA").addEventListener("click", async function() {
    DesactiverReponses();

    await VerifierQuestion("A");

    boutonSuivant.style.display = "block";
});

document.getElementById("ReponseB").addEventListener("click", async function() {
    DesactiverReponses();

    await VerifierQuestion("B");

    boutonSuivant.style.display = "block";
});

document.getElementById("ReponseC").addEventListener("click", async function() {
    DesactiverReponses();

    await VerifierQuestion("C");

    boutonSuivant.style.display = "block";
});

document.getElementById("ReponseD").addEventListener("click", async function() {
    DesactiverReponses();

    await VerifierQuestion("D");

    boutonSuivant.style.display = "block";
});

// Passer a la question suivante
boutonSuivant.addEventListener("click", () => {
    boutonSuivant.style.display = "none";
    texteIndice.style.visibility = "hidden"

    ReinitialiserCouleurs();
    AfficherQuestion();
    ActiverReponses();
    afficherScoreTheme();
});