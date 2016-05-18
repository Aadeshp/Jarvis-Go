package rest

import (
	"sync"
)

type Trie struct {
	mutex *sync.RWMutex
	root  *Node
}

func NewTrie() *Trie {
	return &Trie{
		root: newNode([]byte(" ")[0]),
	}
}

func (this *Trie) Insert(val []byte, route *Route) {
	//this.mutex.Lock()
	//defer this.mutex.Unlock()

	this.root.Insert(val, route)
}

func (this Trie) Find(httpMethod string, val []byte) *Route {
	//this.mutex.Lock()
	//defer this.mutex.Unlock()

	return this.root.Find(httpMethod, val)
}
