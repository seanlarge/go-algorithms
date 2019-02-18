package main

import (
	"fmt"
	"math/rand"
	"math"
)

// points data structure
type Points struct {
	X int
	Y int
}
// util to sort x points
func sortX(a[] Points) []Points {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i].X < a[right].X {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	sortX(a[:left])
	sortX(a[left+1:])

	return a
}

// util to sort y
// todo refactor to one function
func sortY(a[] Points) []Points {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i].Y < a[right].Y {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	sortY(a[:left])
	sortY(a[left+1:])

	return a
}

// A utility function to find the distance between two points
func dist(p1 Points, p2 Points) float64{
	num := float64((p1.X - p2.X)*(p1.X - p2.X) +
		(p1.Y - p2.Y)*(p1.Y - p2.Y))
	return math.Sqrt(num)
}

func closestPair(xP[] Points, yP[] Points ){
	// brute force
	if len(xP) <= 3 {
		bruteForce(xP, len(xP))
	}
}

func bruteForce(p [] Points, n int) float64{
	var min float64 = math.Inf(3)
	for i := 0; i < n; i++ {
		for j := i+1; i < n; j++ {
			if dist(p[i],p[j]) < min{
				min = dist(p[i], p[j])
			}
		}
	}
	return min
}

func main(){
	fmt.Println("D and C alog closest pair of points")

	var points=  [] Points {{2,3}, {12, 30}, {50, 54},}
	// {80, 4}, {12,18}, {3,4}, {40,50}, {5, 1}, {12, 10}, {3, 4}
	// sort the points by the x coordinates into a list called x
	var x = sortX(points)


	// same for y
	 var y = sortY(points)


	// divide the points into lists P and Q with all points in P to the left of all points in Q
	// i e divide in half
	closestPair(x,y)

	// closestPoint( P, N/2), closestPoint(Q, N/2)

}

//closestPair of (xP, yP)
//	where xP is P(1) .. P(N) sorted by x coordinate, and
//		  yP is P(1) .. P(N) sorted by y coordinate (ascending order)
//	if N ≤ 3 then
//	return closest points of xP using brute-force algorithm
