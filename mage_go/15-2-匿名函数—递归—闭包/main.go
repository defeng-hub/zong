package main

type foo func(a, b int)

func func1(f foo) {
	f(1, 2)
}

func main() {
	func2 := func() string {
		return "匿名函数执行"
	}

	func2()
}
