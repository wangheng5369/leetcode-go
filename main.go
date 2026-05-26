package leetcodego

func A() {
	a := make([][]int, 200000)
	b := make([][]int, 200000)
	for i := range a {
		a[i] = make([]int, 200000)
		b[i] = make([]int, 200000)
	}
	for i := 0; i < 200000; i++ {
		for j := 0; j < 200000; j++ {
			a[i][j] = a[i][j] + b[i][j]
		}
	}
}

func B() {
	a := make([][]int, 200000)
	b := make([][]int, 200000)
	for i := range a {
		a[i] = make([]int, 200000)
		b[i] = make([]int, 200000)
	}
	for i := 0; i < 200000; i++ {
		for j := 0; j < 200000; j++ {
			a[i][j] = a[i][j] + b[i][j]
		}
	}
}
