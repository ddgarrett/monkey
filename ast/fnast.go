package ast

import (
	"bytes"
	"monkey/token"
	"strings"
)

// Function Literal defines a function of the form:
//
//	fn(x, y) {
//		return x + y;
//	}
//
// See also parser package functions in fnparser.go to parse a function:
//  - func (p *Parser) parseFunctionLiteral() ast.Expression
//  - func (p *Parser) parseFunctionParameters() []*ast.Identifier
//
type FunctionLiteral struct {
	Token      token.Token // The 'fn' token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}
