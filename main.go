package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
    charIndexMap := make(map[byte]int) 
    maxLength := 0                     
    start := 0                      

    for end := 0; end < len(s); end++ {
        currentChar := s[end]           
        if index, ok := charIndexMap[currentChar]; ok && index >= start {
            start = index + 1           
        }
        charIndexMap[currentChar] = end 
        currentLength := end - start + 1 
        if currentLength > maxLength {
            maxLength = currentLength   
        }
    }

    return maxLength
}

func main() {
    fmt.Println(lengthOfLongestSubstring("abcabcbb")) //* Output: 3
    fmt.Println(lengthOfLongestSubstring("pwwkew"))   //* Output: 3
    fmt.Println(lengthOfLongestSubstring("bbbb"))     //* Output: 1
    fmt.Println(lengthOfLongestSubstring(""))         //* Output: 0
}
