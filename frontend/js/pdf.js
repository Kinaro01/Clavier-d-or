const boutonPDF = document.getElementById("export_pdf");

boutonPDF.addEventListener("click", async () => {
    const idSauvegarde = sessionStorage.getItem("id_sauvegarde");
    const theme = JSON.parse(sessionStorage.getItem("theme_selectionne"));

    if (!idSauvegarde || !theme || !theme.id) {
        console.error("Identifiants de sauvegarde ou de thème manquants.");
        return;
    }

    const url = `/api/exportpdf?id_sauvegarde=${encodeURIComponent(idSauvegarde)}&id_theme=${encodeURIComponent(theme.id)}`;

    try {
        const reponse = await fetch(url);

        if (!reponse.ok) {
            throw new Error("Erreur lors de la génération du PDF.");
        }

        // PDF sous forme de blob (données binaires)
        const pdf = await reponse.blob();

        // Création d'un lien temporaire
        const lien = document.createElement("a");
        const urlPDF = URL.createObjectURL(pdf);

        lien.href = urlPDF;
        lien.download = "resultats_quizz.pdf";

        document.body.appendChild(lien);
        lien.click();
        lien.remove();

        // Libération de l'url temporaire
        URL.revokeObjectURL(urlPDF);

    } catch (erreur) {
        console.error("Impossible de télécharger le PDF :", erreur);
    }
});