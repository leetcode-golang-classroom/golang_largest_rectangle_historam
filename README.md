# golang_largest_rectangle_histogram

Given an array of integers `heights` representing the histogram's bar height where the width of each bar is `1`, return *the area of the largest rectangle in the histogram*.

# Examples

Example 1:

![](https://i.imgur.com/kLFmYVy.jpg)

Input: heights = [2,1,5,6,2,3]
Output: 10
Explanation: The above is a histogram where width of each bar is 1.
The largest rectangle is shown in the red area, which has an area = 10 units.
Example 2:

![](https://i.imgur.com/3n58Yls.jpg)

Input: heights = [2,4]
Output: 4

Constraints:

1 <= heights.length <= $10^5$
0 <= heights[i] <= $10^4$

## 解析

要算出連續 height 所組成的最大矩型面積

所以關鍵在於怎麼列舉出所有可能

讓我們思考一下

每個 height 能夠成為矩型面積的 height 假設後面的 height 都大於當下的 height

如下圖

![](https://i.imgur.com/h9UYGoY.png)

因為每個矩陣只要考慮最接近的height 恰巧可以使用 stack 後進先出的特性

所以可以透過一個 stack 紀錄每個 height 以其對應的起始位置 index 

每次遇到比當下 height 還要第高則需要 pop 出來並且把該 height 以及其起始位置與當下位置 計算出以height 所形成的 矩型面積

而如果比當下 height 小則 push 進入 stack

舉例如下：

![](https://i.imgur.com/CxhLbQ1.png)

![](https://i.imgur.com/jpAEzqc.png)

![](https://i.imgur.com/QWwhKTB.png)

這樣每次只要輪詢一次每個結點 

而 stack 最長只有 n 個 因此時間複雜度是 O(n) 空間複雜度也是 O(n)

## 程式碼

```go
type Node struct {
    height int
    pos int
}
func largestRectangleArea(heights []int) int {
    result := 0
    if len(heights) == 1 {
        return heights[0]
    }
    stack := []Node{} // store min index
    for idx, value := range heights {
        node := Node{height: value, pos: idx}
        for len(stack) > 0 && stack[len(stack)-1].height > value {
            topIdx := len(stack)-1
            height := stack[topIdx].height
            pos := stack[topIdx].pos
            area := height * (idx - pos)
            if result < area {
                result = area
            }
            stack = stack[0: topIdx]
            node.pos = pos
        }
        stack = append(stack, node)           
    }
    max_idx := len(heights) - 1
    for len(stack) > 0 {
        topIdx := len(stack)-1
        height := stack[topIdx].height
        pos := stack[topIdx].pos
        area := height * (max_idx - pos + 1)
        if result < area {
            result = area
        }
        stack = stack[0: topIdx]
    }
    
    return result
}
```

## 困難點

1. 要想出形成矩陣高度的條件
2. 要知道怎麼做到只檢查鄰近比自己高的長度

# Solve points

- [x]  Understand what is problem request
- [x]  what Data structure could help solving this problem
- [x]  understand how data structure could use to help