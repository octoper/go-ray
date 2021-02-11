package utils

// dfs assumes its input as a slice and walks it depth-first
func dfs(in interface{}, out *[]interface{}) {
	s := in.([]interface{})
	for _, e := range s {
		if e != nil {
			switch v := e.(type) {
			case int:
				*out = append(*out, v)
			case []interface{}:
				dfs(v, out)
			}
		}
	}
}

func Flatten(l interface{}) []interface{} {
	var r []interface{}
	dfs(l, &r)
	return r
}