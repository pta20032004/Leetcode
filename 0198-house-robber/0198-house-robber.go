// output max(max(n), max(n - 1))
// max(n) = max(max(n - 2), max(n - 3)) + n
func rob(nums []int) int {
    if len(nums)<=1 {
        return nums[0]
    } else if len(nums) <=2 {
        return max(nums[0], nums[1])
    } else if len(nums) <= 3 {
        return max(nums[0] + nums[2], nums[1])
    }
    var profits [4]int
    n := 3
    profits[n - 3] = nums[0]
    profits[n - 2] = nums[1]
    profits[n - 1] = nums[2] + nums[0]
    profits[n] = max(profits[n - 2], profits[n - 3]) + nums[3]

    for i:=4; i<=len(nums) - 1; i++ {
       
        profits[n - 3] = profits[n - 2]
        profits[n - 2] = profits[n - 1]
        profits[n - 1] = profits[n]

        profits[n] = max(profits[n - 2], profits[n - 3]) + nums[i]
    } 

    return max(profits[n], profits[n - 1])
}

func max(a int, b int) int {
    if a > b {
        return a
    }

    return b
}