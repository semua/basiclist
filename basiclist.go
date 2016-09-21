package basiclist

import (
	"errors"
	"fmt"
)

type CompareFunc func(val0 interface{}, val1 interface{}) int
type BasicNode struct {
	val      interface{}
	nextNode *BasicNode
	prevNode *BasicNode
}

type BasicList struct {
	head    *BasicNode
	tail    *BasicNode
	mapNode map[interface{}]*BasicNode
	size    int
}

//NewBasicList : Init structure for basic Sorted Linked List.
func NewBasicList() *BasicList {
	initMap := make(map[interface{}]*BasicNode)
	return &BasicList{head: nil, tail: nil, size: 0, mapNode: initMap}
}

/*
 * 有序插入
 */
func (b *BasicList) Insert(compare CompareFunc, key interface{}, value interface{}) {
	if b.head == nil {
		b.head = &BasicNode{val: value, nextNode: nil, prevNode: nil}
		b.tail = b.head
		b.mapNode[key] = b.head
	} else {
		var currentNode *BasicNode
		currentNode = b.head
		var prevNode *BasicNode
		prevNode = b.head.prevNode
		var found bool
		newNode := &BasicNode{val: value, nextNode: nil, prevNode: nil}
		b.mapNode[key] = newNode
		for {
			if compare(currentNode.val, value) > 0 {
				if prevNode == nil {
					newNode.nextNode = currentNode
					b.head = newNode
				} else {
					newNode.nextNode = prevNode.nextNode
					newNode.prevNode = prevNode
					prevNode.nextNode = newNode
				}
				currentNode.prevNode = newNode
				found = true
				break
			}

			if currentNode.nextNode == nil {
				break
			}

			prevNode = currentNode
			currentNode = currentNode.nextNode
		}

		if found == false {
			currentNode.nextNode = newNode
			newNode.prevNode = currentNode
			b.tail = newNode
		}
	}
	b.size = b.size + 1
}

func (b *BasicList) GetByKey(key interface{}) (interface{}, error) {
	if n, ok := b.mapNode[key]; ok {
		return n.val, nil
	}
	return nil, errors.New("not found")
}

func (b *BasicList) Equal(compare CompareFunc, value interface{}) (interface{}, error) {
	currentNode := b.head
	tailNode := b.tail
	for {
		if currentNode == nil || tailNode == nil {
			break
		}
		if compare(currentNode.val, value) == 0 {
			return currentNode, nil
		}
		if compare(tailNode.val, value) == 0 {
			return tailNode, nil
		}
		if tailNode == currentNode || currentNode.nextNode == tailNode {
			break
		}
		currentNode = currentNode.nextNode
		tailNode = tailNode.prevNode
	}
	return nil, errors.New("not found")
}

func (b *BasicList) LessThanOrEqual(compare CompareFunc, value interface{}) (interface{}, error) {
	currentNode := b.head
	tailNode := b.tail
	if currentNode == nil {
		return nil, errors.New("not found.")
	}
	for {
		if compare(currentNode.val, value) > 0 {
			break
		}
		if compare(currentNode.val, value) <= 0 && compare(currentNode.nextNode.val, value) > 0 {
			return currentNode, nil
		}
		if compare(tailNode.val, value) <= 0 {
			return tailNode, nil
		}
		currentNode = currentNode.nextNode
		tailNode = tailNode.prevNode
	}
	return nil, errors.New("not found")
}

func (b *BasicList) GreaterThanOrEqual(compare CompareFunc, value interface{}) (interface{}, error) {
	currentNode := b.head
	tailNode := b.tail
	if currentNode == nil {
		return nil, errors.New("not found.")
	}
	for {
		if compare(currentNode.val, value) >= 0 {
			return currentNode, nil
		}
		if compare(tailNode.val, value) < 0 {
			break
		}
		if compare(tailNode.val, value) >= 0 && tailNode.prevNode != nil && compare(tailNode.prevNode.val, value) < 0 {
			return tailNode, nil
		}
		currentNode = currentNode.nextNode
		tailNode = tailNode.prevNode
	}
	return nil, errors.New("not found")
}
func (b *BasicList) Remove(key interface{}) (interface{}, error) {
	if n, ok := b.mapNode[key]; ok {
		prevNode := n.prevNode
		nextNode := n.nextNode
		if prevNode != nil {
			prevNode.nextNode = nextNode
		} else {
			b.head = n.nextNode
		}
		if nextNode != nil {
			nextNode.prevNode = prevNode
		} else {
			b.tail = prevNode
		}
		delete(b.mapNode, key)
		b.size = b.size - 1
		return n.val, nil
	}
	return nil, errors.New("not found")
}

func (b *BasicList) List(start int, limit int) []interface{} {
	returnList := []interface{}{}
	if start > b.size || limit <= 0 {
		return returnList
	}
	currentNode := b.head
	idx := 0
	for {
		if idx >= start+limit {
			break
		}
		if currentNode == nil {
			break
		}
		if idx >= start {
			returnList = append(returnList, currentNode.val)
		}
		currentNode = currentNode.nextNode
		idx = idx + 1
	}
	return returnList
}
func (b *BasicList) Print() {
	fmt.Println("")
	fmt.Printf("head->")
	if b.size == 0 {
		fmt.Printf("nil\n")
		return
	}
	currentNode := b.head
	for {
		fmt.Printf("[val:%v]->", currentNode.val)
		if currentNode.nextNode == nil {
			break
		}
		currentNode = currentNode.nextNode
	}
	fmt.Printf("nil\n")
}
