package day01

func partOne(data []int) int {
	count := 0
	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			count++
		}
	}
	return count
}

func partTwo(data []int) int {
	list := []int{}
	count := 0

	for i := 2; i < len(data); i++ {
		list = append(list, data[i]+data[i-1]+data[i-2])
	}

	for i := 1; i < len(list); i++ {
		if list[i] > list[i-1] {
			count++
		}
	}
	return count
}
