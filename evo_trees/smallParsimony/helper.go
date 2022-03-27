package main

import "fmt"

func PrintNodes(T []*Node) {
	for _, node := range T {
		if node != nil {
			fmt.Print(node.id, " ", node.seq, " ")
			if node.child1 != nil {
				fmt.Print("node: ", node.id, " child1 id: ", node.child1.id, " child2 id: ", node.child2.id)
			}
			if node.parent != nil {
				fmt.Print(" parent: ", node.parent.id, " len to parent: ", node.lenToParent)
			}
			fmt.Println(" ")
		}
	}
}

func PrintLeavesS(n int, s [][]int) {
	for i := 0; i < n; i++ {
		fmt.Println(s[i])
	}
}

func PrintInternalNodeS(n int, s [][]int) {
	for i := n; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

func PrintAdj(T []*Node) {
	for _, node := range T {
		if node.parent != nil {
			fmt.Printf("%s->%s:%d\n", node.seq, node.parent.seq, node.lenToParent)
			fmt.Printf("%s->%s:%d\n", node.parent.seq, node.seq, node.lenToParent)
		}
	}
}
