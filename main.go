package main

func main() {
	//fmt.Println("main called")

	/*m := make(map[string]func(int, bool) int)
	m["compute"] = Compute
	return m*/
}

func Compute(a int, b bool) int {
	if b {
		return a * 2
	}
	return a - 1
}
