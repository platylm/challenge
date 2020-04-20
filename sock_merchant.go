package merchant

func sockMerchant(number int, arrayNumber []int) int {
	var matching int
	for i, n := range arrayNumber {
		for j := i + 1; j < len(arrayNumber); j++ {
			if n == arrayNumber[j] {
				if arrayNumber[i] != 0 {
					matching++
					arrayNumber[i] = 0
					arrayNumber[j] = 0
				}
			}
		}
	}
	return matching
}
