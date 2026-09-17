package compilersteps

import (
	"fmt"
)

type (
	Node  interface{}
	SNode interface{}
	LNode interface{}
)

type IfNode struct {
	e  ENode
	s1 SNode
	s2 SNode
}

type BeginNode struct {
	s SNode
	l LNode
}

type PrintNode struct {
	e ENode
}

type LEndNode struct{}

type LSemicolonNode struct {
	s SNode
	l LNode
}

type ENode struct {
	left  int
	right int
}

/*
* S -> if E then S else S
* S -> begin S L
* S - > print E
* L -> end
* L -> ; S L
* E -> num = num
* */

// RD is an implementation of the LL(1) CFG above
type RD struct{}

func NewRD() RD {
	return RD{}
}

func (rd *RD) Error(tk Token) error {
	return fmt.Errorf("syntax error: %s", tk.String())
}

func (rd *RD) Next() Token {
	return Token{}
}

func (rd *RD) S() (*SNode, error) {
	tk := rd.Next()
	switch tk.Kind {
	case IF:
		_, err := rd.E()
		if err != nil {
			return nil, err
		}
		nextTk := rd.Next()
		if nextTk.Kind != THEN {
			return nil, rd.Error(nextTk)
		}
		_, err = rd.S()
		if err != nil {
			return nil, err
		}
		nextTk = rd.Next()
		if nextTk.Kind != ELSE {
			return nil, rd.Error(nextTk)
		}
		_, err = rd.S()
		if err != nil {
			return nil, err
		}
		return nil, nil
	case BEGIN:
		_, err := rd.S()
		if err != nil {
			return nil, err
		}
		_, err = rd.L()
		if err != nil {
			return nil, err
		}
		return nil, nil
	case PRINT:
		_, err := rd.E()
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	return nil, rd.Error(tk)
}

func (rd *RD) L() (*LNode, error) {
	tk := rd.Next()
	if tk.Kind == END {
		// todo
		return nil, nil
	} else if tk.Kind == SEMICOLON {
		_, err := rd.S()
		if err != nil {
			return nil, err
		}
		_, err = rd.L()
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	return nil, rd.Error(tk)
}

func (rd *RD) E() (*ENode, error) {
	tk := rd.Next()
	if tk.Kind == NUM {
		return nil, nil
	}
	tk = rd.Next()
	if tk.Kind == EQ {
		return nil, nil
	}
	tk = rd.Next()
	if tk.Kind == NUM {
		return nil, nil
	}
	return nil, rd.Error(tk)
}

func TestParser() {
	rd := NewRD()
	rd.S()
}
