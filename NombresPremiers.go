package main

//Fonction moins rapide
func EstPremier1(nb int) bool {
	for i := 1; i < nb; i++ {
		if nb%i == 0 && i != 1 && i != nb {
			return false
		}
	}
	return true
}

//Fonction opti
func EstPremier(nb int) bool {
	if nb%2 == 0 {
		return false
	}
	for i := 1; i < nb; i += 2 {
		if nb%i == 0 && i != 1 && i != nb {
			return false
		}
	}
	return true
}
