package rulex

import (
	"errors"
	"regexp"
)

const (
	vowel     = `аеёиоуыэюя`
	modifier = `ьъ`
	sonar     = "йлмнр"
)

// Phoneme
type Phoneme struct {
	letters     string
	isVowel     bool
	isSonar     bool
	hasModifier bool
}

var validWord = regexp.MustCompile("^[а-яё]{1,}$")

// Phonemize splits word into phonemes
//
// List of phonemes:
// - тьс, тс
// - зж, дж
// - тт, нн, cc
// - ссч
// - *ь, *ъ
func Phonemize(word string) ([]Phoneme, error) {
	if !validWord.MatchString(word) {
		return nil, errors.New("invalid word, lower russian text expected")
	}
	pp := make([]Phoneme, 0)
	rr := []rune(word)
	l := len(rr)
	for i := 0; i < l; i++ {
		if ok, p := HasAnyPrefix(string(rr[i:]), []string{"тьс", "тс", "зж", "дж", "тт", "нн", "сс", "ссч"}); ok {
			i += len([]rune(p)) - 1
			pp = append(pp, Phoneme{letters: p})
			continue
		}
		if i < l-1 && Is(rr[i+1], modifier) {
			pp = append(pp, Phoneme{
				letters: string(rr[i : i+2]),
				isVowel:  Is(rr[i], vowel),
				isSonar:  Is(rr[i], sonar),
				hasModifier: true,
			})
			i++
			continue
		}
		pp = append(pp, Phoneme{
			letters:  string(rr[i]),
			isVowel:  Is(rr[i], vowel),
			isSonar:  Is(rr[i], sonar),
		})
	}
	return pp, nil
}
