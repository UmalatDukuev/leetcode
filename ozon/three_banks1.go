package main

import (
	"bufio"
	"fmt"
	"os"
)

type Transfer struct {
	From    float64
	To      float64
	Visited bool
}

type Bank struct {
	RToD    Transfer
	RToE    Transfer
	DToR    Transfer
	DToE    Transfer
	EToR    Transfer
	EToD    Transfer
	Visited bool
}

func findMaxDollar(banks []Bank) float32 {
	return 0.015
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var startRubles float64
	startRubles = 1.0
	bankA := Bank{
		RToD: Transfer{From: 100, To: 1},
		RToE: Transfer{From: 100, To: 1},
		DToR: Transfer{From: 1, To: 100},
		DToE: Transfer{From: 3, To: 2},
		EToR: Transfer{From: 1, To: 100},
		EToD: Transfer{From: 2, To: 3},
	}

	bankB := Bank{
		RToD: Transfer{From: 100, To: 1},
		RToE: Transfer{From: 100, To: 1},
		DToR: Transfer{From: 1, To: 100},
		DToE: Transfer{From: 3, To: 2},
		EToR: Transfer{From: 1, To: 100},
		EToD: Transfer{From: 2, To: 3},
	}

	bankC := Bank{
		RToD: Transfer{From: 100, To: 1},
		RToE: Transfer{From: 100, To: 1},
		DToR: Transfer{From: 1, To: 100},
		DToE: Transfer{From: 3, To: 2},
		EToR: Transfer{From: 1, To: 100},
		EToD: Transfer{From: 2, To: 3},
	}
	banks := []Bank{bankA, bankB, bankC}
	result := findMaxDollar(banks)
	for _, bank := range banks {
		firstOper := startRubles * (bank.RToD.To / bank.RToD.From)
	}
	// d := make([]float64, 0)
	// e := make([]float64, 0)

	// d = append(d, startRubles*(float64(bankA.RublesToDollars.To)/float64(bankA.RublesToDollars.From)))
	// d = append(d, startRubles*(float64(bankB.RublesToDollars.To)/float64(bankB.RublesToDollars.From)))
	// d = append(d, startRubles*(float64(bankC.RublesToDollars.To)/float64(bankC.RublesToDollars.From)))
	// e = append(e, startRubles*(float64(bankA.RublesToEuros.To)/float64(bankA.RublesToEuros.From)))
	// e = append(e, startRubles*(float64(bankB.RublesToEuros.To)/float64(bankB.RublesToEuros.From)))
	// e = append(e, startRubles*(float64(bankC.RublesToEuros.To)/float64(bankC.RublesToEuros.From)))
	// // firstOperation = append(firstOperation, )
	// firstMaxD := max(d)
	// firstMaxE := max(e)

	// fmt.Println(firstMaxD)
	// fmt.Println(firstMaxE)
	fmt.Println(result)
}

func max(slice []float64) float64 {
	if len(slice) == 0 {
		return 0
	}
	maxVal := slice[0]
	for _, v := range slice[1:] {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}
