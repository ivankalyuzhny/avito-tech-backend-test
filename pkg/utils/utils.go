package utils

func Unique(slice []string) []string {
	keys := make(map[string]bool)
	list := make([]string, 0)
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func RemoveIntersections(slice1, slice2 []string) ([]string, []string) {
	m := make(map[string]bool)
	for _, s := range slice1 {
		m[s] = true
	}
	res2 := make([]string, 0)
	for _, s := range slice2 {
		if !m[s] {
			res2 = append(res2, s)
		} else {
			m[s] = false
		}
	}
	res1 := make([]string, 0)
	for k, v := range m {
		if v {
			res1 = append(res1, k)
		}
	}
	return res1, res2
}
