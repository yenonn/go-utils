package collection

func Contains[T int | string | float64](collection []T, found T) bool {
	for _, value := range collection {
		if value == found {
			return true
		}
	}
	return false
}

func Pop[T int | string | float64](collection []*T) {
	totalLength := len(collection)
	lastIndex := totalLength - 1

}

func Remove[T int | string | float64](collection []*T, index int) {
}
