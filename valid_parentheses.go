package godsa

import "github.com/emirpasic/gods/stacks/arraystack"

//()[]{}

func IsValidParentheses(s string) bool {
	st := arraystack.New()

	g := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	for _, l := range s {
		if l == rune(')') || l == rune(']') || l == rune('}') {
			if v, ok := st.Pop(); ok {
				if pv, ok := v.(string); ok && pv != g[string(l)] {
					return false
				}
			}
		} else {
			st.Push(string(l))
		}
	}

	return st.Empty()
}
