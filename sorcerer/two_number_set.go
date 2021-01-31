/*
Weights for scales are given. And the final weight that you need to gain with two weights.
It is necessary to find out which two weights can make up the desired weight.
If no such combination exists, return an empty array.
*/
package main

func GetWeight(weights []int, finalWeight int) []int {
	for i := 0; i < len(weights)-1; i++ {
		for j := i + 1; j < len(weights); j++ {
			if weights[i]+weights[j] == finalWeight {
				return []int{weights[i], weights[j]}
			}
		}
	}

	return []int{}
}
