func maxProfit(prices []int) int {
    min := prices[0]
    max := -1
    index := -1
    value := 0
    for i:=0; i<len(prices); i++ {
        if prices[i] <= min {
            min = prices[i]
            index = i
        }

        if prices[i] >= max && i > index && index  != -1 && prices[i] > min {
            max = prices[i]
            if max - min > value {
                value = max - min
            }
        }

        if prices[i] - min > value {
            max = prices[i]
            value = prices[i] - min
        }
    }

    return value

    
}