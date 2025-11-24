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
