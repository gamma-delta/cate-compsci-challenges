package main

//A Transaction is one exchange of money from person to person
type Transaction struct {
	From   string `json:"name"`
	Amount int    `json:"amount"`
}
