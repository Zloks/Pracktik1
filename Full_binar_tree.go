package main

import "fmt"

// Структура Node представляет узел бинарного дерева.
type Node struct {
    Value int
    Left  *Node
    Right *Node
}

// Структура BinaryTree представляет само бинарное дерево.
type BinaryTree struct {
    Root *Node
}

// Метод Add добавляет новый элемент в бинарное дерево.
func (bt *BinaryTree) Add(value int) {
    if bt.Root == nil {
        bt.Root = &Node{Value: value}
    } else {
        bt.Root.add(value)
    }
}

// Метод add добавляет новый элемент в поддерево с корнем в данном узле.
func (n *Node) add(value int) {
    if value < n.Value {
        if n.Left == nil {
            n.Left = &Node{Value: value}
        } else {
            n.Left.add(value)
        }
    } else {
        if n.Right == nil {
            n.Right = &Node{Value: value}
        } else {
            n.Right.add(value)
        }
    }
}

// Метод Remove удаляет элемент из бинарного дерева.
func (bt *BinaryTree) Remove(value int) {
    bt.Root = bt.Root.remove(value)
}

// Метод remove удаляет элемент из поддерева с корнем в данном узле.
func (n *Node) remove(value int) *Node {
    if n == nil {
        return nil
    }
    if value < n.Value {
        n.Left = n.Left.remove(value)
        return n
    }
    if value > n.Value {
        n.Right = n.Right.remove(value)
        return n
    }
    if n.Left == nil && n.Right == nil {
        return nil
    }
    if n.Left == nil {
        return n.Right
    }
    if n.Right == nil {
        return n.Left
    }
    smallestRight := n.Right
    for smallestRight.Left != nil {
        smallestRight = smallestRight.Left
    }
    n.Value = smallestRight.Value
    n.Right = n.Right.remove(smallestRight.Value)
    return n
}

// Функция printInOrder выводит элементы дерева в порядке возрастания.
func printInOrder(node *Node) {
    if node != nil {
        printInOrder(node.Left)
        fmt.Println(node.Value)
        printInOrder(node.Right)
    }
}