package parser

import (
	"fmt"
	"strings"
)

/*
	Trace parser steps. To use

	1. Add following

		func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
			defer untrace(trace("parseExpressionStatement"))
		// [...]
		}

		func (p *Parser) parseExpression(precedence int) ast.Expression {
			defer untrace(trace("parseExpression"))
		// [...]
		}

		func (p *Parser) parseIntegerLiteral() ast.Expression {
			defer untrace(trace("parseIntegerLiteral"))
		// [...]
		}

		func (p *Parser) parsePrefixExpression() ast.Expression {
			defer untrace(trace("parsePrefixExpression"))
		// [...]
		}

		func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
			defer untrace(trace("parseInfixExpression"))
		// [...]
		}


	2. Then run:
		$ go test -v -run TestOperatorPrecedenceParsing ./parser

*/

var traceLevel int = 0

const traceIdentPlaceholder string = "\t"

func identLevel() string {
	return strings.Repeat(traceIdentPlaceholder, traceLevel-1)
}

func tracePrint(fs string) {
	fmt.Printf("%s%s\n", identLevel(), fs)
}

func incIdent() { traceLevel = traceLevel + 1 }
func decIdent() { traceLevel = traceLevel - 1 }

func trace(msg string) string {
	incIdent()
	tracePrint("BEGIN " + msg)
	return msg
}

func untrace(msg string) {
	tracePrint("END " + msg)
	decIdent()
}
