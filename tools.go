package main

import (
	"net/http"
)
var vowel = []rune("АаӘәҮүҰұӨөОоІіЫыЯяЮюЕеИи")

func getReversed(word []rune) {
	n := len(word)
	for i := 0; i < n / 2; i += 1 {
		word[i], word[n - i - 1] = word[n - i - 1], word[i]
	}
}

func isVowel(x rune) bool {
	for i := 0; i < len(vowel); i += 1 {
		if vowel[i] == x {
			return true
		}
	}
	return false
}

func parse(word []rune) []rune {
	n := len(word)
	result := make([]rune, 0)
	used := make([]bool, n)
	positions := make([]int, 0)
	getReversed(word)
	for i := 0; i < n; i += 1 {
		if isVowel(word[i]) {
			used[i] = true
			positions = append(positions, i)
		} else
		if (word[i] == 'у' || word[i] == 'У') && i + 1 < n && !isVowel(word[i + 1]) {
			used[i] = true
			positions = append(positions, i)
		}
	}
	for i := 0; i < len(positions); i += 1 {
		l := positions[i]
		r := l
		if r + 1 < n && !used[r + 1] && word[r + 1] != 'ң' {
			r += 1
		}
		if r == n - 2 && !used[n - 1] {
			r = n - 1
		}
		for l - 1 >= 0 && !used[l - 1] {
			l -= 1
		}
		for j := l; j <= r; j += 1 {
			used[j] = true
			result = append(result, word[j])
		}
		result = append(result, '-')
	}

	if len(result) > 0 && result[0] == '-' {
		result = result[1:]
	}
	if len(result) > 0 && result[len(result) - 1] == '-' {
		result = result[0:len(result) - 1]
	}
	for i := 0; i + 1 < len(result); i += 1 {
		if result[i] == '-' && result[i + 1] == '-' {
			copy(result[i:], result[i + 1:])
			result = result[:len(result) - 1]
			i -= 1
			continue
		}
	}
	getReversed(result)
	return result
}

func getSyllables(w string, wr http.ResponseWriter) {
	s := []rune(w)
	words := make([][]rune, 0)
	cur := make([]rune, 0)
	for i := 0; i < len(s); i += 1 {
		if s[i] != ' ' {
			cur = append(cur, s[i])
		} else {
			if len(cur) > 0 {
				words = append(words, parse(cur))
				cur = make([]rune, 0)
			}
		}
	}

	if len(cur) > 0 {
		words = append(words, parse(cur))
	}

	for i := 0; i < len(words); i += 1 {
		wr.Write([]byte(string(words[i]) + " "))
	}
}
