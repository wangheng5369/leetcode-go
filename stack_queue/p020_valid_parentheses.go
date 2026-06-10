package stackqueue

// p020 有效的括号
// 给定一个只包括 '(', ')', '{', '}', '[', ']' 的字符串 s，判断字符串是否有效。
// 有效字符串需满足：左括号必须用相同类型的右括号闭合，左括号必须以正确的顺序闭合。
//
// 示例 1：
// 输入：s = "()"
// 输出：true
//
// 示例 2：
// 输入：s = "()[]{}"
// 输出：true
//
// 提示：
// - 1 <= s.length <= 10^4
// - s 只包含括号字符

func isValid(s string) bool {
	stack := []rune{}
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else if c == ')' {
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		} else if c == '}' {
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		} else if c == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
