package ast

import (
	"github.com/kevguy/kevpretation/token"
)

// let IDENT = <expression>
type LetStatement struct {
	Token token.Token // the token.LET token
	Name *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}


