package main

import (
	"fmt"
)

//var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")

func main() {
	/*
		var err error
		nb := 1          // Valeur de départ
		valMax := 100000 //Valeur du nombre maximum calculé
		th := 16         //runtime.NumCPU() //Utilise le nombre max de coeur dispo ou a remplacer par le nb voulu.
		flag.Parse()
		if flag.NArg() > 0 {
			nb, err = strconv.ParseFloat(flag.Arg(0), 32)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		if *cpuprofile != "" {
			f, err := os.Create(*cpuprofile)
			if err != nil {
				log.Fatal("could not create CPU profile: ", err)
			}
			defer f.Close() // error handling omitted for example
			if err := pprof.StartCPUProfile(f); err != nil {
				log.Fatal("could not start CPU profile: ", err)
			}
			defer pprof.StopCPUProfile()
		}
	*/
	var tabNbPremier []int
	tabNbPremier = calculInterval(100, tabNbPremier)
	fmt.Println(tabNbPremier)

}

func calculInterval(valMax int, tabNbPremier []int) []int {
	for i := 0; i < valMax; i++ {
		if EstPremier(i, valMax) {
			tabNbPremier = append(tabNbPremier, i)
		}
	}
	return tabNbPremier
}

func EstPremier(nb int, valMax int) bool {
	for i := 1; i < valMax; i++ {
		if nb%i == 0 && i != 1 && i != nb {
			return false
		}
	}
	return true
}
