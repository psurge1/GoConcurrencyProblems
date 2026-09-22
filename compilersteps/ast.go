package compilersteps

import (
	"fmt"
)

type (
	Node interface {
		node()
		Accept(Visitor)
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
	return fmt.Sprintf("%#v", n)
}

type SIfNode struct {
	E  ENode
	S1 SNode
	S2 SNode
}

type SBeginNode struct {
	S SNode
	L LNode
}

type SPrintNode struct {
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

func (*SIfNode) node()        {}
func (*SBeginNode) node()     {}
func (*SPrintNode) node()     {}
func (*LEndNode) node()       {}
func (*LSemicolonNode) node() {}
func (*EGeneralNode) node()   {}

func (*SIfNode) sNode()    {}
func (*SBeginNode) sNode() {}
func (*SPrintNode) sNode() {}

func (*LEndNode) lNode()       {}
func (*LSemicolonNode) lNode() {}

func (*EGeneralNode) eNode() {}

func (n *SIfNode) Accept(v Visitor) {
	v.VisitSIf(n)
	n.E.Accept(v)
	n.S1.Accept(v)
}

func (n *SBeginNode) Accept(v Visitor) {
	v.VisitSBegin(n)
	n.S.Accept(v)
	n.L.Accept(v)
}

func (n *SPrintNode) Accept(v Visitor) {
	v.VisitSPrint(n)
	n.E.Accept(v)
}

func (n *LEndNode) Accept(v Visitor) {
	v.VisitLEnd(n)
}

func (n *LSemicolonNode) Accept(v Visitor) {
	v.VisitLSemicolon(n)
	n.S.Accept(v)
	n.L.Accept(v)
}

func (n *EGeneralNode) Accept(v Visitor) {
	v.VisitEGeneral(n)
}
