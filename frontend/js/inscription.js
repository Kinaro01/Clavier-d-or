
const boutonInscription = document.getElementById("btn_inscription");

boutonInscription.addEventListener("click", async () => {
    const pseudo = document.getElementById("pseudo").value.trim();
    const message = document.getElementById("message");

    if (pseudo === "") {
        message.textContent = "Veuillez entrer un pseudo.";
        return;
    }

    try {
        const response = await fetch("/api/inscription", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ pseudo: pseudo })
        });

        const data = await response.json();

        // si ok alors stockage de l'id dans une variable + redirection vers l'accueil
        // sinon erreur
        if (response.ok) {
            sessionStorage.setItem(
                "id_sauvegarde",
                data.id_sauvegarde
            );

            window.location.href = "/accueil";
        } else {
            message.textContent =
                data.message || "Une erreur est survenue.";
        }
    } catch (error) {
        console.error("Erreur de connexion :", error);
        message.textContent =
            "Impossible de contacter le serveur.";
    }
});
