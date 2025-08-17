package main

import (
	"fmt"
	"math"
)

type ParenthesesStack struct {
	stack  []byte
	result []string
}

func New(n int) *ParenthesesStack {
	return &ParenthesesStack{
		stack:  make([]byte, 0, 2*n),
		result: make([]string, 0, int(math.Pow(4.0, float64(n)))),
	}
}

func (ps *ParenthesesStack) Push(p byte) {
	ps.stack = append(ps.stack, p)
}

func (ps *ParenthesesStack) Pop() {
	ps.stack = ps.stack[:len(ps.stack)-1]
}

func (ps *ParenthesesStack) AppendResult() {
	ps.result = append(ps.result, string(ps.stack))
}

func (ps *ParenthesesStack) Backtrack(opened, closed int) {
	if opened+closed == cap(ps.stack) {
		ps.AppendResult()
		return
	}
	if opened < cap(ps.stack)/2 {
		ps.Push('(')
		ps.Backtrack(opened+1, closed)
		ps.Pop()
	}
	if closed < opened {
		ps.Push(')')
		ps.Backtrack(opened, closed+1)
		ps.Pop()
	}
}

func (ps ParenthesesStack) Result() []string {
	return ps.result
}

func generateParenthesis(n int) []string {
	ps := New(n)
	ps.Backtrack(0, 0)
	return ps.Result()
}

func main() {
	fmt.Println(generateParenthesis(1)) // ()
	fmt.Println(generateParenthesis(2)) // (()) ()()
}
