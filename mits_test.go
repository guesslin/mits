package mits

import (
	"reflect"
	"testing"
)

func TestCountSingleWord(t *testing.T) {
	cases := []struct {
		sentence []rune
		result   map[string]int
	}{
		{sentence: []rune("今天的天氣真的很好的說"), result: map[string]int{"今": 1, "天": 2, "的": 3, "氣": 1, "真": 1, "很": 1, "好": 1, "說": 1}},
		{sentence: []rune("是不是這樣的日子裡，你才會這樣的想起我"), result: map[string]int{"會": 1, "的": 2, "日": 1, "子": 1, "才": 1, "想": 1, "我": 1, "這": 2, "不": 1, "樣": 2, "裡": 1, "，": 1, "是": 2, "起": 1, "你": 1}},
	}
	for _, c := range cases {
		result := CountSingleWord(c.sentence)
		eq := reflect.DeepEqual(result, c.result)
		if !eq {
			t.Errorf("\nExpect\t%+v\nGet\t%+v\n", c.result, result)
		}
	}
}

func TestCountTwinWord(t *testing.T) {
	cases := []struct {
		sentence []rune
		result   map[string]int
	}{
		{sentence: []rune("今天的天氣真的很好的說"), result: map[string]int{"今天": 1, "天的": 1, "的天": 1, "天氣": 1, "氣真": 1, "真的": 1, "的很": 1, "很好": 1, "好的": 1, "的說": 1}},
	}
	for _, c := range cases {
		result := CountTwinWord(c.sentence)
		eq := reflect.DeepEqual(result, c.result)
		if !eq {
			t.Errorf("\nExcept\t%+v\nGet\t%+v\n", c.result, result)
		}
	}
}

func TestCountTermFreq(t *testing.T) {
	cases := []struct {
		sentence                 []rune
		singleResult, twinResult map[string]int
	}{
		{sentence: []rune("今天的天氣真的很好的說"), singleResult: map[string]int{"今": 1, "天": 2, "的": 3, "氣": 1, "真": 1, "很": 1, "好": 1, "說": 1}, twinResult: map[string]int{"今天": 1, "天的": 1, "的天": 1, "天氣": 1, "氣真": 1, "真的": 1, "的很": 1, "很好": 1, "好的": 1, "的說": 1}},
		//		{sentence: []rune("是不是這樣的日子裡，你才會這樣的想起我"), singleResult: map[string]int{"會": 1, "的": 2, "日": 1, "子": 1, "才": 1, "想": 1, "我": 1, "這": 2, "不": 1, "樣": 2, "裡": 1, "，": 1, "是": 2, "起": 1, "你": 1}},
	}
	for _, c := range cases {
		single, twin := CountTermFreq(c.sentence)
		eq1 := reflect.DeepEqual(single, c.singleResult)
		eq2 := reflect.DeepEqual(twin, c.twinResult)
		if !eq1 {
			t.Errorf("except\t%+v\nGet\t%+v\n", c.singleResult, single)
		}
		if !eq2 {
			t.Errorf("except\t%+v\nGet\t%+v\n", c.twinResult, twin)
		}
	}

}
