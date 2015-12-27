package lexer

import "unicode/utf8"
import "fmt"

type TokenType int

type Token struct {
	typ TokenType
	val interface{}
}

const (
	// specials
	NewLine = iota
	EOF

	// parens
	OpenBracket
	CloseBracket

	// types
	Integer

	// operators
	Plus
	Minus
)

type Lexer struct {
	input string
}

// public

// NewLexer Lexer constructor
func NewLexer(input string) *Lexer {
	l := new(Lexer)
	l.input = input

	return l
}

// GetTokens get token
func (lexer *Lexer) GetTokens() ([]*Token, error) {
	var tokens []*Token

	for lexer.inputBytesRemaining() > 0 {
		token, err := lexer.nextToken()

		if err != nil {
			return nil, err
		}

		tokens = append(tokens, token)
	}

	return tokens, nil
}

// private

func (lexer *Lexer) inputBytesRemaining() int64 {
	return int64(len(lexer.input))
}

func (lexer *Lexer) nextToken() (*Token, error) {
	r, size := utf8.DecodeRuneInString(lexer.input)
	lexer.input = lexer.input[size:]

	token := new(Token)

	switch r {
	case ' ':
		nextToken, err := lexer.nextToken()

		if err != nil {
			token = nextToken
			break
		}
		return nil, err

	case '1', '2', '3', '4', '5', '6', '7', '8', '9':
		token.typ = Integer
		token.val = int(r)
	case '(':
		token.typ = OpenBracket
	case ')':
		token.typ = CloseBracket
	case '+':
		token.typ = Plus
	default:
		return nil, fmt.Errorf("Unknown token %#U", r)
	}

	return token, nil
}
