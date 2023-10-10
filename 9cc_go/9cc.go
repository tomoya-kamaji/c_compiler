package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type TokenKind int

const (
	TK_RESERVED TokenKind = iota
	TK_NUM
	TK_EOF
)

type Token struct {
	kind TokenKind
	next *Token
	val  int
	str  string
}

var user_input string
var token *Token

func errorAt(loc, fmtStr string, args ...interface{}) {
	pos := len(user_input) - len(loc)
	fmt.Fprintln(os.Stderr, user_input)
	fmt.Fprintf(os.Stderr, "%*s", pos, "")
	fmt.Fprintln(os.Stderr, "^")
	fmt.Fprintf(os.Stderr, fmtStr, args...)
	fmt.Fprintln(os.Stderr, "")
	os.Exit(1)
}

func consume(op rune) bool {
	if token.kind != TK_RESERVED || rune(token.str[0]) != op {
		return false
	}
	token = token.next
	return true
}

func expect(op rune) {
	if token.kind != TK_RESERVED || rune(token.str[0]) != op {
		errorAt(token.str, "expected '%c'", op)
	}
	token = token.next
}

func expectNumber() int {
	if token.kind != TK_NUM {
		errorAt(token.str, "expected a number")
	}
	val := token.val
	token = token.next
	return val
}

func atEOF() bool {
	return token.kind == TK_EOF
}

func newToken(kind TokenKind, cur *Token, str string) *Token {
	tok := &Token{
		kind: kind,
		str:  str,
	}
	cur.next = tok
	return tok
}

func tokenize() *Token {
	p := user_input
	head := &Token{}
	cur := head

	for i, ch := range p {
		if unicode.IsSpace(ch) {
			continue
		}

		if ch == '+' || ch == '-' {
			cur = newToken(TK_RESERVED, cur, string(ch))
			continue
		}

		if unicode.IsDigit(ch) {
			start := i
			for i < len(p) && unicode.IsDigit(rune(p[i])) {
				i++
			}
			numStr := p[start:i]
			cur = newToken(TK_NUM, cur, numStr)
			cur.val, _ = strconv.Atoi(numStr)
			continue
		}

		errorAt(string(ch), "unexpected character")
	}

	newToken(TK_EOF, cur, "")
	return head.next
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "%s: invalid number of arguments\n", os.Args[0])
		os.Exit(1)
	}
	// ファイルを開く（または新しく作成する）
	file, err := os.Create("tmp.s")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	user_input = os.Args[1]
	token = tokenize()

	fmt.Println(".global main")
	fmt.Println("main:")

	fmt.Printf("  mov x0, %d\n", expectNumber())

	for !atEOF() {
		if consume('+') {
			fmt.Printf("  add x0, x0, %d\n", expectNumber())
			continue
		}

		expect('-')
		fmt.Printf("  sub x0, x0, %d\n", expectNumber())
	}

	fmt.Println("  ret")
}
