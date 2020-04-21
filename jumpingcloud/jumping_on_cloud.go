package jumpingcloud

func jumpingOnClouds(cloud []int) int {
	var count int
	var current int
	for index, c := range cloud {
		if index == 0 {
			current = c
			continue
		}

		if current == c {
			count++
			current = c
		}
	}
	return count
}
