package strutil

func CutOffDots(text string, length int) string {
	return CutOff(text, length, "...")
}

func CutOff(text string, length int, end string) string {
	lenText := len(text)
	lenEnd := len(end)
	if lenText <= length || lenText <= lenEnd {
		return text
	}
	if lenEnd >= length {
		return end
	}
	return text[:length-len(end)] + end
}
