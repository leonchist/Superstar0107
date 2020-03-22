package iteration

func Repeat(character string, count int) string {
	var repetead string

	for i := 0; i < count; i++ {
		repetead += character
	}

	return repetead
}
