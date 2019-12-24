package main

import (
	"fmt"
)

const (
	Add      int64 = 1
	Multiply int64 = 2
	End      int64 = 99
)

var result []int64

type Program struct {
	OP     int64
	A      int
	B      int
	Result int
}

func main() {
	in, err := scanFile("input.txt")
	checkFatal(err)
	result = in

	pgs := make([]Program, 0)

	for i := 0; i < len(in); i = i + 4 {
		if in[i] == Add || in[i] == Multiply {
			p := Program{
				OP:     in[i],
				A:      i + 1,
				B:      i + 2,
				Result: i + 3,
			}
			pgs = append(pgs, p)
		} else if in[i] == End {
			break
		}
	}

	for _, v := range pgs {
		fmt.Println(v)
	}
}

func Calculate(res []int64, programs []Program) []int64 {
	return nil
}
