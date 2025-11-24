package process

import (
	"testing"

	"dev.sgwrth/txtcompr/files"
)

func TestTokenize(t *testing.T) {
	file := files.OpenFile("../data/test.txt")
	tokens := GetTokens(file)
	tokensLen := len(tokens)
	if tokensLen != 10 {
		t.Errorf("Wrong length of tokens slice: %v", tokensLen)
	}
}
