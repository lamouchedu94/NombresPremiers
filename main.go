package main

import (
	"fmt"
	"time"
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
	tableauNombrePremier := []int{}
	NombreMax := 100
	th := 1
	debut := 0
	interval := Interval(NombreMax, th)
	echantillon := interval
	deb := time.Now()

	//var wg sync.WaitGroup

	for i := 0; i < th; i++ {
		if echantillon > NombreMax {
			echantillon = debut + (NombreMax - debut)
		}
		tableauNombrePremier = calcul(tableauNombrePremier, debut, echantillon, NombreMax)
		debut += interval
		echantillon += interval
	}

	fin := time.Now()
	fmt.Println(tableauNombrePremier)
	fmt.Println(fin.Sub(deb))
}

func calcul(tableauNombrePremier []int, debut int, interval int, NombreMax int) []int {
	for i := debut; i < interval; i++ {
		if EstPremier(i, NombreMax) {
			tableauNombrePremier = append(tableauNombrePremier, i)
		}
	}
	return tableauNombrePremier
}

func EstPremier(nb int, valMax int) bool {
	for i := 1; i < valMax; i++ {
		if nb%i == 0 && i != 1 && i != nb {
			return false
		}
	}
	return true
}

func Interval(valMax int, th int) int {
	return (valMax / th) + 1
}
