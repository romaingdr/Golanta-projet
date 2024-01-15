package controller

import (
	"Golanta/backend"
	"Golanta/templates"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func AccueilPage(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "accueil", nil)
}

func CreatePage(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "create", nil)
}

func SuccessCreate(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "success", nil)
}

func SubmitCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()

		existingIDs, _ := backend.GetAventurierIDs()

		var nouveauID int
		for {
			nouveauID = backend.GenerateRandomID()
			if !backend.Contains(existingIDs, nouveauID) {
				break
			}
		}

		nouveauAventurier := backend.Aventurier{
			ID:           nouveauID,
			Nom:          r.FormValue("nom"),
			Prenom:       r.FormValue("prenom"),
			Surnom:       r.FormValue("surnom"),
			Age:          backend.ParseInt(r.FormValue("age")),
			Sexe:         r.FormValue("sexe"),
			Tribu:        r.FormValue("tribu"),
			Force:        backend.ParseInt(r.FormValue("force")),
			Intelligence: backend.ParseInt(r.FormValue("intelligence")),
			Strategie:    backend.ParseInt(r.FormValue("strategie")),
			Description:  r.FormValue("description"),
		}

		var aventuriersData backend.AventuriersData

		file, _ := ioutil.ReadFile("aventuriers.json")

		json.Unmarshal(file, &aventuriersData)

		aventuriersData.Aventuriers = append(aventuriersData.Aventuriers, nouveauAventurier)

		data, _ := json.MarshalIndent(aventuriersData, "", "  ")

		ioutil.WriteFile("aventuriers.json", data, 0644)

		fmt.Println("Aventurier ajouté avec succès")
		http.Redirect(w, r, "/success_create", http.StatusSeeOther)
		return
	}

	templates.Temp.ExecuteTemplate(w, "create", nil)
}

func AventuriersPage(w http.ResponseWriter, r *http.Request) {
	var aventuriersData backend.AventuriersData

	jsonData, _ := ioutil.ReadFile("aventuriers.json")

	err := json.Unmarshal(jsonData, &aventuriersData)
	if err != nil {
		fmt.Println("Erreur lors du décodage JSON:", err)
		return
	}
	templates.Temp.ExecuteTemplate(w, "aventuriers", aventuriersData)
}

func AventurierPage(w http.ResponseWriter, r *http.Request) {
	var aventurierData backend.AventuriersData
	queryParams := r.URL.Query()

	idParam := queryParams.Get("id")

	id, _ := strconv.Atoi(idParam)

	jsonData, _ := ioutil.ReadFile("aventuriers.json")

	json.Unmarshal(jsonData, &aventurierData)

	var aventurierRecherche backend.Aventurier
	for _, aventurier := range aventurierData.Aventuriers {
		if aventurier.ID == id {
			aventurierRecherche = aventurier
			break
		}
	}

	if aventurierRecherche.ID == 0 {
		fmt.Println("Aventurier non trouvé avec l'ID:", id)
		http.Error(w, "Aventurier non trouvé", http.StatusNotFound)
		return
	}

	templates.Temp.ExecuteTemplate(w, "aventurier", aventurierRecherche)
}
