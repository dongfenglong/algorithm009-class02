package Week_04

func lemonadeChange(bills []int) bool {
	if bills[0] != 5{
		return false
	}

	five, ten := 0, 0

	for _, bill := range bills{
		switch bill{
		case 5:
			five++
		case 10:
			if five > 0{
				five--
				ten++
			}else{
				return false
			}
		case 20:
			if ten > 0 && five > 0{
				ten--
				five--
				continue
			}else if five >=3{
				five -= 3
				continue
			}else{
				return false
			}
		}
	}
	return true
}