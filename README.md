# Библиетка для анализа русского текста

[![GoDoc](https://godoc.org/github.com/pahanini/go-ru-lex.svg/stm?status.svg)](https://godoc.org/github.com/pahanini/go-ru-lex)
[![Build Status](https://travis-ci.org/pahanini/go-ru-lex.svg)](https://travis-ci.org/pahanini/go-ru-lex)

Библиотека в разработке и может пока выдавать неверные результаты, например, разбиение на слоги на слова "подъезд"

## Разбиение слов на слоги

Правила переноса слов и правила разбиения на слоги в русском языке отличаются.
Эта библиотека предназначена именно для разбиения слов на слоги и не годится для расстановки переносов!

```go
s, err := SyllablizeString("уезжать", "-")
fmt.Println("s") // у-е-зжать
```

## Разбиение на фонемы

```go
pp, err := Phonemize("вьётся")
fmt.Println(pp[0].ToString()) // вь
fmt.Println(pp[1].ToString()) // ё
fmt.Println(pp[2].ToString()) // тс
fmt.Println(pp[3].ToString()) // я
```

