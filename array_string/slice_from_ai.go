package array_string

import "fmt"

func SliceTest() {
	// 1. 初始化 mainBuffer
	mainBuffer := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("Initial - main: %v, len: %d, cap: %d\n", mainBuffer, len(mainBuffer), cap(mainBuffer))

	// 2. 提取 subView (索引 2 到 4)
	// TODO: 你的代码
	subView := mainBuffer[2:4]

	// 3. 修改 subView[0] 并 append 500
	// TODO: 你的代码
	fmt.Println("\n--- After first modify and append ---")
	subView[0] = 300
	subView = append(subView, 500)

	// 4. 强制扩容 (循环 append)
	// TODO: 你的代码
	for i := 0; i < 10; i++ {
		subView = append(subView, 1)
	}

	// 5. 再次修改 subView[0] 为 999
	// TODO: 你的代码
	subView[0] = 999
	fmt.Println("\n--- After expansion and second modify ---")

	// 最终打印对比
	fmt.Printf("Final - main: %v\n", mainBuffer)
	fmt.Printf("Final - sub : %v, len: %d, cap: %d\n", subView, len(subView), cap(subView))
}
