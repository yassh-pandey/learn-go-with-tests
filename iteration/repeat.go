package iteration

// RepeatChars repeates a given set of character n no of times
func RepeatChars(char string, no int) string {
	repeatedStr := ""
	for i := 0; i < no; i++ {
		repeatedStr += char
	}
	return repeatedStr
}
