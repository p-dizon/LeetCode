package containermostwater

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	mArea := calculateWaterArea(height, 0, len(height)-1)
	for i, j := 0, len(height)-1; i < j; {
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
		mArea = max(mArea, calculateWaterArea(height, i, j))
	}
	return mArea
}

func calculateWaterArea(h []int, i, j int) int {
	var height int
	if h[i] < h[j] {
		height = h[i]
	} else {
		height = h[j]
	}
	width := j - i
	area := width * height
	if area < 0 {
		return -area
	}
	return area
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
