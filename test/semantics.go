package test

import (
    "fmt"
    "github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// SemanticListener checks for argument compatibility etc
type SemanticListener struct {
    BasetestListener
}

func NewSemanticListener() *SemanticListener {
    return &SemanticListener{}
}

func (t *SemanticListener) ExitQuery(_ antlr.ParserRuleContext) {
}

func (t *SemanticListener) ExitListExp(list *ListExpContext) {
    
    // Many ways we can do this
    //
    op := list.GetOp()
    l := list.List().AllElement()
    
    switch op.GetTokenType() {
    
    // Strings required
    //
    case testParserIN, testParserNIN,
	testParserHALL, testParserHANY, testParserHNONE:
	// These operators require a string list or an int list, but cannot mix
	//
	if !checkListType(l, testParserSTRING) {
	    // Wasn't a string list, is it an int list?
	    //
	    if !checkListType(l, testParserINT) {
		showError(op, "requires a list of strings or a list of integers - they cannot be mixed")
	    }
	}
    
    case testParserWAO:
	
	// This requires a list of lists, but every list must be  a list of integer or float
	//
	for _, e := range l {
	    
	    els := e.List().AllElement()
	    if e.GetStart().GetTokenType() != testParserLBRACKET {
		showError(op, "requires a list of lists of integers or floats")
		return
	    }
	    
	    // Check that every list in the list is a list of integers or floats
	    //
	    for _, el := range els {
		// Should probably check that this isn't yet another list and call an error about lists being 2 deep only
		if !checkListAny(el.List().AllElement(), testParserINT, testParserFLOAT) {
		    showError(op, "requires a list of lists of integers or floats")
		}
	    }
	}
    
    case testParserITM:
	// This requires a list of lists, but every list must be  a list of STRING
	//
	for _, e := range l {
	    
	    if e.GetStart().GetTokenType() != testParserLBRACKET {
		showError(op, "requires a list of lists of integers or floats")
		return
	    }
	    
	    // Check that every list in the list is a list of integers or floats
	    //
	    els := e.List().AllElement()
	    for _, el := range els {
		if !checkListType(el.List().AllElement(), testParserSTRING) {
		    showError(op, "requires a list of lists of integers or floats")
		}
	    }
	}
    default:
	showError(op, "is an illegal list operator")
    }
}

func checkListType(l []IElementContext, ttype int) bool {
    for _, e := range l {
	if e.GetStart().GetTokenType() != ttype {
	    return false
	}
    }
    return true
}

// Checks that all tokens in the  given element list are of one of the ttypes supplied. It is more
// efficient to use checkListType if the list must be all the same type
//
func checkListAny(l []IElementContext, ttype ...int) bool {
    for _, e := range l {
	pass := false
	for _, t := range ttype {
	    if e.GetStart().GetTokenType() == t {
		pass = true
		continue
	    }
	}
	// Exit early if we didn't find a match for this element
	if !pass {
	    return false
	}
    }
    return true
}

func showError(t antlr.Token, msg string) {
    fmt.Println("Error: Line", t.GetLine(), ", Column", t.GetColumn(), ": Operator ", "'"+t.GetText()+"'",
	msg)
    
}
