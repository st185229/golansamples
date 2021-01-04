package main

import(
	"fmt"
	"sync"
)

func Sort(arr []int, wg *sync.WaitGroup, seg int){
	defer wg.Done()
	defer fmt.Println("Ending Go Routine:", seg)
	fmt.Println("Starting Go Routine:", seg)
	ln := len(arr)
	for i:=ln-1 ; i>0 ; i--{
		isSorted := true
		for j:=0 ; j<i ; j++{
			if arr[j] > arr[j+1]{
				arr[j], arr[j+1] = arr[j+1], arr[j]
				isSorted = false
			}
		}
		if isSorted{
			break
		}
	}
}

func main(){
	var size int
	fmt.Println("Enter the size of the array:")
	fmt.Scan(&size)
	fmt.Println("Enter the array:")
	arr := make([]int, size)

	for i:=0 ; i<size ; i++{
		fmt.Scan(&arr[i])
	}

	var wg sync.WaitGroup
	wg.Add(4)
	go Sort(arr[0:size/4], &wg, 0)
	go Sort(arr[size/4:size/2], &wg, 1)
	go Sort(arr[size/2:(3*size)/4], &wg, 2)
	go Sort(arr[(3*size)/4:size], &wg, 3)
	wg.Wait()

	var ptr [4]int = [4]int{0, size/4, size/2, (3*size)/4}
	srtArr := make([]int, size)
	cur := 0
	for;cur<size;{
		smallestSeg := -1
		smallestEle := 10000000
		for j:=1 ; j<=4 ; j++{
			if ptr[j-1] >= (j*size)/4{continue}

			if arr[ptr[j-1]] < smallestEle{
				smallestEle = arr[ptr[j-1]]
				smallestSeg = j
			} 
		}
		if smallestSeg == -1{break}
		srtArr[cur] = smallestEle
		cur++
		ptr[smallestSeg-1]++
	}

	fmt.Println("Sorted Array:")
	for i:=0 ; i<size ; i++{
		fmt.Printf("%d ", srtArr[i])
	}
	fmt.Println()

}