// Copyright (c) 2020 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package main

import (
	swearfilter "github.com/JoshuaDoes/gofuckyourself"
	"regexp"
	"strings"
)

// Filter variable that used for profanity checking
var Filter *profanityFilter

type profanityFilter struct {
	SwearFilter   *swearfilter.SwearFilter
	ProfanityList []string
}

func init() {
	if Filter == nil {
		badWords := getListOfBadwords()
		swearFilter := swearfilter.New(true, false, false, false, false, badWords...)

		Filter = &profanityFilter{
			SwearFilter:   swearFilter,
			ProfanityList: badWords,
		}
	}
}

// ProfanityCheck checks if a message contains blacklisted words
func ProfanityCheck(word string) (bool, []string, error) {
	isSwearFound, swearsFound, err := Filter.SwearFilter.Check(word)

	return isSwearFound, swearsFound, err
}

func FilterSentence(word string) string {

	var wordRmove string
	word = strings.ReplaceAll(word, "İ", "I")
	_, badWordFound, _ := ProfanityCheck(strings.ToLower(word))

	var star string

	wordRmove = word
	for k, v := range badWordFound {
		for c := 0; c < len(v); c++ {
			star += "*"
		}

		re := regexp.MustCompile(`(?i)` + badWordFound[k])
		wordRmove = re.ReplaceAllString(wordRmove, star)
		star = ""
	}

	return wordRmove
}

func FilterUserName(w string) bool {

	words := getListOfBadwords()

	for _, word := range words {
		if word == strings.ToLower(w) {
			return true
		}
	}

	return false
}
