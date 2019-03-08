package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main(){
 // random generator here for different N to help analyze
 // Numbers for N 20k, 50k, 200k, 300k, 500k
 P := make([]Points, 100000000)
 for i := 0; i < 100000000; i++{
 	P[i].X = rand.Intn(101)
 	P[i].Y = rand.Intn(101)
 }
    // start time
	now := time.Now()
	/// print time after
	defer func() {
		fmt.Println(time.Since(now).Nanoseconds())
	}()
	// run program
   fmt.Println(stairCase(P))
}

type Points struct{
	X int
	Y int
}


func stairCase(originalPoints [] Points)(ParetoPoints [] Points){
	sortedPoints := mergeSort(originalPoints)
	// large negative number that outputs to: -9223372036854775808
	currentY :=  -math.MaxInt64
	for i :=0; i < len(sortedPoints) ; i++  {
		if sortedPoints[i].Y > currentY {
			ParetoPoints = append(ParetoPoints, sortedPoints[i])
			currentY = sortedPoints[i].Y
		}
	}
	return
}


func mergeSort(items []Points) []Points {
	var num = len(items)

	if num == 1 {
		return items
	}

	middle := int(num / 2)
	var (
		left = make([]Points, middle)
		right = make([]Points, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return mergeDescending(mergeSort(left), mergeSort(right))
}

// descending order for this use case
func mergeDescending(left, right []Points) (result []Points) {
	result = make([]Points, len(left) + len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0].X > right[0].X {
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