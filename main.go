package main

type Data struct {
	//numero del dato
	Index string `json:"Ind"`
	//features de la data
	CAI              string `json:"Cai"`
	Edad             int    `json:"Ed"`
	Trabajo          bool   `json:"Tra"`
	Vinculo          bool   `json:"Vin"`
	TipoVio          int    `json:"Tip"`
	ConsumeAlcohol   bool   `json:"ConA"`
	Fuma             bool   `json:"Fum"`
	ConsumeDroga     bool   `json:"ConD"`
	Adiccion         int    `json:"Adi"`
	RiesgoPresuntivo int    `json:"Rie"`
	Mes              int    `json:"M"`
}

func main() {
	router := mux.NewRouter()
}
