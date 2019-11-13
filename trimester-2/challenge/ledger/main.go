package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gonum.org/v1/gonum/stat/distuv"
)

const (
	problemSize   = 100
	moneyVariance = 50
)

func main() {
	out := make(map[string][]Transaction, problemSize)

	dist := distuv.Normal{
		Mu:    0,
		Sigma: moneyVariance,
	}

	//The part that actually makes the data.
	for c := 0; c < problemSize; c++ {
		name1, name2 := randomNamePair()
		money := 0       //init
		for money == 0 { //I don't want anyone giving anyone $0...
			money = int(dist.Rand()) //...so I keep generating numbers until I get one which isn't 0
		}
		moneyComplement := -money //each send has a recieve
		out[name1] = append(out[name1], Transaction{From: name2, Amount: money})
		out[name2] = append(out[name2], Transaction{From: name1, Amount: moneyComplement})
	}

	//Marshall to json so the kids can actually read it
	jsonOut, _ := json.MarshalIndent(out, "", "    ") //I live dangerously and discard errors

	//Create a file with an incremented name
	//For example, if I have `ledger0.dat`, `ledger1.dat` and `ledger2.dat`,
	//I'll make `ledger3.dat`
	var fileName string
	for c := 0; ; c++ { //just count up c forever
		fileName = fmt.Sprintf("ledger%d.dat", c) //create the filename
		_, err := os.Stat(fileName)               //gen info on this file
		if os.IsNotExist(err) {                   //this file doesn't exist!
			//so we're free to create it
			break
		}
	}

	ioutil.WriteFile(fileName, jsonOut, os.ModePerm) //i eat errors for breakfast
	//this is bad practice you should not discard errors like this
	fmt.Println("Successfully created file", fileName)
}
