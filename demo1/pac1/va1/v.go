package va1

import (
	"_/demo1/pac1/va2"
	"fmt"
)

func SI()  {
	fmt.Print("Enter the Principal or Total Amount = ")
	fmt.Scanln(&va2.Amount)

	fmt.Print("Enter the rate of Interest = ")
	fmt.Scanln(&va2.InterestRate)

	fmt.Print("Enter the Total number of Years = ")
	fmt.Scanln(&va2.Time)

	va2.SimpleI = (va2.Amount * va2.InterestRate * va2.Time) / 100

	fmt.Println("\nThe Simple Interest  = ", va2.SimpleI)
	
}