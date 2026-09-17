package compilersteps

import (
	"fmt"
)

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
	return fmt.Errorf("syntax error on token: %s", tk.String())
}

func (rd *RD) Next() Token {
	return Token{}
}

func (rd *RD) S() (SNode, error) {
	tk := rd.Next()
	switch tk.Kind {
	case IF:
		e, err := rd.E()
		if err != nil {
			return nil, err
		}
		nextTk := rd.Next()
		if nextTk.Kind != THEN {
			return nil, rd.Error(nextTk)
		}
		sOne, err := rd.S()
		if err != nil {
			return nil, err
		}
		nextTk = rd.Next()
		if nextTk.Kind != ELSE {
			return nil, rd.Error(nextTk)
		}
		sTwo, err := rd.S()
		if err != nil {
			return nil, err
		}
		return &IfNode{e, sOne, sTwo}, nil
	case BEGIN:
		s, err := rd.S()
		if err != nil {
			return nil, err
		}
		l, err := rd.L()
		if err != nil {
			return nil, err
		}
		return &BeginNode{s, l}, nil
	case PRINT:
		e, err := rd.E()
		if err != nil {
			return nil, err
		}
		return &PrintNode{e}, nil
	}
	return nil, rd.Error(tk)
}

func (rd *RD) L() (LNode, error) {
	tk := rd.Next()
	if tk.Kind == END {
		return &LEndNode{}, nil
	} else if tk.Kind == SEMICOLON {
		s, err := rd.S()
		if err != nil {
			return nil, err
		}
		l, err := rd.L()
		if err != nil {
			return nil, err
		}
		return &LSemicolonNode{s, l}, nil
	}
	return nil, rd.Error(tk)
}

func (rd *RD) E() (ENode, error) {
	tkOne := rd.Next()
	if tkOne.Kind == NUM {
		tkTwo := rd.Next()
		if tkTwo.Kind == EQ {
			tkThree := rd.Next()
			if tkThree.Kind == NUM {
				return &EGeneralNode{tkOne.Value, tkThree.Value}, nil
			}
			return nil, rd.Error(tkThree)
		}
		return nil, rd.Error(tkTwo)
	}
	return nil, rd.Error(tkOne)
}

func TestParser() {
	rd := NewRD()
	tree, err := rd.S()
	fmt.Println(NodeString(tree))

	if err == nil {
		fmt.Println(NodeString(tree))
		//Accept(tree)
	} else {
		fmt.Println(err)
	}
}
