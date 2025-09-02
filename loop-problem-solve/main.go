package main

func main() {
	values := []string{"a", "b", "c"}
	var prints []func()
	for _, v := range values {
		_v := v
		prints = append(prints, func() {
			println(v)
			println(_v)
		})
	}

	for _, print := range prints {
		print()
	}
}
