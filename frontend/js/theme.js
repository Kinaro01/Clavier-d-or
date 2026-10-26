// retour a la page de connexion
const BoutonDeco = document.getElementById("btn_deco")

BoutonDeco.addEventListener("click", () => {
    sessionStorage.removeItem("id_sauvegarde");
    window.location.href = "/";
})

async function afficherTheme() {
    try {
        const idSauvegarde = sessionStorage.getItem("id_sauvegarde");

        if (!idSauvegarde) {
            console.error("Aucune sauvegarde trouvée.");
            return;
        }

        const response = await fetch(
            `/api/affichertheme?id_sauvegarde=${encodeURIComponent(idSauvegarde)}`
        );

        if (!response.ok) {
            console.error("Statut HTTP :", response.status);
            console.error("Réponse Go :", await response.text());
            return;
        }

        const themes = await response.json();

        // création d'un bouton pour chaque thème avec un nom de classe commun
        themes.forEach(theme => {
        const carte = document.createElement("button");

        carte.textContent = theme.nom;
        carte.classList.add("ThemeName");

        carte.addEventListener("click", () => {
            sessionStorage.setItem("theme_selectionne", JSON.stringify(theme));
            window.location.href = "/question";
        });

       document.getElementById("ThemeCard").appendChild(carte);
    });

    } catch (error) {
        console.error(error);
    }
}

afficherTheme();