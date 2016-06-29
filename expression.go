package main

import (
	"fmt"
)

const (
	OPERATOR = "operator"
	NUMBER   = "number"
)

type Node struct {
	nodeType string
	value    int
	op       byte
}

var (
	priority = map[byte]int{
		'(': 0,
		')': 0,
		'-': 1,
		'+': 1,
		'*': 2,
		'/': 2,
	}
)

type Stack []Node

func (s *Stack) push(x Node) {
	*s = append(*s, x)
}

func (s *Stack) pop() Node {
	x := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return x
}

func (s Stack) top() Node {
	return s[len(s)-1]
}

func (s Stack) length() int {
	return len(s)
}

func isNumber(c byte) bool {
	return c >= '0' && c <= '9'
}

func isOperator(c byte) bool {
	return !isNumber(c)
}

func nextChar(exp string, idx int) int {
	for idx < len(exp) && exp[idx] == ' ' {
		idx += 1
	}
	return idx
}

func nextNode(exp string, idx int) (node Node, nextIdx int) {
	idx = nextChar(exp, idx)
	if isOperator(exp[idx]) {
		return Node{
			nodeType: OPERATOR,
			op:       exp[idx],
		}, idx + 1
	} else {
		x := 0
		idx = nextChar(exp, idx)
		for idx < len(exp) && isNumber(exp[idx]) {
			x = x*10 + int(exp[idx]) - 48
			idx = nextChar(exp, idx+1)
		}
		return Node{
			nodeType: NUMBER,
			value:    x,
		}, idx
	}
}

func isPriorityHigher(a Node, b Node) bool {
	return priority[a.op] > priority[b.op]
}

func cal(a Node, op Node, b Node) int {
	switch op.op {
	case '+':
		return a.value + b.value
	case '-':
		return a.value - b.value
	case '*':
		return a.value * b.value
	case '/':
		return a.value / b.value
	}
	panic("unknow op" + string(op.op))
}

func solve(exp string) {
	i := 0
	var ns, ops Stack
	var previousNode Node
	for i < len(exp) {
		var node Node
		node, i = nextNode(exp, i)
		switch node.nodeType {
		case NUMBER:
			ns.push(node)
		case OPERATOR:
			if ops.length() == 0 || node.op == '(' {
				// node can only be '+', '-'
				ops.push(node)
			} else if isPriorityHigher(node, ops.top()) || (previousNode.op == '(' && node.op != ')') {
				if previousNode.op == '(' {
					ns.push(Node{
						nodeType: NUMBER,
						value:    0,
					})
				}
				ops.push(node)
			} else {
				for ops.length() > 0 && ops.top().op != '(' && !isPriorityHigher(node, ops.top()) {
					b := ns.pop()
					op := ops.pop()
					var a Node
					if ns.length() == 0 {
						a = Node{
							nodeType: NUMBER,
							value:    0,
						}
					} else {
						a = ns.pop()
					}
					newNode := Node{
						nodeType: NUMBER,
						value:    cal(a, op, b),
					}
					ns.push(newNode)
					// cal
				}
				if node.op == ')' {
					ops.pop()
				} else {
					ops.push(node)
				}
			}
			break
		}
		previousNode = node
	}
	for ops.length() > 0 {
		var v int
		if ns.length() == 1 {
			v = cal(Node{nodeType: NUMBER, value: 0}, ops.pop(), ns.pop())
		} else {
			b := ns.pop()
			op := ops.pop()
			a := ns.pop()

			v = cal(a, op, b)
		}
		ns.push(Node{
			nodeType: NUMBER,
			value:    v,
		})
	}
	fmt.Printf("%s = ", exp)
	fmt.Println(ns.pop().value)
}

func main() {
	solve("232442")
	solve("23  24 42 + 1 1")
	solve("23  24 42 + (1 1)")
	solve("1")
	solve("-1")
	solve("(-1)")
	solve("1-1")
	solve("1-1*2")
	solve("1-1*2+9")
	solve("1-1*(-2)+9")
	solve("1-1*(2+9)")
	solve("1-3*(2+9)")
	solve("1-3*(2+ 9)")
	solve("-1-3*(2+9)")
	solve("-1-3*(2+9)-1")
	solve("7*(-1-3*(2+9)-1)")
}
