package problems

import (
	"fmt"
	"log"
	model "probsolbygolang/model"
)

func NewSLLNode(data int) *model.SLLNode{
	var node model.SLLNode
	node.Data = data
	return &node
}

func DisplayDataSLL(list *model.SLLNode){
	temp:=list
	for temp!=nil{
		fmt.Print(temp.Data)
		if temp.Next!=nil{
			fmt.Print("->")
		}
		temp=temp.Next
	}
	fmt.Println()
}

func AppendDataSLL(list *model.SLLNode,data int)*model.SLLNode{
	if list==nil{
		newNode:=NewSLLNode(data)	
		list=newNode
		return list
	}
	temp:=list
	for temp.Next!=nil{
		temp=temp.Next
	}
	newNode:=NewSLLNode(data)
	temp.Next=newNode
	return list
}

func InsertDataSLL(list *model.SLLNode,data int,position int)*model.SLLNode{
	if position<1{
		log.Println("Error: Invalid position")
		return list
	}
	if position==1{
		temp:=list
		newNode:=NewSLLNode(data)
		newNode.Next=temp
		return newNode
	}
	temp:=list
	counter:=1
	for temp!=nil&&counter<position-1{
		temp=temp.Next
		counter++
	}
	if temp==nil{
		log.Println("Error: given position is greater than size of list")
		return list
	}
	newNode:=NewSLLNode(data)
	newNode.Next=temp.Next
	temp.Next=newNode
	return list
}

func DeleteSpecificDataSLL(list *model.SLLNode,position int)*model.SLLNode{
	if position<1{
		log.Println("Error: Invalid position")
		return list
	}
	if position==1{
		temp:=list
		list=temp.Next
		return list
	}
	temp:=list
	counter:=1
	for temp!=nil&&counter<position-1{
		temp=temp.Next
		counter++
	}
	if temp==nil||temp.Next==nil{
		log.Println("Error: given position is greater than size of list")
		return list
	}
	delData:=temp.Next
	temp.Next=delData.Next
	return list
}

// by reversing the addresses
func ReverseTheSLL(list *model.SLLNode)*model.SLLNode{
	var prev,temp,next *model.SLLNode
	temp=list
	for temp!=nil{
		next=temp.Next
		temp.Next=prev
		prev=temp
		temp=next
	}
	list=prev
	return list
}

// using stack
func ReverseSLLByStack(list *model.SLLNode)*model.SLLNode{
	var st []int
	temp:=list
	for temp!=nil{
		st=append(st,temp.Data)
		temp=temp.Next
	}
	temp=list
	for i:=len(st)-1;i>=0;i--{
		temp.Data=st[i]
		temp=temp.Next
	}
	return list
}