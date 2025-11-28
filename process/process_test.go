package process

import (
	"testing"

	"dev.sgwrth/txtcompr/files"
)

func TestTokenize(t *testing.T) {
	file := files.OpenFile("../data/allwork.txt")
	tokens := GetTokens(file)
	tokensLen := len(tokens)
	if tokensLen != 9 { // Not counting one-letter word 'a'.
		t.Errorf("Wrong length of tokens slice: %v", tokensLen)
	}
}

func TestCountOccurrences(t *testing.T) {
	tokens := GetTokens(files.OpenFile("../data/arose.txt"))
	roseOccurrences := CountOccurrences("rose", tokens)
	if roseOccurrences != 3 {
		t.Errorf("Wrong number of occurrences: %v", roseOccurrences)
	}
}

func TestBuildFreqMap(t *testing.T) {
	tokens := GetTokens(files.OpenFile("../data/arose.txt"))
	freqMap := BuildFreqMap(tokens)
	if freqMap["rose"] != 3 {
		t.Errorf("Wrong frequency count: %v", freqMap["rose"])
	}
	if freqMap["is"] != 2 {
		t.Errorf("Wrong frequency count: %v", freqMap["is"])
	}
}

func TestKeysOfMap(t *testing.T) {
	stringIntMap := map[string]int{
		"one":   404,
		"two":   69,
		"three": 42,
	}
	keys := KeysOfMap(stringIntMap)
	if len(keys) != 3 {
		t.Errorf("Wrong length of 'keys' slice: %v", len(keys))
	}
	if keys[0] != "one" {
		t.Errorf("Wrong value (%v) at index 0", keys[0])
	}
	if keys[1] != "two" {
		t.Errorf("Wrong value (%v) at index 1", keys[1])
	}
	if keys[2] != "three" {
		t.Errorf("Wrong value (%v) at index 2", keys[2])
	}
}
