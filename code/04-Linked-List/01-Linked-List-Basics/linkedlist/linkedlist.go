/*
 * @Author: duanhaobin
 * @Date: 2021-03-21 10:56:12
 * @LastEditTime: 2021-03-21 11:03:32
 * @FilePath: \DSA_study\code\04-Linked-List\01-Linked-List-Basics\linkedlist\linkedlist.go
 */
package linkedlist

import "fmt"

// Node 基础节点结构定体
type Node struct {
	e    interface{}
	next *Node
}

func (n *Node) String() string {
	return fmt.Sprint(n.e)
}
