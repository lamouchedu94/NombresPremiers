package main

import (
	"flag"
	"fmt"
)

func arguments() (bool, int, bool, int) {
	//Affichage := false
	arret := false
	//th := 0
	//var err error
	ArgAffichage := flag.Bool("v", false, "Argument affichage")
	ArgThread := flag.Int("th", 0, "Nb Thread")
	ArgEstPremier := flag.Int("n", 0, "Valeur précise à vérifier")
	flag.Parse()

	if len(flag.Args()) == 0 && *ArgEstPremier == 0 {
		fmt.Println("./main {option} nombre")
		fmt.Println("Option : \n-v pour afficher tableau de nombre\n-th {nombre thread} pour choisir nb thread (max dispo par defaut)")
		arret = true
	}

	return *ArgAffichage, *ArgThread, arret, *ArgEstPremier
}
