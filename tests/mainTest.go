package main

var noMutante = [6]string{
	"ATGCGA",
	"CAGTGC",
	"TTATTT",
	"AGACGG",
	"GCGTCA",
	"TCACTG",
}

var mutante = [6]string{
	"ATGCGA",
	"CAGTGC",
	"TTATGT",
	"AGAAGG",
	"CCCCTA",
	"TCACTG",
}

func main() {
	if isMutant(mutante) {
		println("Mutante encontrado!")
	} else {
		println("No se encontró ningún mutante.")
	}
}
