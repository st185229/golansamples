// A data race occurs when two goroutines access the same variable concurrently
// and at least one of the accesses is a write.

// Races occur due to communication
package main

import (
	"fmt"
	"sort"
	"sync"
)

var parts = 4

func RoutineSort(numbers []int, id int, wg *sync.WaitGroup) []int {
	sort.Ints(numbers)
	fmt.Println("Gourotine ", id+1, ":", numbers)
	defer wg.Done()
	return numbers
}

func read(n int) ([]int, error) {
	in := make([]int, n)
	for i := range in {
		_, err := fmt.Scan(&in[i])
		if err != nil {
			return in[:i], err
		}
	}
	return in, nil
}

func chunkSlice(slice []int, chunkSize int) [][]int {
	var chunks [][]int
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func main() {

	fmt.Println("Insert 12 integers: ")
	var wg sync.WaitGroup
	ints, err := read(12)

	if err != nil {
		fmt.Println(err)
		return
	}

	chunkSize := (len(ints) + parts - 1) / parts

	partitions := chunkSlice(ints, chunkSize)

	wg.Add(parts)
	for i := 0; i < parts; i++ {

		go RoutineSort(partitions[i], i, &wg)

	}
	wg.Wait()

	s0 := partitions[0]
	s1 := append(s0, partitions[1]...)
	s2 := append(s1, partitions[2]...)
	s3 := append(s2, partitions[3]...)

	sort.Ints(s3)

	fmt.Println("Sorted Output", s3)

}
