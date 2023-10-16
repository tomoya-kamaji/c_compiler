package main

import (
	"fmt"
	"strconv"
	"unicode"
)

// トークンの種類
type TokenType int

const (
	TOKEN_INT TokenType = iota
	TOKEN_PLUS
	TOKEN_MINUS
	TOKEN_MUL
	TOKEN_DIV
	TOKEN_LPAREN
	TOKEN_RPAREN
	TOKEN_EOF
)

// トークンを表現する構造体
type Token struct {
	Type  TokenType
	Value string
}

var tokens []Token
var pos int

// トークンを消費するための関数
func consume(t TokenType) bool {
	if tokens[pos].Type == t {
		pos++
		return true
	}
	return false
}

// トークンを期待するための関数
func expect(t TokenType) {
	if tokens[pos].Type != t {
		fmt.Printf("Expected token %v, but got %v\n", t, tokens[pos].Type)
		return
	}
	pos++
}

// 数値を期待するための関数
func expectNumber() int {
	if tokens[pos].Type != TOKEN_INT {
		fmt.Println("Number expected")
		return 0
	}
	val, _ := strconv.Atoi(tokens[pos].Value)
	pos++
	return val
}

// primary = num | "(" expr ")"
func primary() int {
	if consume(TOKEN_LPAREN) {
		val := expr()
		expect(TOKEN_RPAREN)
		return val
	}
	return expectNumber()
}

// mul = primary ("*" primary | "/" primary)*
func mul() int {
	val := primary()

	for {
		if consume(TOKEN_MUL) {
			val *= primary()
		} else if consume(TOKEN_DIV) {
			val /= primary()
		} else {
			break
		}
	}
	return val
}

// expr = mul ("+" mul | "-" mul)*
func expr() int {
	val := mul()

	for {
		if consume(TOKEN_PLUS) {
			val += mul()
		} else if consume(TOKEN_MINUS) {
			val -= mul()
		} else {
			break
		}
	}
	return val
}

// トークナイザ
func tokenize(input string) []Token {
	var tokens []Token

	for i := 0; i < len(input); i++ {
		switch {
		case unicode.IsDigit(rune(input[i])):
			j := i
			for j < len(input) && unicode.IsDigit(rune(input[j])) {
				j++
			}
			tokens = append(tokens, Token{TOKEN_INT, input[i:j]})
			i = j - 1 // update loop counter

		case input[i] == '+':
			tokens = append(tokens, Token{TOKEN_PLUS, string(input[i])})

		case input[i] == '-':
			tokens = append(tokens, Token{TOKEN_MINUS, string(input[i])})

		case input[i] == '*':
			tokens = append(tokens, Token{TOKEN_MUL, string(input[i])})

		case input[i] == '/':
			tokens = append(tokens, Token{TOKEN_DIV, string(input[i])})

		case input[i] == '(':
			tokens = append(tokens, Token{TOKEN_LPAREN, string(input[i])})

		case input[i] == ')':
			tokens = append(tokens, Token{TOKEN_RPAREN, string(input[i])})

		case unicode.IsSpace(rune(input[i])):
			// スペースを無視する

		default:
			fmt.Printf("Unexpected character: %c\n", input[i])
			return nil
		}
	}
	tokens = append(tokens, Token{TOKEN_EOF, ""}) // EOFトークンを追加
	return tokens
}

func main() {
	input := "2*(10-3)*3"
	tokens = tokenize(input)

	result := expr()
	fmt.Printf("Result: %d\n", result)
}
