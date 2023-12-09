package longestconsecutivesequence

func longestConsecutive(nums []int) int {

	//create a set
	set := make(map[int]struct{})
	for _, n := range nums {
		set[n] = struct{}{}
	}

	longest := 0
	for _, n := range nums {

		//if number does not have a previous, it is the beginning of the sequence
		if _, ok := set[n-1]; !ok {
			count := 1

			//try to find next number of the sequence
			for _, ok := set[n+count]; ok; _, ok = set[n+count] {
				count++
			}

			if count > longest {
				longest = count
			}

			//if the current sequence is greather than half of len(nums),
			//the next sequence cant be greather than current one.
			if count > len(nums)/2 {
				break
			}
		}
	}

	return longest
}
