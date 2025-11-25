package process

import (
	"bufio"
	// "encoding/base64"
	"os"
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
		token := scanner.Text()
		if len(token) >= 2 {
			tokens = append(tokens, token)
		}
	}
	return tokens
}

func CountOccurrences(token string, tokens []string) int {
	counter := 0
	for _, element := range tokens {
		if element == token {
			counter++
		}
	}
	return counter
}

func BuildFreqMap(tokens []string) map[string]int {
	freqMap := make(map[string]int)
	for _, token := range tokens {
		freqMap[token]++
	}
	return freqMap
}
