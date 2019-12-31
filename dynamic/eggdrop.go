package main

import (
	"fmt"
	"time"
)

func main(){
	now := time.Now()
	/// print time after
	defer func() {
		fmt.Println(time.Since(now).Nanoseconds())
	}()
	// run program
	fmt.Printf("We can guess the floor in %v moves \n",eggDrop(10,50000))

}

// keep count in terms of m(number of moves
func eggDrop(K int, N int) int{
	storage :=  make([]int, K+1)
	// number of moves
	var m int
	for m = 0; storage[K] < N; m++ {
		for  k := K; k > 0; k-- {
		storage[k] += storage[k - 1] + 1
	}
		}
	return m
}

//Then, we take 1 move to a floor and the egg will break or not. If the egg breaks, it means we should find the answer under this level,
//and it also means this level can not be higher than dp[m-1][k-1] + 1, otherwise we will unable to get the answer.
//If the egg doesn't break, we will use the k eggs and m-1 moves to go higher. So the highest level we can touch is dp[m-1][k-1] + 1 + dp[m-1][k].
//We can find the answer use k eggs and m moves in any level which not higher than it.