package main

import (
	"fmt"
	"math"
)

func main() {
	var loan float64
	fmt.Printf("Enter amount of loan: ")
	fmt.Scanf("%f", &loan)

	var credit_period float64
	fmt.Printf("Enter period in years: ")
	fmt.Scanf("%f", &credit_period)

	var interest float64
	fmt.Printf("Enter interest rate: ")
	fmt.Scanf("%f", &interest)
	interest = interest / 100

	num_of_month := credit_period * 12

	fmt.Println("|#\t|Payment\t|Principal\t|Interest\tBalance Owed\t|")

	monthly_payment := loan * (interest * math.Pow((1+interest), num_of_month)) / ((math.Pow(1+interest, num_of_month)) - 1)

	for i := 0; i < int(num_of_month); i++ {
		interest_payment := loan * interest
		principal_payment := monthly_payment - interest_payment
		loan = loan - principal_payment
		fmt.Printf("|%d\t|%.2f\t|%.2f\t|%.2f\t|%.2f\t|\n", i+1, monthly_payment, principal_payment, interest_payment, loan)
	}
}
