func removeDuplicates(nums []int) int {
    if len(nums) == 0 || len(nums) == 1 {
        return len(nums)
    }

    count := 1
    temp := nums[0]
    currentIndex := 1

    for i:=0; i<len(nums); i++ {
        if temp != nums[i] {
            temp = nums[i]
            count++
            nums[currentIndex] = temp
            currentIndex++
        }
    }
    return count
}