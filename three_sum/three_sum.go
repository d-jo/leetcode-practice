package three_sum

func ThreeSum(nums []int) [][]int {
	// fix one number and solve the two number problem??

	var dataMap = make(map[int]int)
	var result [][]int
	var i int = 0
	var j int = 1
	var k int = 2

	for x := 0; x < len(nums); x++ {
		dataMap[nums[x]] = x
	}

}
