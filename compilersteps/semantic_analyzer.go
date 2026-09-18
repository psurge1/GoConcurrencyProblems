package compilersteps

import (
	"fmt"
)

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
	Errors []error
}

func (v *SemanticAnalyzer) VisitSIf(n *SIfNode) {
	n.E.Accept(v)
	n.S1.Accept(v)
	n.S2.Accept(v)

	fmt.Println(NodeString(n))
	// TODO: Semantic Analysis
}

func (v *SemanticAnalyzer) VisitSBegin(n *SBeginNode) {
	n.S.Accept(v)
	n.L.Accept(v)

	fmt.Println(NodeString(n))
	// TODO: Semantic Analysis
}

func (v *SemanticAnalyzer) VisitSPrint(n *SPrintNode) {
	n.E.Accept(v)

	fmt.Println(NodeString(n))
	// TODO: Semantic Analysis
}

func (v *SemanticAnalyzer) VisitLEnd(n *LEndNode) {
	fmt.Println(NodeString(n))
	// TODO: Semantic Analysis (not much to do here at the end of the program)
}

func (v *SemanticAnalyzer) VisitLSemicolon(n *LSemicolonNode) {
	n.S.Accept(v)
	n.L.Accept(v)

	fmt.Println(NodeString(n))
	// TODO: Semantic Analysis
}

func (v *SemanticAnalyzer) VisitEGeneral(n *EGeneralNode) {
	fmt.Println(NodeString(n))
	// TODO: Semantic Analysis
}
