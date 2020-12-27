package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter integers: ")
	input, _ := reader.ReadString('\n')
	stringSlice := strings.Fields(input)
	var numbers []int
	for _, elem := range stringSlice {
		num, _ := strconv.Atoi(elem)
		numbers = append(numbers, num)
	}
	size := len(numbers) / 4
	for i := 0; i < 4; i++ {
		start := i * size
		wg.Add(1)
		if i == 3 {
			go mySort(numbers[start :], i, &wg)
		} else {
			end := (i + 1) * size
			go mySort(numbers[start : end], i, &wg)
		}
	}
	wg.Wait()
	fmt.Println(numbers)
	var sortedNumbers []int
	for i := 0; i < len(numbers); i++ {
		if i < 1 * size {
			sortedNumbers = append(sortedNumbers, numbers[i])
		} else {
			if numbers[i] <= sortedNumbers[0] {
				sortedNumbers = append(sortedNumbers, 0)
				copy(sortedNumbers[1:], sortedNumbers[0:])
				sortedNumbers[0] = numbers[i]
			} else if numbers[i] >= sortedNumbers[len(sortedNumbers) - 1] {
				sortedNumbers = append(sortedNumbers, numbers[i])
			} else {
				l := len(sortedNumbers)
				for j := 0; j < l - 1; j++ {
					if sortedNumbers[j] <= numbers[i] && numbers[i] <= sortedNumbers[j + 1] {
						sortedNumbers = append(sortedNumbers, 0)
						copy(sortedNumbers[j + 1 :], sortedNumbers[j :])
						sortedNumbers[j + 1] = numbers[i]
						break
					}
				}
			}
		}
	}
	fmt.Printf("The final sorted array: %v\n", sortedNumbers)
}

func mySort(a []int, i int, wg *sync.WaitGroup) {
	fmt.Printf("The %d routine will sort %v\n", i, a)
	sort.Ints(a)
	wg.Done()
}