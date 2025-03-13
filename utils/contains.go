package utils

func Contains(slice []uint, value uint) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}