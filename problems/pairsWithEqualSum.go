package problems

import (
	model "probsolbygolang/model"
)

func PairsWithEqualSum(arr []int) map[int][]model.IntergerPair {
	pairsWithEqualSum := make(map[int][]model.IntergerPair)
	var pairs []model.IntergerPair
	for i := range arr {
		for j := range arr {
			if i != j {
				pair := model.IntergerPair{Number1: arr[i], Number2: arr[j]}
				pairs = append(pairs, pair)
			}
		}
	}
	for i := range pairs {
		_, present := pairsWithEqualSum[pairs[i].Number1+pairs[i].Number2]
		if present {
			pairsWithEqualSum[pairs[i].Number1+pairs[i].Number2] = append(pairsWithEqualSum[pairs[i].Number1+pairs[i].Number2], pairs[i])
		} else {
			pairsWithEqualSum[pairs[i].Number1+pairs[i].Number2] = []model.IntergerPair{pairs[i]}
		}
	}
	return pairsWithEqualSum
}
