package leetcode

import "container/list"

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
func isValid(s string) bool {
	list := list.New()

	for i := 0; i < len(s); i++ {
		b := s[i]

		switch b {
		case '(', '[', '{':
			list.PushBack(b)
		case ')':
			element := list.Back()
			if element == nil {
				return false
			}

			if element.Value.(byte) == '(' {
				list.Remove(element)
				continue
			} else {
				return false
			}
		case ']':
			element := list.Back()
			if element == nil {
				return false
			}
			if element.Value.(byte) == '[' {
				list.Remove(element)
				continue
			} else {
				return false
			}

		case '}':
			element := list.Back()
			if element == nil {
				return false
			}
			if element.Value.(byte) == '{' {
				list.Remove(element)
				continue
			} else {
				return false
			}
		}
	}

	return list.Len() == 0
}
