package collection

func Contains[T int | string | float32](collection []T, found T) bool {
	for _, value := range collection {
		if value == found {
			return true
		}
	}
	return false
}
