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
type RD struct {
	tokens    []Token
	idx       int
	numTokens int
}

func NewRD(tokens []Token) RD {
	return RD{tokens, 0, len(tokens)}
}

func (rd *RD) Error(tk Token) error {
	return fmt.Errorf("syntax error on token: %s", tk.String())
}

func (rd *RD) Next() Token {
	if rd.idx >= rd.numTokens {
		return Token{}
	}
	tk := rd.tokens[rd.idx]
	rd.idx += 1
	return tk
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
		return &SIfNode{e, sOne, sTwo}, nil
	case BEGIN:
		s, err := rd.S()
		if err != nil {
			return nil, err
		}
		l, err := rd.L()
		if err != nil {
			return nil, err
		}
		return &SBeginNode{s, l}, nil
	case PRINT:
		e, err := rd.E()
		if err != nil {
			return nil, err
		}
		return &SPrintNode{e}, nil
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
	tokens := []Token{
		{BEGIN, nil},
		{IF, nil},
		{NUM, IntValue(5)},
		{EQ, nil},
		{NUM, IntValue(5)},
		{THEN, nil},
		{PRINT, nil},
		{NUM, IntValue(5)},
		{EQ, nil},
		{NUM, IntValue(5)},
		{ELSE, nil},
		{PRINT, nil},
		{NUM, IntValue(10)},
		{EQ, nil},
		{NUM, IntValue(10)},
		{END, nil},
	}
	rd := NewRD(tokens)
	tree, err := rd.S()

	if err == nil {
		DFSParser(tree)
		//Accept(tree)
	} else {
		fmt.Println(err)
	}
}
