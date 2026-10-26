async function afficherScoreTotal() {
    try {
        // récupération id de la save
        const idSauvegarde = sessionStorage.getItem("id_sauvegarde");

        if (!idSauvegarde) {
            console.error("Aucune sauvegarde trouvée.");
            return;
        }

        const response = await fetch(`/api/scoretotal?id_sauvegarde=${encodeURIComponent(idSauvegarde)}`);

        if (!response.ok) {
        const erreur = await response.text();

        console.error("Statut HTTP :", response.status);
        console.error("Réponse Go :", erreur);

        return;
    }

    const data = await response.json();

    document.getElementById("ScoreTotal").textContent = data.score_total;

    } catch (error) {
        console.error(error);
    }
}

afficherScoreTotal();