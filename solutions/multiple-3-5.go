package solutions

func GetSum() int {
	var naturalNum int = 1
	var sum int = 0

	for naturalNum < 1000 {
		if naturalNum%3 == 0 || naturalNum%5 == 0 {
			sum += naturalNum
		}
		naturalNum++
	}

	return sum
}
