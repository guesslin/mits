/*
Package mits implements simple text segmentation by calculate words' mutual information value

*/
package mits

import (
//	"math"
)

// CountSingleWord count frequency of each single rune word appear
// in single file
func CountSingleWord(sentence []rune) (single map[string]int) {
	single = make(map[string]int)
	for _, word := range sentence {
		single[string(word)] += 1
	}
	return
}

// CountTwinWord
func CountTwinWord(sentence []rune) (twin map[string]int) {
	twin = make(map[string]int)
	for i, _ := range sentence {
		if i == len(sentence)-1 {
			break // last word of sentence
		}
		two := sentence[i : i+2]
		twin[string(two)] += 1
	}
	return
}

// CountTermFreq count frequency of single word and twin-word term
func CountTermFreq(sentence []rune) (single, twin map[string]int) {
	single = make(map[string]int)
	twin = make(map[string]int)
	for i, word := range sentence {
		single[string(word)] += 1
		if i < len(sentence)-1 {
			two := sentence[i : i+2]
			twin[string(two)] += 1
		}
	}
	return
}

// CalcMI use single word frequency and twin word term frequency to
// calculate mutual information value of continue term
// mutual information value formula:
//                    p(x[i], x[i+1]) * log(p(x[i], x[i+1])
// MI(x[i], x[i+1]) = -------------------------------------
//                           (p(x[i]) * p(x[i+1])))
func CalcMI(single, twin map[string]int) (mi float64) {
	return
}
