package adder


var default_add = 1

// function documentation
func internalAdd(n int) int {
	return n + default_add
}

func Addnum(n int) int {
	return internalAdd(n)
}