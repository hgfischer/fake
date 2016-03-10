package fake

import "math/rand"

// Brand generates brand name
func Brand() string {
	return Company()
}

// ProductName generates product name
func ProductName() string {
	productName := lookup(lang, "adjectives", true) + " " + lookup(lang, "nouns", true)
	if rand.Intn(2) == 1 {
		productName = lookup(lang, "adjectives", true) + " " + productName
	}
	return productName
}

// Product generates product title as brand + product name
func Product() string {
	return Brand() + " " + ProductName()
}

// Model generates model name that consists of letters and digits, optionally with a hyphen between them
func Model() string {
	seps := []string{"", " ", "-"}
	return CharactersN(rand.Intn(3)+1) + seps[rand.Intn(len(seps))] + Digits()
}
