const boutonConnexion = document.getElementById("btn_connexion");

boutonConnexion.addEventListener("click", async () => {
    const pseudo = document.getElementById("pseudo").value.trim();
    const message = document.getElementById("message");

    if (pseudo === "") {
        message.textContent = "Veuillez entrer un pseudo.";
        return;
    }

    try {
        const response = await fetch("/api/connexion", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            // stringify = conversion de javascript a json
            body: JSON.stringify({ pseudo: pseudo })
        });

        const data = await response.json();

        // si ok alors stockage de l'id dans une variable + redirection vers l'accueil
        // sinon redirection vers l'inscription
        if (response.ok) {
            sessionStorage.setItem(
                "id_sauvegarde",
                data.id_sauvegarde
            );

            window.location.href = "/accueil";
        } else if (response.status === 404) {
            window.location.href = "/inscription";
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
