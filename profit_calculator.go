package main

import (
	"errors"
	"fmt"
	"os"
)

//Goals
// 1) Validate user input
//	=> Show error message and exit if invalid input is provided
//	- no negative numbers
//	- Not 0
// 2) Store calculated results into file

func main() {
	revenue, err := askValues("Revenue: ")

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		return
	}

	expenses, err := askValues("Expenses: ")

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		return
	}

	taxRate, err := askValues("Tax Rate: ")

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		return
	}

	earningsBeforeTax, earningsAfterTax, ratio := calculate(revenue, expenses, taxRate)

	writeCalculationsToFile(earningsBeforeTax, earningsAfterTax, ratio)

	fmt.Printf("The EBT is %.2f\n", earningsBeforeTax)
	fmt.Printf("The Profit is %.2f\n", earningsAfterTax)
	fmt.Printf("The Ratio is %.2f\n", ratio)
}

func askValues(text string) (returnValue float64, err error) {
	fmt.Print(text)
	fmt.Scan(&returnValue)
	if returnValue == 0 {
		return 0, errors.New("Value can't be 0.")
	} else if returnValue < 0 {
		return 0, errors.New("Value must be greater then 0.")
	}
	return returnValue, nil
}

func calculate(revenue float64, expenses float64, taxRate float64) (earningsBeforeTax float64, earningsAfterTax float64, ratio float64) {
	earningsBeforeTax = revenue - expenses

	earningsAfterTax = earningsBeforeTax * (1 - taxRate/100)

	ratio = earningsBeforeTax / earningsAfterTax

	return earningsBeforeTax, earningsAfterTax, ratio
}

func writeCalculationsToFile(earningsBeforeTax float64, earningsAfterTax float64, ratio float64) {
	data := fmt.Sprintf("Earning Before Tax: %.2f\nEarnings After tax: %.2f\nRatio: %.2f", earningsBeforeTax, earningsAfterTax, ratio)
	os.WriteFile("CalculatedResults.txt", []byte(data), 0644)
}
