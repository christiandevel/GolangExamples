package basico

import "fmt"

func GetMultiplication(number int) {

	fmt.Printf("The multiplication table of %d is -------------\n", number)
	counter := 0
	for counter <= 10 {
		result := number * counter
		fmt.Printf("%d x %d = %d\n", number, counter, result)
		counter++
	}

}
