package ast

import (
	"github.com/kevguy/kevpretation/token"
)

// return <expression>;
type ReturnStatement struct {
	Token token.Token // the token.LET token
	Value Expression
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
