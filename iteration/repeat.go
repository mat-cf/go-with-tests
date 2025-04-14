package iteration

const repeatCount = 5

func Repeat(character string) string {
	var repeated string
	// modern syntax
	for range repeatCount {
		repeated += character
	}
	return repeated
}
