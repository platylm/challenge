package jumpingcloud

func jumpingOnClouds(cloud []int) int {
	var count int

	for current := 0; current < len(cloud); {
		if current < len(cloud)-2 && cloud[current] == cloud[current+2] {
			count++
			current = current + 2
			continue
		}
		if current < len(cloud)-1 && cloud[current] == cloud[current+1] {
			count++
			current = current + 1
			continue
		}
		current++
	}
	return count
}
