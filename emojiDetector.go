package uniseg

import (
	"unicode/utf8"
)

func EmojiCount(str string) int {
	if len(str) == 0 {
		return 0
	}

	emjCnt, state := 0, -1
	c := ""
	for len(str) > 0 {
		c, str, _, state = StepString(str, state)
		r, _ := utf8.DecodeRuneInString(c)
		if property(graphemeCodePoints, r) == prExtendedPictographic ||
			property(graphemeCodePoints, r) == prRegionalIndicator {
			emjCnt += 1
		}
	}
	return emjCnt
}

func ContinuousEmojiCount(str string) int {
	if len(str) == 0 {
		return 0
	}

	continuous, state, maxContinue := 0, -1, 0
	c := ""
	for len(str) > 0 {
		c, str, _, state = StepString(str, state)
		r, _ := utf8.DecodeRuneInString(c)
		if property(graphemeCodePoints, r) == prExtendedPictographic ||
			property(graphemeCodePoints, r) == prRegionalIndicator ||
			property(graphemeCodePoints, r) == prEmojiPresentation {
			continuous += 1
			maxContinue = maxInt(maxContinue, continuous)
		} else if c != " " {
			continuous = 0
		}

	}

	return maxContinue
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
