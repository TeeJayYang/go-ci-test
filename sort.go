package sort

func SortBubble(input []int) {
	if input == nil {
		return
	}
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if input[i] < input[j] {
				temp := input[i]
				input[i] = input[j]
				input[j] = temp
			}
		}
	}
}
