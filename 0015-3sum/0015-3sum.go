import "sort"

func threeSum(nums []int) [][]int {
    sort.Ints(nums)    
    out := [][]int{}
    i := 0
    for {
        if i >= len(nums) - 1 {
            break
        }
        head := i+1
        tail := len(nums)-1
        for {
            if (head >= tail) {
                break
            }


            sum := nums[i] + nums[head] + nums[tail]
            if sum == 0 {
                out = append(out, []int{nums[i], nums[head], nums[tail]})
                next(&head, nums, nums[head])
                down(&tail, nums, nums[tail])
            } else if sum < 0 {
                next(&head, nums, nums[head])
            } else {
                down(&tail, nums, nums[tail])
            } 
        }

        if nums[i] == nums[i + 1] {
        next(&i, nums, nums[i])} else {
            i++
        }
    }

    return out
}

func next(i *int, nums []int, value int){
    *i++
    for {
        if *i >= len(nums) || nums[*i] != value {
            break
        }
        *i++
    }
}

func down(i *int, nums []int, value int){
    *i--
    for {
        if *i < 0 || nums[*i] != value {
            break
        }
        *i--
    }
}