package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	// "regexp"
	// "strconv"
	// "strings"
	// "math"
)

// Антон учит Степана делить массивы одинаковой длины n друг на друга следующим образом:
// если ai нацело делится на bi для всех 1≤i≤n, то массив a делится на массив b.
// В качестве упражнения, Степан всегда использует массив [1,2,…,n] в качестве делителя.
// Теперь он хочет посчитать количество возможных массивов a, которые делятся на заданный делитель.
// Но Антон догадался, что таких массивов бесконечно много, поэтому для Степана
// он наложил на построение массива a следующие ограничение:
// li ≤ ai ≤ ri для всех 1≤i≤n, где l и r — подготовленные заранее Антоном массивы.
// Но после введения таких ограничений, посчитать ответ оказалось слишком трудно, поэтому ребята просят вас им помочь.

// 3
// 3
// 1 1 1
// 2 5 5
// 5
// 1 2 7 10 20
// 1 4 10 30 40
// 4
// 1000 1000 1000 1000
// 2000 2000 2000 2000

// 4
// 50
// 916957796

func main() {
	// fmt.Print("Start")
	in, out := getBuffers()
	defer out.Flush()

	var nGroup int
	var res int64
	var nStr int64

	fmt.Fscan(in, &nGroup)
	for range nGroup {

		fmt.Fscan(in, &nStr)
		l := make([]int64, nStr)
		r := make([]int64, nStr)

		res = 1

		var i int64
		for i = 0; i < nStr; i++ {
			fmt.Fscan(in, &l[i])
		}
		fmt.Println(l)

		for i = 0; i < nStr; i++ {
			fmt.Fscan(in, &r[i])
		}
		fmt.Println(r)

		for i = 0; i < nStr; i++ {
			n := i + 1
			a := l[i] / n
			b := r[i] / n

			fmt.Println("Номер прохода:", i, "; a:", a, "; b:", b, "; n:", n, "; l[i]:", l[i], "; r[i]:", r[i])

			if (a*n == l[i]) && (b*n == r[i]) && (l[i] != r[i]) {
				fmt.Println("Прибавил к res +1")
				res *= (b - a) + 1
				// print(10)
			} else if l[i] == r[i] {
				continue
			} else {
				res *= (b - a)
			}
		}
	}
	print("Result:", res%1000000007)
}

func getBuffers() (*bufio.Reader, *bufio.Writer) {
	var input io.Reader = os.Stdin

	for i, arg := range os.Args {
		if arg == "-f" && i+1 < len(os.Args) {
			f, err := os.Open(os.Args[i+1])
			if err == nil {
				input = f
			}
		}
	}
	return bufio.NewReader(input), bufio.NewWriter(os.Stdout)
}
