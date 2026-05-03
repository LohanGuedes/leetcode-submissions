func twoSum(nums []int, target int) []int {
    items := map[int]int{}

    for i, num := range nums {
        items[num] = i
    }

    for i, num := range nums {
        try := target - num
        j, ok := items[try]
        if i == j {
            continue
        }
        if ok {
            return []int{i, j}
        }
    }

    return []int{}
}
