package main

import (
	"bufio"
	"go/token"
	"io"

	"github.com/cznic/golex/lex"
)

// Allocate Character classes anywhere in [0x80, 0xFF].
const (
	classUnicodeLeter = iota + 0x80
	classUnicodeDigit
	classOther
)

type lexer struct {
	*lex.Lexer
	stateStack []int
}

func newLexer(src io.Reader, dst io.Writer, fName string) *lexer {
	file := token.NewFileSet().AddFile(fName, -1, 1<<31-1)
	lx, err := lex.New(file, bufio.NewReader(src), lex.RuneClass(rune2Class))
	if err != nil {
		panic(err)
	}
	return &lexer{lx, []int{0}}
}

func (l *lexer) ungetN(n int) []byte {
	l.Unget(l.Lookahead())

	chars := l.Token()

	for i := 1; i <= n; i++ {
		char := chars[len(chars)-i]
		l.Unget(char)
	}

	buf := l.TokenBytes(nil)
	buf = buf[:len(buf)-n]

	return buf
}

func (l *lexer) pushState(state int) {
	l.stateStack = append(l.stateStack, state)
}

func (l *lexer) popState() {
	len := len(l.stateStack)
	if len <= 1 {
		return
	}

	l.stateStack = l.stateStack[:len-1]
}

func (l *lexer) begin(state int) {
	len := len(l.stateStack)
	l.stateStack = l.stateStack[:len-1]
	l.stateStack = append(l.stateStack, state)
}

func (l *lexer) getCurrentState() int {
	return l.stateStack[len(l.stateStack)-1]
}