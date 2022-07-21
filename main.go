package main

import (
	"fmt"
	"sync"
)

// Конкурентно порахувати суму кожного слайсу int, та роздрукувати результат.
// Потрібно використовувати WaitGroup.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// Порядок друку не важливий.
// “slice 1: 10”
// “slice 2: 16”
func main() {
	// Розкоментуй мене)
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	// Ваша реалізація
	var wg sync.WaitGroup
	var outputMutex sync.Mutex

	for i, subSlice := range n {

		wg.Add(1)
		sSlCopy, iCopy := subSlice, i

		go func() {
			outputMutex.Lock()
			defer wg.Done()
			fmt.Printf("slice %d: ", iCopy)
			sum(sSlCopy)
			fmt.Println()
			outputMutex.Unlock()
		}()
	}
	wg.Wait()
}

func sum(arr []int) {

	var sumInt int

	// Ваша реалізація
	for _, element := range arr {
		sumInt += element

	}

	fmt.Printf("%v", sumInt)
}
