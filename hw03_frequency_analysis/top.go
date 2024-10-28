package hw03frequencyanalysis

import (
	"strings"

	"sort"
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
	var ordered = sortKeys(counter)
	return trim(ordered, n)
}

func sortKeys(counter map[string]int) []string {
	keys := make([]string, 0, len(counter))
	for key := range counter {
		keys = append(keys, key)
	}

	sortLex(keys)
	sortVal(keys, counter)
	return keys
}

// value sorting
func sortVal(keys []string, counter map[string]int) {
	sort.Slice(keys, func(i, j int) bool {
		if counter[keys[i]] == counter[keys[j]] {
			return keys[i] < keys[j]
		}
		return counter[keys[i]] > counter[keys[j]]
	})
}

//Expected :[]string{"он", "а", "и", "ты", "что", "-", "Кристофер", "если", "не", "то"}
//Actual   :[]string{"он", "а", "и", "ты", "что", "то", "Кристофер", "не", "если", "-"}

// lexicographic sorting
func sortLex(keys []string) {
	sort.Strings(keys)
}

func trim(ordered []string, n int) []string {
	var res []string
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
