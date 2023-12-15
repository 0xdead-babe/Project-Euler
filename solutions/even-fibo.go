package solutions

func GetEvenFiboSum() int {
	const fourMillion int = 4000_000
	var start1 int = 1
	var start2 int = 2

	var sum int = 2

	var nextTerm int = 0

	for nextTerm <= fourMillion {
		nextTerm = start1 + start2
		start1 = start2
		start2 = nextTerm

		if nextTerm%2 == 0 {
			sum += nextTerm
		}

	}

	return sum
}
