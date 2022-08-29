package api

import (
	"fmt"
	"testing"
)

func TestGetSentence(t *testing.T) {
	sentenceEnglish, sentenceChinese, err := GetSentence()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println(sentenceChinese)
	fmt.Println(sentenceEnglish)
}
