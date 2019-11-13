package main

import "math/rand"

var names = [...]string{
	"Alice",
	"Benjamin",
	"Claudia",
	"Devon",
	"Ella",
	"Frankie",
	"Gray",
	"Helen",
	"Iggy",
	"Juliana",
	"Kelly",
	"Lex",
	"Manny",
	"Oswald",
	"Peter",
	"Quentin",
	"Rachel",
	"Sera",
	"Theo",
	"Ursula",
	"Valerie",
	"Xavier",
	"Yingtian",
	"Zoe",
}

//these will never be the same name
func randomNamePair() (string, string) {
	var i1, i2 int
	for i1 == i2 {
		i1 = rand.Intn(len(names))
		i2 = rand.Intn(len(names))
	}
	return names[i1], names[i2]
}
