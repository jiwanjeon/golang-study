package main

var out string

func main() {

	test := 2
	switch test {
	case 0:
		out = "zero"
	case 2:
		out = "two"
		fallthrough
	default:
		out = "other"
	}
	println(test)
}
