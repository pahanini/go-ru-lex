package rulex

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestPhonemize(t *testing.T) {

	p, err := Phonemize("мир")
	require.NoError(t, err)
	require.Equal(t, 3, len(p))
	require.Equal(t, "м", p[0].letters)
	require.Equal(t, "и", p[1].letters)
	require.Equal(t, "р", p[2].letters)

	p, err = Phonemize("вьётся")
	require.NoError(t, err)
	require.Equal(t, 4, len(p))
	require.Equal(t, "вь", p[0].letters)
	require.Equal(t, "ё", p[1].letters)
	require.Equal(t, "тс", p[2].letters)
	require.Equal(t, "я", p[3].letters)

	p, err = Phonemize("ванька")
	require.NoError(t, err)
	require.Equal(t, 5, len(p))
	require.Equal(t, "в", p[0].letters)
	require.Equal(t, "а", p[1].letters)
	require.Equal(t, "нь", p[2].letters)
	require.True(t, p[2].isSonar)
	require.Equal(t, "к", p[3].letters)
	require.Equal(t, "а", p[4].letters)

	p, err = Phonemize("уезжать")
	require.NoError(t, err)
	require.Equal(t, 5, len(p))
	require.Equal(t, "у", p[0].letters)
	require.Equal(t, "е", p[1].letters)
	require.Equal(t, "зж", p[2].letters)
	require.Equal(t, "а", p[3].letters)
	require.Equal(t, "ть", p[4].letters)

	_, err = Phonemize("Test")
	require.Error(t, err)

	_, err = Phonemize("Тест")
	require.Error(t, err)
}
