package compilersteps

import (
	"fmt"
)

type Visitor struct{}

type (
	Node interface {
		node()
	}

	SNode interface {
		Node
		sNode()
	}
	LNode interface {
		Node
		lNode()
	}
	ENode interface {
		Node
		eNode()
	}
)

func NodeString(n Node) string {
	return fmt.Sprintf("%+v", n)
}

type IfNode struct {
	E  ENode
	S1 SNode
	S2 SNode
}

type BeginNode struct {
	S SNode
	L LNode
}

type PrintNode struct {
	E ENode
}

type LEndNode struct{}

type LSemicolonNode struct {
	S SNode
	L LNode
}

type EGeneralNode struct {
	Left  Value
	Right Value
}

func (*IfNode) node()         {}
func (*BeginNode) node()      {}
func (*PrintNode) node()      {}
func (*LEndNode) node()       {}
func (*LSemicolonNode) node() {}
func (*EGeneralNode) node()   {}

func (*IfNode) sNode()    {}
func (*BeginNode) sNode() {}
func (*PrintNode) sNode() {}

func (*LEndNode) lNode()       {}
func (*LSemicolonNode) lNode() {}

func (*EGeneralNode) eNode() {}
