/*
Package mits implements simple text segmentation by calculate words' mutual information value

*/
package mits

import (
	"math"
)

// Terms contains result of MITS which means "term in string => probability of each term"
type Terms map[string]float64

// Mits read sentences and translate them into pieces of single char and double chars
// use
func Mits() {
}

// prob returns probability of single word in all documents/sentence
func prob() {
}

// CountSingleWord count frequency of each single rune word appear
// in single file
func CountSingleWord(sentence []rune) (single map[string]int, count int) {
	single = make(map[string]int)
	count = len(sentence)
	for _, word := range sentence {
		single[string(word)] += 1
	}
	return
}

// CountTwinWord
func CountTwinWord(sentence []rune) (twin map[string]int, count int) {
	twin = make(map[string]int)
	count = len(sentence) - 1
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
func CountTermFreq(sentence []rune) (single, twin map[string]int, s_len, t_len int) {
	single = make(map[string]int)
	twin = make(map[string]int)
	s_len = len(sentence)
	for i, word := range sentence {
		single[string(word)] += 1
		if i < len(sentence)-1 {
			two := sentence[i : i+2]
			twin[string(two)] += 1
			t_len += 1
		}
	}
	return
}

// CalcMI use single word frequency and twin word term frequency to
// calculate mutual information value of continue term
// mutual information value formula:
// 	MI(x[i], x[i+1]) = p(x[i], x[i+1]) * log(p(x[i], x[i+1]) / (p(x[i]) * p(x[i+1])))
func CalcMI(single, twin map[string]int, word string) (mi float64) {
	return
}

// Segment sentence and return segmented terms by mutual information
// threshold
func Segment(sentence []rune, theta float64) (terms []string) {
	terms = make([]string, 0)
	single, twin, s_len, t_len := CountTermFreq(sentence)
	var term = string(sentence[0])
	for i := 1; i <= len(sentence)-1; i++ {
		x2 := string(sentence[i-1 : i+1])
		posofx2 := float64(twin[x2]) / float64(t_len)
		posofxi := float64(single[string(sentence[i-1])]) / float64(s_len)
		posofxj := float64(single[string(sentence[i])]) / float64(s_len)
		mi := posofx2 * math.Log(posofx2/(posofxi*posofxj))
		if mi < theta {
			terms = append(terms, term)
			term = string(sentence[i])
		} else {
			term += string(sentence[i])
		}
	}
	return
}
