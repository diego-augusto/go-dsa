package deletethemiddlenodeofalinkedlist

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		desc  string
		input *ListNode
	}{
		{
			desc: "",
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			deleteMiddle(tc.input)
		})
	}
}
