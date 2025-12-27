package utils

// GenerateFromMask returns a receive-only channel that yields password
// candidates produced by substituting the '?' placeholders in `mask` with
// characters from `charset`. The function runs generation in a separate
// goroutine and closes the channel when all combinations are exhausted.
func GenerateFromMask(mask string, charset string) <-chan string {
	out := make(chan string, 100)

	go func() {
		defer close(out)

		var positions []int
		for i, c := range mask {
			if c == '?' {
				positions = append(positions, i)
			}
		}

		indexes := make([]int, len(positions))

		for {
			runes := []rune(mask)
			for i, pos := range positions {
				runes[pos] = rune(charset[indexes[i]])
			}

			out <- string(runes)

			for i := len(indexes) - 1; i >= 0; i-- {
				indexes[i]++
				if indexes[i] < len(charset) {
					break
				}
				indexes[i] = 0
				if i == 0 {
					return
				}
			}
		}
	}()

	return out
}
