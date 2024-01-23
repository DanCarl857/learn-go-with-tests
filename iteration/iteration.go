package iteration

const repeatCount = 5

func Iteration(character string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}
