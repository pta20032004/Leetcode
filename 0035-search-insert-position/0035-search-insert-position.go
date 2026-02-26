import "fmt"
func searchInsert(nums []int, target int) int {
   head := 0;
   tail := len(nums) - 1;
   var mid int;
   for {
        mid = (head + tail) / 2
        if nums[mid] == target {

            return mid
        }

        if tail - head <= 1 {
            if target > nums[tail] {

                return tail + 1
            } else if (target < nums[head]) {

                return head 
            } else {

                return tail
            }
        }

        if nums[mid] < target {
            head = mid
        } else {
            tail = mid
        }
   } 
   return -1;
}