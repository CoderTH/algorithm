package LinearSearch

import "reflect"

func Search(data []interface{}, target interface{}) int {
	for i := 0; i < len(data); i++ {
		if reflect.DeepEqual(data[i], target) {
			return i
		}
	}
	return -1
}
