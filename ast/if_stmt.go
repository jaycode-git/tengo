package ast

import "github.com/d5/tengo/source"

type IfStmt struct {
	IfPos source.Pos
	Init  Stmt
	Cond  Expr
	Body  *BlockStmt
	Else  Stmt // else branch; or nil
}

func (s *IfStmt) stmtNode() {}

func (s *IfStmt) Pos() source.Pos {
	return s.IfPos
}

func (s *IfStmt) End() source.Pos {
	if s.Else != nil {
		return s.Else.End()
	}

	return s.Body.End()
}

func (s *IfStmt) String() string {
	var initStmt, elseStmt string
	if s.Init != nil {
		initStmt = s.Init.String() + "; "
	}
	if s.Else != nil {
		elseStmt = " else " + s.Else.String()
	}

	return "if " + initStmt + s.Cond.String() + " " + s.Body.String() + elseStmt
}
