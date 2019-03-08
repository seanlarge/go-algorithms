package main

import (
	"fmt"
	"math"
	"math/rand"
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

func bruteForce(p [] Points) (float64, [] Points){
	 min := math.Inf(1)
	 closestPoints := make([]Points, 2)
	for i := 0; i <= len(p) - 1; i++ {
		for j := i+1; j <= len(p) - 1; j++ {
			if dist(p[i],p[j]) < min{
				min = dist(p[i], p[j])
				closestPoints[0] = p[i]
				closestPoints[1] = p[j]
			}
		}
	}
	return min, closestPoints
}

func closestPair( p [] Points ) (float64, [] Points){

	// brute force
	//if len(xP) <= 10 {
	//	return bruteForce(xP)
	//}

	//	xL ← points of xP from 1 to ⌈N/2⌉
	fmt.Println(len(p))

	xl := sortX(p[ :len(p) / 2  ])
	fmt.Println(xl)
	//	xR ← points of xP from ⌈N/2⌉+1 to N
	xr := sortX(p[len(p) / 2:  ])

	fmt.Println(xr)
	//	xm ← xP(⌈N/2⌉)x
	//	yL ← { p ∈ yP : px ≤ xm }
	//	yR ← { p ∈ yP : px > xm }
	//	(dL, pairL) ← closestPair of (xL, yL)
	//	(dR, pairR) ← closestPair of (xR, yR)
	//	(dmin, pairMin) ← (dR, pairR)
	return 0, nil
}

func main(){
	fmt.Println("D and C alog closest pair of points")

	var points =  [] Points {{12, 30},{2,3}, {1,10},{5,80},{7,100}, {50, 54},{60, 54}, {100, 100}}
	// {80, 4}, {12,18}, {3,4}, {40,50}, {5, 1}, {12, 10}, {3, 4}
	// sort the points by the x coordinates into a list called x


	// divide the points into lists P and Q with all points in P to the left of all points in Q
	// i e divide in half
	//var pair  = closestPair(x,y)
	fmt.Println(closestPair(points))

	// closestPoint( P, N/2), closestPoint(Q, N/2)

}

//closestPair of (xP, yP)
//	where xP is P(1) .. P(N) sorted by x coordinate, and
//		  yP is P(1) .. P(N) sorted by y coordinate (ascending order)
//	if N ≤ 3 then
//	return closest points of xP using brute-force algorithm
//else
//	xL ← points of xP from 1 to ⌈N/2⌉
//	xR ← points of xP from ⌈N/2⌉+1 to N
//	xm ← xP(⌈N/2⌉)x
//	yL ← { p ∈ yP : px ≤ xm }
//	yR ← { p ∈ yP : px > xm }
//	(dL, pairL) ← closestPair of (xL, yL)
//	(dR, pairR) ← closestPair of (xR, yR)
//	(dmin, pairMin) ← (dR, pairR)
//if dL < dR then
//(dmin, pairMin) ← (dL, pairL)
//endif
//yS ← { p ∈ yP : |xm - px| < dmin }
//nS ← number of points in yS
//(closest, closestPair) ← (dmin, pairMin)
//for i from 1 to nS - 1
//k ← i + 1
//while k ≤ nS and yS(k)y - yS(i)y < dmin
//if |yS(k) - yS(i)| < closest then
//(closest, closestPair) ← (|yS(k) - yS(i)|, {yS(k), yS(i)})
//endif
//k ← k + 1
//endwhile
//endfor
//return closest, closestPair
//endif