package ast

import (
	"github.com/atushi-koga/interpreter/token"
)

type Node interface {
	TokenLiteral() string
}

// @todo: interfaceが入れ子になるってどういうこと？interfaceを継承するってこと？
type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.statements) > 0 {
		return p.statements[0].TokenLiteral()
	}

	return ""
}

func (p Program) Statements() []Statement {
	return p.statements
}

type LetStatement struct {
	token token.Token // token.LET トークン
	name  *Identifier
	value Expression
}

func (ls *LetStatement) statementNode() {

}

func (ls *LetStatement) TokenLiteral() string {
	return ls.token.Literal
}

// メソッド名を 'name' にすると、ls.name はメソッドを返すようになり、属性アクセスができない。（これはPHPも同じ？）
// ゲッターのメソッド名を属性名とは異なるものにする
func (ls LetStatement) Name() *Identifier {
	return ls.name
}

type Identifier struct {
	token token.Token // token.IDENT トークン
	value string
}

func (i *Identifier) expressionNode() {

}

func (i *Identifier) TokenLiteral() string {
	return i.token.Literal
}

func (i Identifier) Value() string {
	return i.value
}
