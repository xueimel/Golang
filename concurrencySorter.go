package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var max = 25000 // make this a ratio of size
//
func main() {
	arrSize := 1000000 //number of objects to sort
	maxValue := 10000  //max number of Values in array
	//make copies of unsorted arrays.
	nums := generateSlice(arrSize, maxValue)
	nums1 := copySlice(nums, arrSize)
	nums2 := copySlice(nums, arrSize)
	nums3 := copySlice(nums, arrSize)
	nums4 := copySlice(nums, arrSize)
	nums5 := copySlice(nums, arrSize)

	func() { //handy copy and paste lambda/anonymous functions
		startTime := time.Now()
		mergeSort(nums2)
		elapsed := time.Since(startTime)
		fmt.Println("Merge Sort: Elapsed: ", elapsed)
	}()
	func() {
		startTime := time.Now()
		mergeSortPar(nums3)
		elapsed := time.Since(startTime)
		fmt.Println("Merge Sort2: Elapsed: ", elapsed)
	}()
	func() {
		startTime := time.Now()
		quicksort(nums4)
		elapsed := time.Since(startTime)
		fmt.Println("Quick Sort: Elapsed: ", elapsed)
	}()
	func() {
		startTime := time.Now()
		quicksortpar(nums5)
		elapsed := time.Since(startTime)
		fmt.Println("Quick Sort2: Elapsed: ", elapsed)
	}()
	func() {
		fmt.Println("Bubble sort, because why not...")
		startTime := time.Now()
		timer()
		bubblesort(nums1)
		elapsed := time.Since(startTime)
		fmt.Println("Bubble Sort: Start time: Elapsed: ", elapsed)
	}()

}
func timer() {
	for i := 0; true; i++ {
		time.Sleep(10 * time.Second)
		if i%2 == 0 {
			fmt.Println("Still trying")
		} else {
			fmt.Println("...I promise, I'm working on it")
		}
	}
}

func quicksortpar(a []int) []int {
	if len(a) < 2 {
		return a
	}
	left, right := 0, len(a)-1
	pivot := rand.Int() % len(a)
	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]

	if len(a) < max { //this check allows for optimization between concurrency and its overhead costs
		quicksort(a[:left])
		quicksort(a[left+1:])

		return a

	} else {
		var waiter sync.WaitGroup //ensures that the concurrent threads aren't forgotten by the main 
		waiter.Add(2)

		go func() {
			quicksortpar(a[:left])
			waiter.Done()
		}()
		go func() {
			quicksortpar(a[left+1:])
			waiter.Done()
		}()
		waiter.Wait()

		return a
	}
}

func quicksort(a []int) []int {

	if len(a) < 2 {
		return a
	}
	left, right := 0, len(a)-1
	pivot := rand.Int() % len(a)
	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a

}

func mergeSortPar(items []int) { //based off of teivah's implementation of mergesort concurrency algo. Can be found at  https://gist.github.com/teivah/392a6b209f5c9c5aec82ef5043edbecc#file-mergesortv2-go
	len := len(items)

	if len > 1 {
		if len <= max { // part that is Sequential
			mergeSort(items)
		} else { // part that runs in parallel
			middle := len / 2

			var waiter sync.WaitGroup
			waiter.Add(2)

			go func() {
				mergeSortPar(items[:middle])
				waiter.Done()
			}()

			go func() {
				mergeSortPar(items[middle:])
				waiter.Done()
			}()
			waiter.Wait()
		}
	}
}

func mergeSort(items []int) []int {
	var num = len(items)

	if num == 1 {
		return items
	}

	middle := int(num / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) (result []int) {

	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}

func bubblesort(items []int) {
	startTime := time.Now()
	var n = len(items)
	var sorted = false
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
	elapsed := time.Since(startTime)
	fmt.Println("Bubble Sort: Start time: ", startTime.Nanosecond(), "(nanoseconds) Elapsed: ", elapsed)
}

func generateSlice(size int, maxValue int) []int {
	lice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		lice[i] = rand.Intn(100)
	}
	return slice
}

func copySlice(items []int, n int) []int {
	slice := make([]int, n)
	copy(slice, items)
	return slice
}
