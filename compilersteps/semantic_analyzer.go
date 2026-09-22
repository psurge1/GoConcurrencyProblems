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
	Errors []error
}

func (v *SemanticAnalyzer) VisitSIf(n *SIfNode) {
	fmt.Println(NodeString(n))
	// TODO: Semantic Analysis
}

func (v *SemanticAnalyzer) VisitSBegin(n *SBeginNode) {
	fmt.Println(NodeString(n))
	// TODO: Semantic Analysis
}

func (v *SemanticAnalyzer) VisitSPrint(n *SPrintNode) {
	fmt.Println(NodeString(n))
	// TODO: Semantic Analysis
}

func (v *SemanticAnalyzer) VisitLEnd(n *LEndNode) {
	fmt.Println(NodeString(n))
	// TODO: Semantic Analysis (not much to do here at the end of the program)
}

func (v *SemanticAnalyzer) VisitLSemicolon(n *LSemicolonNode) {
	fmt.Println(NodeString(n))
	// TODO: Semantic Analysis
}

func (v *SemanticAnalyzer) VisitEGeneral(n *EGeneralNode) {
	fmt.Println(NodeString(n))
	// TODO: Semantic Analysis
}
