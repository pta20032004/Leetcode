//min(n + 1) = min(  min(n), min(n - 1))
//Nếu min(n) == min(n - 1) thì ưu tiên chọn min(n)
func minCostClimbingStairs(cost []int) int {
    n := len(cost) - 1
    a := cost[0]
    b := cost[1]

    for i:=2; i<=n; i++ {
        temp := a
        a = b
        b = min(temp, b) + cost[i]
    }

    return min(a, b)
}

func min(a int, b int) int {
    if a < b {
        return a
    }
    return b
}