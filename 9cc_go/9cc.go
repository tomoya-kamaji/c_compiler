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
	position int
}

func (p *Parser) getNextToken() Token {
	tok := p.tokens[p.position]
	p.position++
	return tok
}

func (p *Parser) parsePrimary() int {
	token := p.getNextToken()
	switch token.Type {
	case TOKEN_INT:
		val, _ := strconv.Atoi(token.Value)
		return val
	case TOKEN_LPAREN:
		val := p.parseExpr()
		if p.getNextToken().Type != TOKEN_RPAREN {
			panic("Expected closing parenthesis!")
		}
		return val
	default:
		panic("Expected integer or parenthesis!")
	}
}

// 掛け算
func (p *Parser) parseMul() int {
	val := p.parsePrimary()

	for {
		token := p.getNextToken()
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
		token := p.getNextToken()
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
	input := "3+4*10" // You can change this to any expression that fits the given grammar

	tokens := tokenize(input)
	parser := Parser{tokens, 0}
	result := parser.parseExpr()

	fmt.Printf("%s = %d\n", input, result)
}
