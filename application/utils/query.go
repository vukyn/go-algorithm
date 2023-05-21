package utils

// Where function filter elements based on conditions.
// Where return a new slice.
//
//	 Example:
//		evens := Filter(list, func(i int) bool {return i % 2 == 0})
func Where[T any](list []T, f func(T) bool) []T {
	var newList []T
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// Find function find first element based on conditions.
// Find return a new value T.
//
//	 Example:
//		item := Find(list, func(n name) bool {return n == "ABC"})
func Find[T any](list []T, f func(T) bool) T {
	var newValue T
	for _, v := range list {
		if f(v) {
			newValue = v
			break
		}
	}
	return newValue
}

// Remove function remove current element based on index.
// Remove return current slice without removed element.
//
//	 Example:
//		items := Remove(items, 1)
func Remove[T any](list []T, i int) []T {
	list[i] = list[len(list)-1]
	return list[:len(list)-1]
}

// IndexOf function return an index in a list.
//
//	 Example:
//		index := IndexOf(list, func(n name) bool {return n == "ABC"})
func IndexOf[T any](list []T, f func(T) bool) int {
	index := -1
	for i, v := range list {
		if f(v) {
			index = i
			break
		}
	}
	return index
}

// Distinct function remove duplicates from slice.
// Distinct return a new slice.
//
//	 Example:
//		newList := Distinct(oldList)
func Distinct[T string | int32 | int64 | float32 | float64](list []T) []T {
	allKeys := make(map[T]bool)
	newList := []T{}
	for _, item := range list {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			newList = append(newList, item)
		}
	}
	return newList
}