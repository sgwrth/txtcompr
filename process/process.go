package process

import (
	"bufio"
	"bytes"
	"os"
	"sort"
)

func CompressFile(file *os.File) {
	// defer file.Close()
	// tokens := GetTokens(file)
}

func GetTokens(file *os.File) []string {
	var tokens []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		tokens = append(tokens, word)
	}
	return tokens
}

func FreqMap(tokens []string) map[string]int {
	freqMap := make(map[string]int)
	for _, token := range tokens {
		if len(token) > 1 {
			freqMap[token]++
		}
	}
	return freqMap
}

func KeysOfMap(stringIntMap map[string]int) []string {
	keys := make([]string, 0, len(stringIntMap))
	for key := range stringIntMap {
		keys = append(keys, key)
	}
	return keys
}

func KeysSortedByVal(stringIntMap map[string]int) []string {
	keys := KeysOfMap(stringIntMap)
	sort.Slice(keys, func(i, j int) bool {
		return stringIntMap[keys[i]] > stringIntMap[keys[j]]
	})
	return keys
}

func DuplicateEntries(freqMap map[string]int) []Entry {
	sortedKeys := KeysSortedByVal(freqMap)
	var entries []Entry
	for _, key := range sortedKeys {
		if freqMap[key] > 1 {
			entries = append(entries, Entry{key, freqMap[key]})
		}
	}
	return entries
}

func BuildDict(entries []Entry) map[string]byte {
	dict := make(map[string]byte)
	code := byte(1)
	for _, e := range entries {
		dict[e.word] = code
		code++
	}
	return dict
}

func CompressedTextBytes(tokens []string, dict map[string]byte) []byte {
	var compressedText bytes.Buffer
	ESCAPE := byte(0)
	for _, word := range tokens {
		if code, ok := dict[word]; ok {
			compressedText.WriteByte(code)
		} else {
			compressedText.WriteByte(ESCAPE)
			compressedText.WriteString(word + " ")
		}
	}
	return compressedText.Bytes()
}
