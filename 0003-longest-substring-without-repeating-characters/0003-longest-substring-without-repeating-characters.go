import "fmt"
func lengthOfLongestSubstring(s string) int {

    curLen := 0
    maxLen := 0
    mapper := make(map[byte]int)
    curIndex := 0

    for i:=0; i<len(s); i++ {
        v, ok := mapper[s[i]]

        if ok  {
            if v >= curIndex {
            curLen = i - v
            curIndex = v + 1
            } else {
                curLen++
            }
        } else{
            curLen ++
        }

        mapper[s[i]] = i

        if curLen > maxLen {
            maxLen = curLen
        }

    }   
    return maxLen 
}