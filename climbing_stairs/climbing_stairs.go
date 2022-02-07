package climbingstairs

var stairValues map[int]int = make(map[int]int)

func ClimbStairs(n int) int {

	if p, ok := stairValues[n]; ok {
		return p
	}

	stairValues[0] = 1
	stairValues[1] = 1
	stairValues[2] = 2

	for i := 3; i <= n; i++ {
		if _, ok := stairValues[i]; !ok {
			stairValues[i] = stairValues[i-1] + stairValues[i-2]
		} else {
			continue
		}
	}

	return stairValues[n]
}
