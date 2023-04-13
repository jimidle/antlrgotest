package main

import (
    "fmt"
    "github.com/antlr/antlr4/runtime/Go/antlr/v4"
    "os"
    "testrig/test"
)

func main() {
    testRun("input")
}

func testRun(inf string) {
    
    // Pre-initialize so that we can distinguish this initialization from the lexing nad parsing rules
    test.TestLexerInit()
    test.TestParserInit()
    
    input, err := antlr.NewFileStream(inf)
    if err != nil {
	fmt.Println("Error opening file '", inf, "'", err)
	os.Exit(1)
    }
    
    lexer := test.NewtestLexer(input)
    stream := antlr.NewCommonTokenStream(lexer, 0)
    p := test.NewtestParser(stream)
    
    // Add a diagnostic listener to get messages about conflicts in the grammar - debug only!
    //
    p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
    
    // We want the tree, as we are going to check it for semantics, then execute it
    //
    p.BuildParseTrees = true
    tree := p.Query()
    // Perform the semantic checks
    //
    antlr.ParseTreeWalkerDefault.Walk(test.NewSemanticListener(), tree)
    
    // Here we can execute the tree with a different listener
}
