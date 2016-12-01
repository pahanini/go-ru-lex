package rulex

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const sep = "-"

func joinSyllables(ss []Syllable) string {
	var r []string;
	for _, s := range ss {
		r = append(r, s.ToString())
	}
	return strings.Join(r, sep);
}

func joinSyllablesAndRules(ss []Syllable) string {
	var r []string;
	for _, s := range ss {
		r = append(r, s.ToString() + "(" + s.rule + ")")
	}
	return strings.Join(r, sep);
}

func TestSyllablize(t *testing.T) {
	examples := []string{
		"а-ист",
		"а-у-ра",
		"бул-ка",
		"ве-сна",
		"вьё-тся",
		"и-зжить",
		"кла-ссный",
		"кор-ка",
		"конь-ки",
		"ло-па-та",
		"метр",
		"ми-ни-а-тю-ра",
		"мо-шка",
		"мы-ться",
		"о-ття-нуть",
		//"подъ-езд",
		"поль-ка",
		"по-чта",
		"по-э-ма",
		"ра-ссчи-та-нный",
		"слой-ка",
		"со-лом-ка",
		"со-ци-аль-на-я",
		"то-чка",
		"у-е-зжать",
	}
	for _, example := range examples {
		res, err := Syllablize(strings.Replace(example, "-", "", -1))
		assert.NoError(t, err)
		assert.Equal(t, example, joinSyllables(res), joinSyllablesAndRules(res))
	}
}

func TestSyllablizeString(t *testing.T) {
	s, _ := SyllablizeString("уезжать", "-")
	assert.Equal(t, "у-е-зжать", s)
}
