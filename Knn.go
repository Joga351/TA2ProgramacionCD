package main

import (
	"math"
	"sort"
)

type Data struct {
	//numero del dato
	Index string
	//features de la data
	CAI              string
	Edad             int
	Trabajo          bool
	Vinculo          bool
	TipoVio          int
	ConsumeAlcohol   bool
	Fuma             bool
	ConsumeDroga     bool
	Adiccion         int
	RiesgoPresuntivo int
	Mes              int
}

type inferir struct {
	Index     string
	Distancia int
}

type VecinoMasCercano interface {
	Exec(obj Data, train []Data) string
	CalcularDistancia(a, b Data) int
}

type KNNAlgoritmo struct {
	KModificaciones int
}

func NuevoVecinoMasCercano(KModificaciones int) VecinoMasCercano {
	return &KNNAlgoritmo{
		KModificaciones: KModificaciones,
	}
}

func (k *KNNAlgoritmo) CalcularDistancia(a, b Data) float64 {
	return math.Sqrt(math.Pow(a.X1-b.X1, 2) + math.Pow(a.X2-b.X2, 2) + math.Pow(a.X3-b.X3, 2) + math.Pow(a.X4-b.X4, 2) + math.Pow(a.X5-b.X5, 2))
}

func (k *KNNAlgoritmo) Exec(a Data, train []Data) string {
	var dists []inferir
	inf := map[string]int{
		"0": 0,
		"1": 0,
		"2": 0,
		"3": 0,
	}

	// Calculate all distance between object to train data.
	for _, b := range train {
		var dist inferir
		dist.Index = b.CAI
		dist.Distancia = int(k.CalcularDistancia())(a, b)
		dists = append(dists, dist)
	}

	sort.Slice(dists, func(x, y int) bool {
		if dists[x].Distance == dists[y].Distance {
			return dists[x].Name < dists[y].Name
		}
		return dists[x].Distance < dists[y].Distance
	})

	// Get K nearest data.
	for j := 0; j < k.K; j++ {
		if dists[j].Name == "0" {
			inf["0"]++
		} else if dists[j].Name == "1" {
			inf["1"]++
		} else if dists[j].Name == "2" {
			inf["2"]++
		} else if dists[j].Name == "3" {
			inf["3"]++
		}
	}

	max := inf["0"]
	key := "0"
	for index, e := range inf {
		if max < e {
			max = e
			key = index
		}
	}

	// Return the most inference showed up in the first K data.
	return key
}
