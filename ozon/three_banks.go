package main

import (
	"bufio"
	"os"
)

type Transfer struct {
	From int
	To   int
}

type Bank struct {
	RublesToDollars Transfer
	RublesToEuros   Transfer
	DollarsToRubles Transfer
	DollarsToEuros  Transfer
	EurosToRubles   Transfer
	EurosToDollars  Transfer
	Visited         bool
}

func findMaxDollar(banks []Bank) float32 {
	return 0.015
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	bankA := Bank{
		RublesToDollars: Transfer{From: 100, To: 1},
		RublesToEuros:   Transfer{From: 100, To: 1},
		DollarsToRubles: Transfer{From: 1, To: 100},
		DollarsToEuros:  Transfer{From: 3, To: 2},
		EurosToRubles:   Transfer{From: 1, To: 100},
		EurosToDollars:  Transfer{From: 2, To: 3},
	}

	bankB := Bank{
		RublesToDollars: Transfer{From: 100, To: 1},
		RublesToEuros:   Transfer{From: 100, To: 1},
		DollarsToRubles: Transfer{From: 1, To: 100},
		DollarsToEuros:  Transfer{From: 3, To: 2},
		EurosToRubles:   Transfer{From: 1, To: 100},
		EurosToDollars:  Transfer{From: 2, To: 3},
	}

	bankC := Bank{
		RublesToDollars: Transfer{From: 100, To: 1},
		RublesToEuros:   Transfer{From: 100, To: 1},
		DollarsToRubles: Transfer{From: 1, To: 100},
		DollarsToEuros:  Transfer{From: 3, To: 2},
		EurosToRubles:   Transfer{From: 1, To: 100},
		EurosToDollars:  Transfer{From: 2, To: 3},
	}
	banks := []Bank{bankA, bankB, bankC}
	result := findMaxDollar(banks)
	println(bankA)
	println(result)
}
