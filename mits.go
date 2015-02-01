/*
Package mits implements simple text segmentation by calculate words' mutual information value

*/
package mits

import (
//	"math"
)

func CountSingleWord(sentence string) (words map[string]int) {
	words = make(map[string]int)
	for _, word := range sentence {
		words[string(word)] += 1
	}
	return
}

/*
func FilterStopWords(stopword, sentence string) string {
	return ""
}
*/

func CalcMI(single, twin map[string]int) (mi float64) {
	return
}

/*
func MutualInformation(word1, word2 rune) float64 {
	fmt.Println("hello")
	return 0.01
}
*/
