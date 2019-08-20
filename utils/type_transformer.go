package utils

func interface2map(i interface{}) map[string]interface{} {
	res_data := i.(map[string]interface{})
	return res_data
}
