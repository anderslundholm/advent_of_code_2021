package day06

// First itteration, save all fish in a list, does not scale to higher numbers of days (150+).
func simulateFish(fishList []int, days int) int {
	for i := 0; i < days; i++ {
		tempList := fishList
		for i, fish := range tempList {
			if fish == 0 {
				tempList[i] = 6
				tempList = append(tempList, 8)
			} else {
				tempList[i]--
			}
		}
		fishList = tempList
	}

	return len(fishList)
}

// Second itteration, counts the amount of fish of each age instead of saving all fish.
func simulateFish2(fishList []int, days int) int {
	var fishAges = make(map[int]int)
	var spawnFish, sum int

	for _, age := range fishList {
		fishAges[age] += 1
	}

	for i := 0; i < days; i++ {
		spawnFish = fishAges[0]
		for j := 0; j < 8; j++ {
			fishAges[j] = fishAges[j+1]
		}
		fishAges[6] += spawnFish
		fishAges[8] = spawnFish
	}

	for _, fish := range fishAges {
		sum += fish
	}

	return sum
}
