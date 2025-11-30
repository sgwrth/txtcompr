package process

import (
	"testing"

	"dev.sgwrth/txtcompr/files"
)

func TestTokenize(t *testing.T) {
	file := files.OpenFile("../data/allwork.txt")
	defer file.Close()
	tokens := GetTokens(file)
	tokensLen := len(tokens)
	if tokensLen != 9 { // Not counting one-letter word 'a'.
		t.Errorf("Wrong length of tokens slice: %v", tokensLen)
	}
}

func TestCountOccurrences(t *testing.T) {
	file := files.OpenFile("../data/arose.txt")
	defer file.Close()
	tokens := GetTokens(file)
	roseOccurrences := CountOccurrences("rose", tokens)
	if roseOccurrences != 3 {
		t.Errorf("Wrong number of occurrences: %v", roseOccurrences)
	}
}

func TestFreqMap(t *testing.T) {
	tokens := GetTokens(files.OpenFile("../data/arose.txt"))
	freqMap := FreqMap(tokens)
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

func TestSortedKeysOfMap(t *testing.T) {
	stringIntMap := map[string]int{
		"one":   42,
		"two":   69,
		"three": 404,
	}
	keys := KeysSortedByVal(stringIntMap)
	if len(keys) != 3 {
		t.Errorf("Wrong length of 'keys' slice: %v", len(keys))
	}
	if keys[0] != "three" {
		t.Errorf("Wrong value (%v) at index 0", keys[0])
	}
	if keys[1] != "two" {
		t.Errorf("Wrong value (%v) at index 1", keys[1])
	}
	if keys[2] != "one" {
		t.Errorf("Wrong value (%v) at index 2", keys[2])
	}
}

func TestDuplicateEntries(t *testing.T) {
	file := files.OpenFile("../data/arose.txt")
	defer file.Close()
	tokens := GetTokens(file)
	freqMap := FreqMap(tokens)
	entries := DuplicateEntries(freqMap)
	if len(entries) != 2 { // 'rose', 'is'
		t.Errorf("Wrong length of entries: %v", len(entries))
	}
	if entries[0].word != "rose" {
		t.Errorf("Wrong top entry: %v", entries[0].word)
	}
	if entries[1].word != "is" {
		t.Errorf("Wrong runner-up entry: %v", entries[1].word)
	}
}

func TestBuildDict(t *testing.T) {
	file := files.OpenFile("../data/arose.txt")
	defer file.Close()
	tokens := GetTokens(file)
	freqMap := FreqMap(tokens)
	duplicateEntries := DuplicateEntries(freqMap)
	dict := BuildDict(duplicateEntries)
	_, ok := dict["rose"]
	if !ok {
		t.Errorf("Dict entry 'rose' missing")
	}
	if dict["rose"] != byte(1) {
		t.Errorf("Wrong value for 'rose': %v", dict["rose"])
	}
	_, ok = dict["is"]
	if !ok {
		t.Errorf("Dict entry 'is' missing")
	}
	if dict["is"] != byte(2) {
		t.Errorf("Wrong value for 'is': %v", dict["is"])
	}
}
