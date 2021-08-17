package ch

import (
	"ftc/tt"
)

type Node struct {
	Next *Node
	Prev *Node
	Data interface{}
}

type Link struct {
	Count int
	Head *Node
	Last *Node
}

func NewLink()*Link {
	return &Link{

	}
}

func (obj *Link)Push(data interface{})  {
	node := &Node{
		Data:data,
	}

	if obj.Count==0 {
		obj.Head = node
		obj.Last = node
	}else {
		node.Prev = obj.Last
		obj.Last.Next = node
		obj.Last= node
	}

	obj.Count++
}
func (obj *Link)Pop()(data interface{}){
	if obj.Last== nil  {
		return data
	}
	data = obj.Last.Data
	if obj.Last==obj.Head {
		obj.Head = nil
		obj.Last = nil
	}else {
		obj.Last=obj.Last.Prev
		obj.Last.Next= nil
	}
	obj.Count--
	return data
}

func (obj *Link)Shift()(data interface{} ) {
	if obj.Head== nil {
		return data
	}
	data=obj.Head.Data
	if obj.Last == obj.Head {
		obj.Last = nil
		obj.Head = nil
	}else {
		obj.Head= obj.Head.Next
		obj.Head.Prev = nil
		obj.Count--
	}
	return data
}

func (obj *Link)Unshift(data interface{})  {
	node := &Node{
		Data: data,
	}
	if obj.Count == 0 {
		obj.Head = node
		obj.Last = node
	}else{
		node.Next=obj.Head
		obj.Head.Prev=node
		obj.Head= node
	}
	obj.Count++
}
func (obj *Link)Length ()int  {
	return obj.Count
}

func (obj *Link)Delete(num int)bool {
	if obj.Count < num-1 {
		return false
	}
	var curr *Node
	tt.GoFor(func(o *tt.Gfo) {
		if o.Num!=num{
			return
		}
		if obj.Last==curr {
			obj.Last = curr.Prev
			obj.Last.Next= nil
			return
		}
		if obj.Head==curr {
			obj.Head = curr.Next
			obj.Head.Prev= nil
			return
		}
		curr.Prev.Next=curr.Next
		curr.Next.Prev=curr.Prev
		obj.Count--

	}).Cond(func(o *tt.Gfo) bool {
		if curr == nil {
			curr= obj.Head
		}else{
			curr = curr.Next
		}
		return o.Num <= num
	}).Action()
	return true
}

func (obj *Link)Insert(num int,data interface{})  {
	if obj.Count < num-1 {
		return
	}
	node :=&Node{
		Data: data,
	}
	curr := obj.Head

	tt.GoFor(func(o *tt.Gfo) {
		if o.Num==num{
			node.Prev=curr
			node.Next= curr.Next
			curr.Next=node

			if obj.Last==curr {
				obj.Last = node
			}

			obj.Count++
		}
		curr = curr.Next

	}).Cond(func(o *tt.Gfo) bool {
		return o.Num <= num
	}).Action()
}
func (obj *Link)List()[]interface{} {
	list:=make([]interface{},0,obj.Count)
	curr := obj.Head
	list = append(list, curr.Data)

	tt.GoFor(func(o *tt.Gfo) {
			curr = curr.Next
			list = append(list, curr.Data)
	}).Cond(func(o *tt.Gfo) bool {
		return curr!=obj.Last
	}).Action()
	return list
}

