package compilersteps

import (
	"fmt"
)

// Visitor is an interface for visiting nodes. It must define a visit function per node.
// dfs tree traversal is done in the visitor pattern
type Visitor interface {
	VisitSIf(*SIfNode)
	VisitSBegin(*SBeginNode)
	VisitSPrint(*SPrintNode)
	VisitLEnd(*LEndNode)
	VisitLSemicolon(*LSemicolonNode)
	VisitEGeneral(*EGeneralNode)
}

type SemanticAnalyzer struct {
	NodeCount int
	Errors    []error
}

func (v *SemanticAnalyzer) VisitSIf(n *SIfNode) {
	v.NodeCount += 1
	n.E.Accept(v)
	n.S1.Accept(v)
	n.S2.Accept(v)
	fmt.Println(NodeString(n))
	// TODO: Semantic Analysis
}

func (v *SemanticAnalyzer) VisitSBegin(n *SBeginNode) {
	v.NodeCount += 1
	n.S.Accept(v)
	n.L.Accept(v)
	fmt.Println(NodeString(n))
	// TODO: Semantic Analysis
}

func (v *SemanticAnalyzer) VisitSPrint(n *SPrintNode) {
	v.NodeCount += 1
	n.E.Accept(v)
	fmt.Println(NodeString(n))
	// TODO: Semantic Analysis
}

func (v *SemanticAnalyzer) VisitLEnd(n *LEndNode) {
	v.NodeCount += 1
	fmt.Println(NodeString(n))
	// TODO: Semantic Analysis (not much to do here at the end of the program)
}

func (v *SemanticAnalyzer) VisitLSemicolon(n *LSemicolonNode) {
	v.NodeCount += 1
	n.S.Accept(v)
	n.L.Accept(v)
	fmt.Println(NodeString(n))
	// TODO: Semantic Analysis
}

func (v *SemanticAnalyzer) VisitEGeneral(n *EGeneralNode) {
	v.NodeCount += 1
	fmt.Println(NodeString(n))
	// TODO: Semantic Analysis
}
