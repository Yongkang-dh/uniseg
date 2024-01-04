package uniseg

import "unicode/utf8"

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
