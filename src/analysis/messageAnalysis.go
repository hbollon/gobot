package analysis

import (
	Yaml "gobot/src/yaml"
)

// Return most high matching response in function of the received message
func FindResponse(m string) string {
	var bestMatchingScore float64
	var currentResponse string

	for _, template := range Yaml.ResponsesPool.Templates {
		for _, message := range template.Messages {
			percent := matchingPercentage(m, message)
			if bestMatchingScore != 0 {
				if percent >= 0.50 && percent > bestMatchingScore {
					bestMatchingScore = percent
					currentResponse = template.Response
				}
			} else {
				if percent >= 0.50 {
					bestMatchingScore = percent
					currentResponse = template.Response
				}
			}
		}
	}

	if currentResponse != "" {
		return currentResponse
	}
	return Yaml.ResponsesPool.DefaultResponse
}

func matchingPercentage(str1, str2 string) float64 {
	// Get Levenshtein distance between these two strings
	var distance = levenshteinDistance(str1, str2)

	// Compare strings lenght and make a matching percentage between them
	if len(str1) >= len(str2) {
		return float64(len(str1)-distance) / float64(len(str1))
	}
	return float64(len(str2)-distance) / float64(len(str2))

}

func levenshteinDistance(str1, str2 string) int {
	// Convert string parameters to rune arrays to be compatible with non-ASCII
	runeStr1 := []rune(str1)
	runeStr2 := []rune(str2)

	// Get and store lenght of these strings
	runeStr1len := len(runeStr1)
	runeStr2len := len(runeStr2)
	if runeStr1len == 0 {
		return runeStr2len
	} else if runeStr2len == 0 {
		return runeStr1len
	}

	column := make([]int, runeStr1len+1)

	for y := 1; y <= runeStr1len; y++ {
		column[y] = y
	}
	for x := 1; x <= runeStr2len; x++ {
		column[0] = x
		lastkey := x - 1
		for y := 1; y <= runeStr1len; y++ {
			oldkey := column[y]
			var i int
			if runeStr1[y-1] != runeStr2[x-1] {
				i = 1
			}
			column[y] = min(column[y]+1, column[y-1]+1, lastkey+i)
			lastkey = oldkey
		}
	}

	return column[runeStr1len]
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}

	return c
}
