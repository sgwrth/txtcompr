package process

import (
	"bufio"
	"os"
)

func CompressFile(file *os.File) {
	defer file.Close()
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
