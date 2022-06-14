package main

import (
	"fmt"
	"sort"
	"sync"
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
	NombreMax := 200
	th := 4
	debut := 0
	interval := Interval(NombreMax, th)
	echantillon := interval
	deb := time.Now()

	var wg sync.WaitGroup
	lockval := sync.Mutex{}
	for i := 0; i < th; i++ {
		if echantillon > NombreMax {
			echantillon = debut + (NombreMax - debut)
		}
		wg.Add(1)
		go func(debut int, echantillon int) {
			val := calcul(tableauNombrePremier, debut, echantillon, NombreMax)
			lockval.Lock()
			tableauNombrePremier = append(tableauNombrePremier, val...)
			lockval.Unlock()
			wg.Done()
		}(debut, echantillon)

		debut += interval
		echantillon += interval
	}
	wg.Wait()
	fin := time.Now()
	sort.Ints(tableauNombrePremier)
	fmt.Println(tableauNombrePremier)
	fmt.Println(fin.Sub(deb))
}

func calcul(tableauNombrePremier []int, debut int, echantillon int, NombreMax int) []int {
	tabProvisoir := []int{}
	for i := debut; i < echantillon; i++ {
		if EstPremier(i, NombreMax) {
			tabProvisoir = append(tabProvisoir, i)
		}
	}
	return tabProvisoir
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
