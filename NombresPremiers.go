package main

func EstPremier(nb int) bool {
	for i := 1; i < nb; i++ {
		if nb%i == 0 && i != 1 && i != nb {
			return false
		}
	}
	return true
}
