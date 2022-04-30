package largest_rect_hist

type Node struct {
	height int
	pos    int
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
			topIdx := len(stack) - 1
			height := stack[topIdx].height
			pos := stack[topIdx].pos
			area := height * (idx - pos)
			if result < area {
				result = area
			}
			node.pos = pos
			stack = stack[0:topIdx]
		}
		stack = append(stack, node)
	}
	max_idx := len(heights) - 1
	for len(stack) > 0 {
		topIdx := len(stack) - 1
		height := stack[topIdx].height
		pos := stack[topIdx].pos
		area := height * (max_idx - pos + 1)
		if result < area {
			result = area
		}
		stack = stack[0:topIdx]
	}

	return result
}
