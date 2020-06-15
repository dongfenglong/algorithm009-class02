package Week_04
func isPerfectSquare(num int) bool {
	if num < 2{
		return true
	}

	left, right := 2, num/2
	for left <= right{
		pivot := (left+right)/2

		if pivot * pivot == num{
			return true
		}

		if pivot * pivot < num {//当前值偏小
			left = pivot+1
			continue
		}

		if pivot * pivot > num{//当前值偏大
			right = pivot - 1
			continue
		}
	}
	return false
}