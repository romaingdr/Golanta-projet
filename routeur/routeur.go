package routeur

import (
	"Golanta/controller"
	"fmt"
	"log"
	"net/http"
)

func Initserv() {

	css := http.FileServer(http.Dir("./assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", css))

	http.HandleFunc("/accueil", controller.AccueilPage)
	http.HandleFunc("/create", controller.CreatePage)
	http.HandleFunc("/submit_create", controller.SubmitCreate)
	http.HandleFunc("/success_create", controller.SuccessCreate)
	http.HandleFunc("/aventuriers", controller.AventuriersPage)
	http.HandleFunc("/aventurier", controller.AventurierPage)
	http.HandleFunc("/delete", controller.DeletePage)
	http.HandleFunc("/equipes", controller.EquipesPage)
	http.HandleFunc("/equipe", controller.EquipePage)
	http.HandleFunc("/edit", controller.EditPage)
	http.HandleFunc("/submit_edit", controller.SubmitEdit)

	// Démarrage du serveur
	log.Println("[✅] Serveur lancé !")
	fmt.Println("[🌐] http://localhost:8080/accueil")
	http.ListenAndServe(":8080", nil)
	log.Fatal()
}
