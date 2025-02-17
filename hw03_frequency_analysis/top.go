package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(s string) []string {
	// Place your code here.
	slices := sliceStr(s)
	counter := countWords(slices)
	top10 := takeTop(counter, 10)
	return top10
}

func sliceStr(s string) []string {
	return strings.Fields(s) // strings.Split(s, " ") //
}

func countWords(slices []string) map[string]int {
	m := make(map[string]int)
	for _, r := range slices {
		m[r]++
	}
	return m
}

func takeTop(counter map[string]int, n int) []string {
	ordered := sortKeys(counter)
	return trim(ordered, n)
}

func sortKeys(counter map[string]int) []string {
	keys := make([]string, 0, len(counter))
	for key := range counter {
		keys = append(keys, key)
	}

	// value sorting with lexicographic sorting
	sort.Slice(keys, func(i, j int) bool {
		if counter[keys[i]] == counter[keys[j]] {
			return keys[i] < keys[j]
		}
		return counter[keys[i]] > counter[keys[j]]
	})
	return keys
}

func trim(ordered []string, n int) []string {
	res := make([]string, 0, n)
	var count int
	for _, s := range ordered {
		if count >= n {
			break
		}
		res = append(res, s)
		count++
	}
	return res
}
