package rulex

import "strings"

// Syllable
type Syllable struct {
	phonemes []Phoneme
	rule     string
}

// NewSyllable creates new syllable
func NewSyllable() Syllable {
	s := Syllable{}
	s.phonemes = make([]Phoneme, 0)
	return s
}

// ToString returns string with syllable
func (s Syllable) ToString() string {
	var r string
	for _, p := range s.phonemes {
		r += p.letters
	}
	return r
}

// Syllablize splits word to syllables
func Syllablize(word string) ([]Syllable, error) {
	p, err := Phonemize(word)
	if err != nil {
		return nil, err
	}
	ss := make([]Syllable, 0)
	l := len(p)
	s := NewSyllable()
	// hs vowels
	hv := func(pp []Phoneme) bool {
		for _, p := range pp {
			if p.isVowel {
				return true
			}
		}
		return false
	}
	// finish syllable
	f := func(rule string) {
		s.rule = rule
		ss = append(ss, s)
		s = NewSyllable()
	}
	for i := 0; i < l; i++ {
		s.phonemes = append(s.phonemes, p[i])
		// current is vowel & next is vowel too
		// e.g аура
		if i < l-1 && p[i].isVowel && p[i+1].isVowel {
			f("vv")
			continue
		}
		// current is vowel && next is consonant & next next is vowel again
		if i < l-2 && p[i].isVowel && !p[i+1].isVowel && p[i+2].isVowel {
			f("vcv")
			continue
		}
		isNotLastSyllable := hv(p[i+1:])
		// it is not first phoneme,
		// current is voiced but next is not vowel and it is not last syllable
		// eg май-ка, слой-ка
		if i > 0 && i < l-1 && p[i].isSonar && !p[i+1].isVowel && isNotLastSyllable {
			f("sc+")
			continue
		}
		// current is vowel next is consonant but not sonar
		// next next is consonant and is is not last syllable
		if i < l-2 && p[i].isVowel && !p[i+1].isVowel && !p[i+1].isSonar && !p[i+2].isVowel && isNotLastSyllable {
			f("vcc+")
			continue
		}

	}
	f("f")
	return ss, nil
}

// SyllablizeString returns word divided to syllables with separator
func SyllablizeString(word, sep string) (string, error) {
	ss, err := Syllablize(word)
	if err != nil {
		return "", err
	}
	var r []string
	for _, s := range ss {
		r = append(r, s.ToString())
	}
	return strings.Join(r, sep), nil
}
