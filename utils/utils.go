package utils

func ExpandStringList(list []interface{}) []string {
	strs := make([]string, len(list))
	for _, s := range list {
		strs = append(strs, s.(string))
	}
	return strs
}
