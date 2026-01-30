// func climbStairs(n int) int {
//     a := 0 //biến để đếm
//     tong := 0
//     helper(&a, tong, n)
//     return a
// }

// func helper(a *int, tong int, n int) { // Truyền vào con trỏ của biến a
//     if tong == n {
//         *a += 1
//     }

//     if tong > n {
//         return
//     }

//     helper(a, tong + 1, n)
//     helper(a, tong + 2, n)
// }

// Toi uu, 1 + 2 lên 3, 2 + 1 cũng lên 3, lúc này 3 có ở hai nhánh và nó lại phải chọn 1 hai lần, chọn 2 2 lần

// Số cách đi tới n bằng số cách đi tới (n - 1) + số cách đi tới (n - 2)
func climbStairs(n int) int {
    a := 1
    b := 2
    
    if n <= 2 {
        return n
    }

    for i:=3; i<n; i++ {
        temp:=a
        a = b
        b = b + temp
    } 

    return a + b

}