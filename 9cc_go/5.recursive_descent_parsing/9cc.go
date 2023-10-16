package main

import (
	"fmt"
	"strconv"
	"strings"
)

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

func (t TokenType) String() string {
	switch t {
	case TOKEN_INT:
		return "TOKEN_INT"
	case TOKEN_PLUS:
		return "TOKEN_PLUS"
	case TOKEN_MINUS:
		return "TOKEN_MINUS"
	case TOKEN_MUL:
		return "TOKEN_MUL"
	case TOKEN_DIV:
		return "TOKEN_DIV"
	case TOKEN_LPAREN:
		return "TOKEN_LPAREN"
	case TOKEN_RPAREN:
		return "TOKEN_RPAREN"
	case TOKEN_EOF:
		return "TOKEN_EOF"
	default:
		return "UNKNOWN_TOKEN_TYPE"
	}
}

type Token struct {
	Type  TokenType
	Value string
}

func tokenize(input string) []Token {
	var tokens []Token

	for i := 0; i < len(input); i++ {
		switch input[i] {
		case ' ':
			continue
		case '+':
			tokens = append(tokens, Token{TOKEN_PLUS, "+"})
		case '-':
			tokens = append(tokens, Token{TOKEN_MINUS, "-"})
		case '*':
			tokens = append(tokens, Token{TOKEN_MUL, "*"})
		case '/':
			tokens = append(tokens, Token{TOKEN_DIV, "/"})
		case '(':
			tokens = append(tokens, Token{TOKEN_LPAREN, "("})
		case ')':
			tokens = append(tokens, Token{TOKEN_RPAREN, ")"})
		default:
			start := i
			for i+1 < len(input) && strings.ContainsAny(string(input[i+1]), "0123456789") {
				i++
			}
			tokens = append(tokens, Token{TOKEN_INT, input[start : i+1]})
		}
	}

	tokens = append(tokens, Token{TOKEN_EOF, ""})
	return tokens
}

type Parser struct {
	tokens   []Token
	position int // その時点で見ているトークンの位置
}

// トークンを1つ読み進める
func (p *Parser) getToken() Token {
	token := p.tokens[p.position]
	p.position++
	return token
}

func (p *Parser) parsePrimary() int {
	token := p.getToken()
	switch token.Type {
	case TOKEN_INT:
		val, _ := strconv.Atoi(token.Value)
		return val
	case TOKEN_LPAREN: // 左括弧があればexprをパースして、カッコが閉じているか確認する
		val := p.parseExpr()
		if p.getToken().Type != TOKEN_RPAREN {
			panic("Expected closing parenthesis!")
		}
		return val
	default:
		panic("Expected integer or parenthesis!")
	}
}

// 乗算： 3*4
func (p *Parser) parseMul() int {
	val := p.parsePrimary()

	for {
		token := p.getToken()
		switch token.Type {
		case TOKEN_MUL:
			val *= p.parsePrimary()
		case TOKEN_DIV:
			val /= p.parsePrimary()
		default:
			p.position--
			return val
		}
	}
}

func (p *Parser) parseExpr() int {
	val := p.parseMul()

	for {
		token := p.getToken()
		switch token.Type {
		case TOKEN_PLUS:
			val += p.parseMul()
		case TOKEN_MINUS:
			val -= p.parseMul()
		default:
			p.position--
			return val
		}
	}
}

func main() {
	input := "3+4*10"

	tokens := tokenize(input)
	parser := Parser{tokens, 0}
	result := parser.parseExpr()

	fmt.Printf("%s = %d\n", input, result)
}

func printTokens(tokens []Token) {
	for i, token := range tokens {
		fmt.Printf("position:%v type:%s value:%s\n", i, token.Type.String(), token.Value)
	}
}
