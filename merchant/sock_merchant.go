package merchant

func sockMerchant(number int, arrayNumber []int) int {
	var matching int
	for oneSideSock, n := range arrayNumber {
		for pairSock := oneSideSock + 1; pairSock < len(arrayNumber); pairSock++ {
			if n == arrayNumber[pairSock] {
				if arrayNumber[oneSideSock] != 0 {
					matching++
					arrayNumber[oneSideSock] = 0
					arrayNumber[pairSock] = 0
				}
			}
		}
	}
	return matching
}
