package compilersteps

import (
	"fmt"
)

type Kind int

const (
	END Kind = iota
	IF
	THEN
	ELSE
	BEGIN
	PRINT
	SEMICOLON
	NUM
	EQ
)

func (k Kind) String() string {
	switch k {
	case END:
		return "END"
	case IF:
		return "IF"
	case THEN:
		return "THEN"
	case ELSE:
		return "ELSE"
	case BEGIN:
		return "BEGIN"
	case PRINT:
		return "PRINT"
	case SEMICOLON:
		return "SEMICOLON"
	case NUM:
		return "NUM"
	case EQ:
		return "EQ"
	default:
		return "UNKNOWN"
	}
}

type Token struct {
	Kind  Kind
	Value string
}

func (tk *Token) String() string {
	return fmt.Sprintf("<%s : %s>", tk.Kind.String(), tk.Value)
}
