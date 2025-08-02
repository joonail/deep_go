package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type node struct {
	key, value  int
	left, right *node
}

type OrderedMap struct {
	root *node
	size int
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{}
}

func (m *OrderedMap) Insert(key, value int) {
	m.root = insertNode(m.root, key, value, &m.size)
}

func insertNode(n *node, key, value int, size *int) *node {
	if n == nil {
		*size++
		return &node{key: key, value: value}
	}
	if key < n.key {
		n.left = insertNode(n.left, key, value, size)
	} else if key > n.key {
		n.right = insertNode(n.right, key, value, size)
	} else {
		n.value = value
	}
	return n
}

func (m *OrderedMap) Erase(key int) {
	var deleted bool
	m.root, deleted = eraseNode(m.root, key)
	if deleted {
		m.size--
	}
}

func eraseNode(n *node, key int) (*node, bool) {
	if n == nil {
		return nil, false
	}
	if key < n.key {
		var deleted bool
		n.left, deleted = eraseNode(n.left, key)
		return n, deleted
	}
	if key > n.key {
		var deleted bool
		n.right, deleted = eraseNode(n.right, key)
		return n, deleted
	}

	// Node found
	if n.left == nil {
		return n.right, true
	}
	if n.right == nil {
		return n.left, true
	}

	successor := minNode(n.right)
	n.key, n.value = successor.key, successor.value
	n.right, _ = eraseNode(n.right, successor.key)
	return n, true
}

func minNode(n *node) *node {
	for n.left != nil {
		n = n.left
	}
	return n
}

func (m *OrderedMap) Contains(key int) bool {
	return containsNode(m.root, key)
}

func containsNode(n *node, key int) bool {
	if n == nil {
		return false
	}
	if key < n.key {
		return containsNode(n.left, key)
	} else if key > n.key {
		return containsNode(n.right, key)
	}
	return true
}

func (m *OrderedMap) Size() int {
	return m.size
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	inOrderTraversal(m.root, action)
}

func inOrderTraversal(n *node, action func(int, int)) {
	if n == nil {
		return
	}
	inOrderTraversal(n.left, action)
	action(n.key, n.value)
	inOrderTraversal(n.right, action)
}

func TestCircularQueue(t *testing.T) {
	data := NewOrderedMap()
	assert.Zero(t, data.Size())

	data.Insert(10, 10)
	data.Insert(5, 5)
	data.Insert(15, 15)
	data.Insert(2, 2)
	data.Insert(4, 4)
	data.Insert(12, 12)
	data.Insert(14, 14)

	assert.Equal(t, 7, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(3))
	assert.False(t, data.Contains(13))

	var keys []int
	expectedKeys := []int{2, 4, 5, 10, 12, 14, 15}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))

	data.Erase(15)
	data.Erase(14)
	data.Erase(2)

	assert.Equal(t, 4, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(2))
	assert.False(t, data.Contains(14))

	keys = nil
	expectedKeys = []int{4, 5, 10, 12}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))
}
