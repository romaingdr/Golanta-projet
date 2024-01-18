package controller

import (
	"Golanta/backend"
	"Golanta/templates"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

func DeletePage(w http.ResponseWriter, r *http.Request) {
	jsonData, _ := ioutil.ReadFile("aventuriers.json")
	queryParams := r.URL.Query()

	idParam := queryParams.Get("id")

	idToDelete, _ := strconv.Atoi(idParam)

	type AventuriersData struct {
		Aventuriers []backend.Aventurier `json:"aventuriers"`
	}

	var aventurierData AventuriersData
	json.Unmarshal(jsonData, &aventurierData)

	if backend.SupprimerAventurierParID(idToDelete, &aventurierData.Aventuriers) {

		jsonUpdated, _ := json.MarshalIndent(aventurierData, "", "  ")

		ioutil.WriteFile("aventuriers.json", jsonUpdated, os.ModePerm)

		http.Redirect(w, r, "/aventuriers", http.StatusSeeOther)
	} else {
		fmt.Printf("Aventurier avec ID %d non trouvé.\n", idToDelete)
	}

}

func EquipesPage(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "equipes", nil)
}

func EquipePage(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	idParam := queryParams.Get("id")

	var aventuriers []backend.Aventurier
	aventuriers, _ = backend.ChargerAventuriersParEquipe("aventuriers.json", idParam)
	templates.Temp.ExecuteTemplate(w, "equipe", aventuriers)
}

func EditPage(w http.ResponseWriter, r *http.Request) {
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

	templates.Temp.ExecuteTemplate(w, "edit", aventurierRecherche)
}

func SubmitEdit(w http.ResponseWriter, r *http.Request) {
	idStr := r.PostFormValue("id")
	id, _ := strconv.Atoi(idStr)

	data, _ := ioutil.ReadFile("aventuriers.json")

	var aventuriersData backend.AventuriersData
	json.Unmarshal(data, &aventuriersData)

	index := -1
	for i, aventurier := range aventuriersData.Aventuriers {
		if aventurier.ID == id {
			index = i
			break
		}
	}

	aventuriersData.Aventuriers[index].Nom = r.PostFormValue("nom")
	aventuriersData.Aventuriers[index].Prenom = r.PostFormValue("prenom")
	aventuriersData.Aventuriers[index].Surnom = r.PostFormValue("surnom")
	aventuriersData.Aventuriers[index].Age, _ = strconv.Atoi(r.PostFormValue("age"))
	aventuriersData.Aventuriers[index].Tribu = r.PostFormValue("tribu")
	aventuriersData.Aventuriers[index].Sexe = r.PostFormValue("sexe")
	aventuriersData.Aventuriers[index].Force, _ = strconv.Atoi(r.PostFormValue("force"))
	aventuriersData.Aventuriers[index].Intelligence, _ = strconv.Atoi(r.PostFormValue("intelligence"))
	aventuriersData.Aventuriers[index].Strategie, _ = strconv.Atoi(r.PostFormValue("strategie"))
	aventuriersData.Aventuriers[index].Description = r.PostFormValue("description")

	nouvellesDonneesJSON, _ := json.MarshalIndent(aventuriersData, "", "  ")

	ioutil.WriteFile("aventuriers.json", nouvellesDonneesJSON, 0644)

	http.Redirect(w, r, "/aventuriers", http.StatusSeeOther)
}
