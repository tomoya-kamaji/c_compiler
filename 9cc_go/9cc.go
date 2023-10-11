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

	i := 0
	for i < len(p) {
		ch := rune(p[i])
		if unicode.IsSpace(ch) {
			i++
			continue
		}

		if ch == '+' || ch == '-' {
			cur = newToken(TK_RESERVED, cur, string(ch))
			i++
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
		i++
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
	// printToken(token)

	fmt.Fprintln(file, ".global main")
	fmt.Fprintln(file, "main:")

	fmt.Fprintf(file, "  mov x0, %d\n", expectNumber())

	for !atEOF() {
		if consume('+') {
			fmt.Fprintf(file, "  add x0, x0, %d\n", expectNumber())
			continue
		}

		expect('-')
		fmt.Fprintf(file, "  sub x0, x0, %d\n", expectNumber())
	}

	fmt.Fprintln(file, "  ret")
}

func printToken(t *Token) {
	if t == nil {
		return
	}

	// Tokenの内容を出力
	fmt.Printf("Token - kind: %v, val: %d, str: %s\n", t.kind, t.val, t.str)

	// 次のTokenの内容も出力したい場合
	printToken(t.next)
}
