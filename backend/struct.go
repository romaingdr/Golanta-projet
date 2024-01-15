package backend

type Aventurier struct {
	ID           int    `json:"id"`
	Nom          string `json:"nom"`
	Prenom       string `json:"prenom"`
	Surnom       string `json:"surnom"`
	Age          int    `json:"age"`
	Tribu        string `json:"tribu"`
	Sexe         string `json:"sexe"`
	Force        int    `json:"force"`
	Intelligence int    `json:"intelligence"`
	Strategie    int    `json:"strategie"`
	Description  string `json:"description"`
}

type AventuriersData struct {
	Aventuriers []Aventurier `json:"aventuriers"`
}
