package goproject

import (
	"container/list"
	"github.com/LeilaBeken/goproject"
)

func Add(a, b int) int {
	return a + b
}

func addTwoNumbers(l1 *list.List, l2 *list.List) (l3 *list.List) {
	carry := 0
	l4 := list.New()
	e1 := l1.Front()
	e2 := l2.Front()
	for {
		sum := carry
		if e1 != nil {
			sum += e1.Value.(int)
			e1 = e1.Next()
		}
		if e2 != nil {
			sum += e2.Value.(int)
			e2 = e2.Next()
		}
		l4.PushBack(sum % 10)
		carry = sum / 10

		if e1 == nil && e2 == nil && carry == 0 {
			break
		}
	}
	return l4
}

func Delete(a, b int) float32 {
	return (float32(a)) / float32(b)
}
