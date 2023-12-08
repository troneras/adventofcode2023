package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Name string
	// other properties...
}

type Graph struct {
	Nodes map[*Node][]*Node
}

func (g *Graph) String() string {
	var b strings.Builder
	for n, cons := range g.Nodes {
		b.WriteString(n.Name)
		b.WriteString(" -> ")
		for _, con := range cons {
			b.WriteString(con.Name)
			b.WriteString(" ")
		}
		b.WriteString("\n")
	}
	return b.String()
}

func (g *Graph) CreateNodeIfNotExist(name string) *Node {
	for n := range g.Nodes {
		if n.Name == name {
			return n
		}
	}
	node := &Node{Name: name}
	g.Nodes[node] = nil
	return node
}

func (g *Graph) AddNode(n *Node, connections []*Node) {
	if g.Nodes == nil {
		g.Nodes = make(map[*Node][]*Node)
	}
	g.Nodes[n] = connections
}

func (g *Graph) GetNode(name string) *Node {
	for n := range g.Nodes {
		if n.Name == name {
			return n
		}
	}
	return nil
}

func parseNodeConnections(b []byte) (string, string) {
	s := string(b)
	s = strings.Trim(s, "()")
	parts := strings.Split(s, ", ")
	return parts[0], parts[1]
}

func main() {
	body, _ := os.ReadFile("input.txt")

	instructions, block_nodes, _ := bytes.Cut(body, []byte("\n\n"))
	fmt.Println(string(instructions))

	nodes := bytes.Split(block_nodes, []byte("\n"))

	graph := Graph{}

	graph.Nodes = make(map[*Node][]*Node)
	for _, node := range nodes {
		// fmt.Println(string(node))
		n, cons, _ := bytes.Cut(node, []byte(" = "))
		part1, part2 := parseNodeConnections(cons)

		con1 := graph.CreateNodeIfNotExist(part1)
		con2 := graph.CreateNodeIfNotExist(part2)
		root := graph.CreateNodeIfNotExist(string(n))

		graph.AddNode(root, []*Node{con1, con2})
	}

	fmt.Println(graph.String())

	// endNode := ""
	// next_node := graph.GetNode("AAA")
	// total := 0

	// for endNode != "ZZZ" {
	// 	for _, inst := range instructions {
	// 		if inst == byte('L') {
	// 			next_node = graph.Nodes[next_node][0]
	// 		} else if inst == byte('R') {
	// 			next_node = graph.Nodes[next_node][1]
	// 		}
	// 	}
	// 	endNode = next_node.Name
	// 	total += len(instructions)
	// }

	// fmt.Println(total)

	// part 2
	next_nodes := []*Node{} // starting nodes: all nodes ending in A
	for n := range graph.Nodes {
		if strings.HasSuffix(n.Name, "A") {
			fmt.Println("start ", n)
			next_nodes = append(next_nodes, n)
		}
	}

	total := 0
	left_right := 0
	found := false
	totals := []int{}
	for _, next := range next_nodes {
		for !found {
			for _, inst := range instructions {
				if inst == byte('L') {
					left_right = 0
				} else if inst == byte('R') {
					left_right = 1
				}
				tmp_node := graph.Nodes[next][left_right]
				// if ends in Z, add to end_nodes
				// fmt.Println("lr", left_right, "cur", next, "next", tmp_node)
				next = tmp_node

				if strings.HasSuffix(tmp_node.Name, "Z") {
					found = true
				}
			}

			total += len(instructions)
		}
		found = false
		fmt.Println("total ", total)
		totals = append(totals, total)
		total = 0
	}

	// calculate LCM of totals
	fmt.Println("lcm", lcmOfTotals(totals))

}
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func lcmOfTotals(totals []int) int {
	result := totals[0]
	for i := 1; i < len(totals); i++ {
		result = lcm(result, totals[i])
	}
	return result
}
