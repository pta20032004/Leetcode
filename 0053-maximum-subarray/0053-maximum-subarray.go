// Giải thuật với O(n^2)  rất đơn giản
// func maxSubArray(nums []int) int {
//     sum := nums[0]

//     for i:=0; i<len(nums); i++ {
//         temp := nums[i] 
//         if temp > sum {
//             sum = temp
//         }
//         for j:=i+1; j<len(nums); j++ {
//             temp += nums[j]
//             if temp > sum {
//                 sum = temp
//             }
//         }
        
//     }

//     return sum
// }


func maxSubArray(nums []int) int {

    if len(nums) == 1 {
        return nums[0]
    }
    maxSum := -100000
    currentSum := 0 

    for i:=0; i<len(nums); i++ {
        
        if currentSum < 0 {
            currentSum = nums[i]
        } else {
            currentSum += nums[i]
        }
        
        

        if currentSum > maxSum {
            maxSum = currentSum
        }
    }

    return maxSum
}