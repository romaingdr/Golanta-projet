package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"
)

func GetAventurierIDs() ([]int, error) {
	filename := "aventuriers.json"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var aventuriersData AventuriersData
	err = json.Unmarshal(data, &aventuriersData)
	if err != nil {
		return nil, err
	}

	var ids []int
	for _, aventurier := range aventuriersData.Aventuriers {
		ids = append(ids, aventurier.ID)
	}

	return ids, nil
}

func ParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Erreur lors de la conversion de la chaîne en entier:", err)
	}
	return i
}

func GenerateRandomID() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(9000) + 1000
}

func Contains(list []int, value int) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

func SupprimerAventurierParID(id int, aventuriers *[]Aventurier) bool {
	index := -1
	for i, a := range *aventuriers {
		if a.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		return false
	}

	*aventuriers = append((*aventuriers)[:index], (*aventuriers)[index+1:]...)

	return true
}

func ChargerAventuriersParEquipe(cheminFichier string, equipe string) ([]Aventurier, error) {
	data, err := ioutil.ReadFile(cheminFichier)
	if err != nil {
		return nil, err
	}

	var aventuriersData AventuriersData

	err = json.Unmarshal(data, &aventuriersData)
	if err != nil {
		return nil, err
	}

	var aventuriersEquipe []Aventurier
	for _, aventurier := range aventuriersData.Aventuriers {
		if aventurier.Tribu == equipe {
			aventuriersEquipe = append(aventuriersEquipe, aventurier)
		}
	}

	return aventuriersEquipe, nil
}
