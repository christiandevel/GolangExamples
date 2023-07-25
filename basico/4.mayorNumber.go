package basico

import "fmt"

func FindMayorNumer(numbers []int) {
	number := 0
	for i := range numbers {
		if numbers[i] > number {
			number = numbers[i]
		}
	}
	fmt.Printf("El numero mayor es %d \n", number)
}
