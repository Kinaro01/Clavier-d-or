const idSauvegarde = sessionStorage.getItem("id_sauvegarde");

const theme = JSON.parse(sessionStorage.getItem("theme_selectionne"));
const idTheme = theme.id;

const boutonRecomencer = document.getElementById("btn_recommencer")

boutonRecomencer.addEventListener("click", async () => {
    try {
        const response = await fetch("/api/recommencertheme", {
            method: "PUT",
            headers: {
                "Content-Type": "application/json"
            },
            // stringify = conversion de javascript a json
            body: JSON.stringify({
                id_sauvegarde: Number(idSauvegarde),
                id_theme: Number(idTheme)
            })
        });
        if (!response.ok) {
            throw new Error("Erreur lors de la réinitialisation du thème");
        }

        window.location.href = "/question"
    } catch (error) {
        console.error(error)
    }
});


// retour au menu
const boutonMenu = document.getElementById("btn_menu")

boutonMenu.addEventListener("click", () => {
    window.location.href = "/accueil"
})