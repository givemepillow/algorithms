package rt

func Fibonacci(value int) int {
	mem := make([]int, 1000000000)
	return fib(value, &mem)
}

func fib(value int, mem *[]int) int {
	if (*mem)[value] != 0 {
		return (*mem)[value]
	} else if value < 1 {
		return 0
	} else if value == 1 {
		return 1
	} else {
		temp := fib(value-1, mem) + fib(value-2, mem)
		(*mem)[value] = temp
		return temp
	}
}
