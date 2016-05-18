package rest

import (
//"fmt"
)

type Node struct {
	children map[byte]*Node
	value    byte
	isLeaf   bool
	routes   []*Route
}

func newNode(val byte) *Node {
	return &Node{
		value:    val,
		children: make(map[byte]*Node),
		routes:   []*Route{},
		isLeaf:   false,
	}
}

func (this *Node) Insert(val []byte, route *Route) {
	children := this.children

	for i := 0; i < len(val); i += 1 {
		c := val[i]

		var n *Node
		if children[c] != nil {
			n = children[c]
		} else {
			n = newNode(c)
			children[c] = n
		}

		children = n.children

		if i == len(val)-1 {
			n.isLeaf = true
			n.routes = append(n.routes, route)
		}
	}
}

func (this Node) Find(httpMethod string, val []byte) *Route {
	children := this.children

	var n *Node
	for i := 0; i < len(val); i += 1 {
		c := val[i]
		if children[c] != nil {
			n = children[c]
			children = n.children
		} else {
			return nil
		}
	}

	if n.isLeaf {
		for _, route := range n.routes {
			if route.HttpMethod == httpMethod {
				return route
			}
		}
	}

	return nil
}
