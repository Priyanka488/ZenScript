package lexer

import (
	"fmt"
	"unicode"

	"github.com/Priyanka488/ZenScript/tokens"
)

const (
	ILLEGAL_CHAR_ERROR = "IllegalCharError"
)

// ************ Lexer ************
type Lexer struct {
	text        string
	pos         Position // current position in input (points to current char)
	currentChar byte     // current char

}

func New(text string, fileName string) *Lexer {
	lexer := &Lexer{
		text:        text,
		pos:         Position{Index: 0, Line: 0, Col: -1, FileName: fileName, FileText: text},
		currentChar: text[0],
	}
	return lexer
}

func (l *Lexer) advance() {
	l.pos.advance(string(l.currentChar))
	if l.pos.Index >= len(l.text) {
		l.currentChar = 0
	} else {
		l.currentChar = l.text[l.pos.Index]
	}
}

func (l *Lexer) generateNumber() *tokens.Token {
	result := ""
	dotCount := 0
	for l.currentChar != 0 && (unicode.IsDigit(rune(l.currentChar)) || l.currentChar == '.') {
		if l.currentChar == '.' {
			if dotCount == 1 {
				break
			}
			dotCount += 1
			result += "."
		} else {
			result += string(l.currentChar)
		}
		l.advance()
	}
	if dotCount == 0 {
		return tokens.New(tokens.TT_INT, result)
	} else {
		return tokens.New(tokens.TT_FLOAT, result)
	}
}

func (l *Lexer) generateTokens() ([]*tokens.Token, *Error) {
	currentTokens := []*tokens.Token{}
	for l.currentChar != 0 {
		switch {
		case l.currentChar == ' ':
			currentTokens = append(currentTokens, tokens.New(tokens.TT_SPACE, " "))
			l.advance()
		case l.currentChar == '+':
			currentTokens = append(currentTokens, tokens.New(tokens.TT_PLUS, "+"))
			l.advance()
		case l.currentChar == '-':
			currentTokens = append(currentTokens, tokens.New(tokens.TT_MINUS, "-"))
			l.advance()
		case l.currentChar == '*':
			currentTokens = append(currentTokens, tokens.New(tokens.TT_MUL, "*"))
			l.advance()
		case l.currentChar == '/':
			currentTokens = append(currentTokens, tokens.New(tokens.TT_DIV, "/"))
			l.advance()
		case l.currentChar == '(':
			currentTokens = append(currentTokens, tokens.New(tokens.TT_LPAREN, "("))
			l.advance()
		case l.currentChar == ')':
			currentTokens = append(currentTokens, tokens.New(tokens.TT_RPAREN, ")"))
			l.advance()
		case l.currentChar >= '0' && l.currentChar <= '9':
			currentTokens = append(currentTokens, l.generateNumber())
		default:
			l.advance()
			posStart := l.pos.Copy()
			message := "Illegal character: " + string(l.currentChar)
			return nil, NewIllegalCharError(message, posStart, l.pos).Error
		}
	}
	return currentTokens, nil
}

func Run(text string, fileName string) ([]*tokens.Token, *Error) {
	lexer := New(text, fileName)
	return lexer.generateTokens()
}

//  ************ Position ************

type Position struct {
	Index    int // current position in input (points to current char)
	Line     int // current line number
	Col      int // current column number
	FileName string
	FileText string
}

func (p *Position) advance(currentChar string) {
	p.Index += 1
	p.Col += 1
	if currentChar == "\n" {
		p.Line += 1
		p.Col = 0
	}
}

func (p *Position) Copy() Position {
	return Position{
		Index:    p.Index,
		Line:     p.Line,
		Col:      p.Col,
		FileName: p.FileName,
		FileText: p.FileText,
	}
}

//  ************ ERROR ************

type Error struct {
	Name          string
	Message       string
	PositionStart Position
	PositionEnd   Position
}

func NewError(name, message string, positionStart Position, positionEnd Position) *Error {
	return &Error{
		Name:          name,
		Message:       message,
		PositionStart: positionStart,
		PositionEnd:   positionEnd,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("ERROR: %s\nMessage:%s\nFile: %s\nLine: %d\nCol: %d\n", e.Name, e.Message, e.PositionStart.FileName, e.PositionStart.Line, e.PositionStart.Col)
}

type IllegalCharError struct {
	*Error
}

func NewIllegalCharError(message string, positionStart Position, positionEnd Position) *IllegalCharError {
	return &IllegalCharError{
		Error: NewError(ILLEGAL_CHAR_ERROR, message, positionStart, positionEnd),
	}
}
