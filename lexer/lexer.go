package lexer

import (
	"fmt"
)

func ManualRegexDFA(text string) bool {
	//regex := "a|bax*"
	dfa := map[int]map[rune]int{
		-1: {},               // error state
		0:  {'a': 1, 'b': 1}, // start state
		1:  {'a': 2},
		2:  {'x': 2}, // final state
	}
	finalState := 2

	currState := 0
	for _, chr := range text {
		nextState, ok := dfa[currState][chr]
		if ok {
			currState = nextState
		} else {
			currState = -1
			break
		}
	}

	return currState == finalState
}

func Test() {
	fmt.Println("Running Basic Regex as DFA:")
	fmt.Println("Testing regex matching for a|bax*")

	matchingStrings := []string{
		"aa",
		"ba",
		"aax",
		"bax",
		"aaxxx",
		"baxxxxxx",
	}

	nonMatchingStrings := []string{
		"",
		"a",
		"b",
		"ab",
		"bb",
		"xxx",
		"ax",
		"bx",
		"aaxa",
		"baxxb",
		"aax123",
	}
	failedTests := 0
	for _, value := range matchingStrings {
		if !ManualRegexDFA(value) {
			fmt.Printf("Failed test %s should match\n", value)
			failedTests += 1
		}
	}
	for _, value := range nonMatchingStrings {
		if ManualRegexDFA(value) {
			fmt.Printf("Failed test %s shouldn't match\n", value)
			failedTests += 1
		}
	}
	if failedTests == 0 {
		fmt.Println("ALL TESTS PASSED")
	}
	fmt.Println("")

	fmt.Println("Running Generic Regex to NFA to DFA:")

	fmt.Println("Running Regex matcher:")

	fmt.Println("Running Lexer:")
}
