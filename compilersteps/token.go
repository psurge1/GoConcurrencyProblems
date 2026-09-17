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

type Value interface {
	ValueString() string
}

type (
	IntValue    int
	StringValue string
)

func (v IntValue) ValueString() string {
	return fmt.Sprintf("%d", v)
}

func (v StringValue) ValueString() string {
	return fmt.Sprintf("%s", v)
}

type Token struct {
	Kind  Kind
	Value Value
}

func (tk *Token) String() string {
	kind := tk.Kind.String()
	value := "uninitialized"
	if tk.Value != nil {
		value = tk.Value.ValueString()
	}
	return fmt.Sprintf("<%s : %s>", kind, value)
}

func TestToken() {
	tk := Token{NUM, IntValue(10)}
	tkTwo := Token{EQ, StringValue("=")}
	fmt.Println(tk.String())
	fmt.Println(tkTwo.String())
}
