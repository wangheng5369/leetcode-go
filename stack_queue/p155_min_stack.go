package stackqueue

import "math"

// p155 最小栈
// 设计一个支持 push、pop、top 操作，并能在常数时间内检索到最小元素的栈。
// - push(x) —— 将元素 x 推入栈中
// - pop() —— 删除栈顶元素
// - top() —— 获取栈顶元素
// - getMin() —— 检索栈中的最小元素
//
// 提示：
// - -2^31 <= val <= 2^31 - 1
// - pop、top 和 getMin 操作都在 O(1) 时间复杂度内

type MinStack struct {
	stack  []int
	minVal int
}

func Constructor() MinStack {
	return MinStack{
		stack:  make([]int, 0),
		minVal: 0,
	}
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 {
		this.minVal = val
		this.stack = append(this.stack, 0)
	} else {
		// 溢出检测
		var diff int
		if val > 0 && this.minVal > math.MaxInt-val {
			// val + this.minVal 会溢出，改存实际值
			diff = val - this.minVal
		} else if val < 0 && this.minVal < math.MinInt-val {
			// val + this.minVal 会溢出
			diff = val - this.minVal
		} else {
			diff = val - this.minVal
		}
		// 其实上面的检测对于 int 的范围来说...
		this.stack = append(this.stack, diff)
		if diff < 0 {
			this.minVal = val
		}
	}
}

func (this *MinStack) Pop() {
	top := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if top < 0 {
		this.minVal -= top
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1] + this.minVal
}

func (this *MinStack) GetMin() int {
	return this.minVal
}
