package process

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"
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

func DictAsBinary(dict map[string]byte) []byte {
	var dictAsBytes bytes.Buffer
	dictAsBytes.WriteByte(byte(len(dict)))

	for word, code := range dict {
		dictAsBytes.WriteByte(code)
		dictAsBytes.WriteByte(byte(len(word)))
		dictAsBytes.WriteString(word)
	}

	return dictAsBytes.Bytes()
}

func WriteDictAndTextToFile(filePath string, dict []byte, compressedText []byte) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	_, err = file.Write(dict)

	if err != nil {
		return err
	}

	_, err = file.Write(compressedText)
	return err
}

func RecreateDict(file []byte, startingPos int) (map[byte]string, int) {
	dictLen := int(file[0])
	dict := make(map[byte]string)
	bytePos := startingPos
	for range dictLen {
		code := file[bytePos]
		bytePos++
		wordLen := file[bytePos]
		var wordAsByte bytes.Buffer
		bytePos++
		for j := 0; j < int(wordLen); j++ {
			wordAsByte.WriteByte(file[bytePos])
			bytePos++
		}
		dict[code] = wordAsByte.String()
	}
	return dict, bytePos
}

func DecodeText(file []byte, startingPos int, dict map[byte]string) []string {
	var tokens []string
	bytePos := startingPos
	for bytePos < len(file) {
		if file[bytePos] != byte(0) {
			tokens = append(tokens, dict[file[bytePos]])
			bytePos++
		} else {
			bytePos++
			delimiter := byte(' ')
			var word bytes.Buffer
			for file[bytePos] != delimiter && bytePos < len(file) {
				word.WriteByte(file[bytePos])
				bytePos++
			}
			tokens = append(tokens, word.String())
			bytePos++
		}
	}
	return tokens
}

func DecompressFile(filepath string) {
	file, err := os.ReadFile(filepath)

	if err != nil {
		panic(err)
	}

	dict, startingPosCompressedText := RecreateDict(file, 1)
	tokens := DecodeText(file, startingPosCompressedText, dict)
	fmt.Println(strings.Join(tokens, " "))
}
