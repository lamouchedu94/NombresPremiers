package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"sync"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")

func main() {
	var err error
	nb := 1.0       // Valeur de départ
	valMax := 10000 //Valeur du nombre maximum calculé
	th := 16

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

	var tabNb []int
	var wg sync.WaitGroup
	deb := time.Now()

	inter := interval(valMax, th)
	for i := 0; i < th; i++ {

		wg.Add(1)
		go func(i int, nb float64, valMax int, inter int, tabNb []int) {
			AffichageAppend(i, nb, valMax, inter, tabNb)
			wg.Done()
		}(i, nb, valMax, inter, tabNb)

	}
	wg.Wait()
	fin := time.Now()
	_ = tabNb
	//fmt.Println(tabNb)
	fmt.Println(fin.Sub(deb))
}

func AffichageAppend(j int, nb float64, valMax int, interval int, tabNb []int) {
	for i := interval * j; i < interval*(j+1); i++ {
		res := testNb(nb, valMax)
		if res == 1 {
			//fmt.Println(nb) //, "Est un nombre premier")
			tabNb = append(tabNb, int(nb))
		} /*else {
			fmt.Println(nb, "N'est pas un nombre premier")
		}*/

		nb += 1.0
	}
	//fmt.Println(tabNb)
}

func testNb(nb float64, max int) int {
	for i := 0; i < max; i++ {
		res := nb / float64(i)
		if res == float64(int(res)) && res != nb && res != 1 {
			//fmt.Println(nb, "n'est pas un nombre premier")
			return 0
		}
	}
	return 1
}

func interval(nb int, th int) int {
	return nb / th
}
