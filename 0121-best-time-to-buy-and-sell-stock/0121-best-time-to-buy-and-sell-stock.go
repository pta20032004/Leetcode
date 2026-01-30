// func maxProfit(prices []int) int {
//     min := prices[0]
//     max := -1
//     index := -1
//     value := 0
//     for i:=0; i<len(prices); i++ {
//         if prices[i] <= min {
//             min = prices[i]
//             index = i
//         }

//         if prices[i] >= max && i > index && index  != -1 && prices[i] > min {
//             max = prices[i]
//             if max - min > value {
//                 value = max - min
//             }
//         }

//         if prices[i] - min > value {
//             max = prices[i]
//             value = prices[i] - min
//         }
//     }

//     return value

// }

//Code tương đối tối ưu về complexity và space nhưng vẫn rất dài dòng và rối
//Tiến hành tối giản
func maxProfit(prices []int) int {
    //Tư duy: Tôi cần tìm ngày có mức giá thấp nhất để mua
    // Tôi cần bán ra vào ngày có profit cao nhất, khi profit cao nhất, tiến hành
    // cập nhật ngày mua, không quan tâm ngày bán

    minDay := prices[0]
    profit := 0
    for i:=0; i<len(prices); i++ {
        if prices[i] <= minDay {
            minDay = prices[i]
        } else {
            if prices[i] - minDay > profit {
                profit = prices[i] - minDay
            }
        }
    }
    return profit
}