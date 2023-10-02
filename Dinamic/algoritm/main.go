package main

import (
	"fmt"
	"strconv"
)

func main(){
	dis, final := loadWays()
	start := 1 

	calculateDistance(start, final, dis)

}

func calculateDistance(start, final int, waysDistance [][]int) {
	totalDistance := make(map[string]int)

	for i := final; i >= 1; i-- {
		lastDistance := 0
		for j := final; j >= 1; j-- {
			way := waysDistance[i][j]
			if way == 0 {
				continue
			}

			if lastDistance == 0 {
				lastDistance = way
			}

			if lastDistance > way {
				lastDistance = way

				totalDistance = map[string]int{
					strconv.FormatInt(int64(i), 10) + "-" + strconv.FormatInt(int64(j), 10): lastDistance,
				}

				
				fmt.Println(strconv.FormatInt(int64(i), 10) + "-" + strconv.FormatInt(int64(j), 10))
			}
		}
	}

	fmt.Println(totalDistance)
}

func loadWays() ([][]int, int) {
	return [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 50, 4, 25, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 5, 36, 0, 0},
		{0, 0, 0, 0, 0, 0, 11, 0, 0},
		{0, 0, 0, 0, 0, 0, 9, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 11},
		{0, 0, 0, 0, 0, 0, 0, 0, 78},
		{0, 0, 0, 0, 0, 0, 0, 0, 42},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}, 8
}
