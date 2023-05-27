package main

import (
	"antlrTest/pokemonSearchParser"
	"github.com/antlr4-go/antlr/v4"
	"strconv"
)

type PokemonData struct {
	Iv    int
	Atk   int
	Def   int
	Sta   int
	Level int
}

type PokemonEvalVisitor struct {
	pokemonSearchParser.BaseAdvancedSearchVisitor
	pokemonData *PokemonData
}

func (v *PokemonEvalVisitor) Visit(tree antlr.ParseTree) bool {
	switch val := tree.(type) {
	case *pokemonSearchParser.EntryContext:
		return v.Visit(val.Expr())
	case *pokemonSearchParser.SimpleExprContext:
		return v.Visit(val.SingleExpr())
	case *pokemonSearchParser.NestedExprContext:
		return v.Visit(val.SingleExpr())
	case *pokemonSearchParser.CommaExprContext:
		return v.VisitCommaExpr(val)
	case *pokemonSearchParser.NegateExprContext:
		return !v.Visit(val.SingleExpr())
	case *pokemonSearchParser.OrExprContext:
		return v.VisitOrExpr(val)
	case *pokemonSearchParser.AndExprContext:
		return v.VisitAndExpr(val)
	case *pokemonSearchParser.IvValueExprContext:
		return v.VisitIvValueExpr(val)
	case *pokemonSearchParser.IvRangeValueExprContext:
		return v.VisitIvRangeValueExpr(val)
	case *pokemonSearchParser.RangeValueExprContext:
		return v.VisitRangeValueExpr(val)
	case *pokemonSearchParser.SingleValueExprContext:
		return v.VisitSingleValueExpr(val)
	}
	return false
}

func (v *PokemonEvalVisitor) VisitCommaExpr(ctx *pokemonSearchParser.CommaExprContext) bool {
	lhs := v.Visit(ctx.GetLhs())
	if lhs == true {
		return true // Shortcut evaluation
	}
	return v.Visit(ctx.GetRhs())
}

func (v *PokemonEvalVisitor) VisitOrExpr(ctx *pokemonSearchParser.OrExprContext) bool {
	lhs := v.Visit(ctx.GetLhs())
	if lhs == true {
		return true // Shortcut evaluation
	}
	return v.Visit(ctx.GetRhs())
}

func (v *PokemonEvalVisitor) VisitAndExpr(ctx *pokemonSearchParser.AndExprContext) bool {
	lhs := v.Visit(ctx.GetLhs())
	if lhs == false {
		return false // Shortcut evaluation
	}
	return v.Visit(ctx.GetRhs())
}

func (v *PokemonEvalVisitor) VisitIvValueExpr(ctx *pokemonSearchParser.IvValueExprContext) bool {
	iv, _ := strconv.ParseInt(ctx.GetVal().GetText(), 10, 64)

	return iv == int64(v.pokemonData.Iv)
}

func (v *PokemonEvalVisitor) VisitIvRangeValueExpr(ctx *pokemonSearchParser.IvRangeValueExprContext) bool {
	min, _ := strconv.ParseInt(ctx.GetMin().GetText(), 10, 64)
	iv64 := int64(v.pokemonData.Iv)
	if iv64 < min {
		return false
	}
	max, _ := strconv.ParseInt(ctx.GetMax().GetText(), 10, 64)
	return iv64 <= max
}

func (v *PokemonEvalVisitor) VisitSingleValueExpr(ctx *pokemonSearchParser.SingleValueExprContext) bool {
	testValue := int64(v.GetValueFromOp(ctx.GetOp()))

	val, _ := strconv.ParseInt(ctx.GetVal().GetText(), 10, 64)

	return val == testValue
}

func (v *PokemonEvalVisitor) VisitRangeValueExpr(ctx *pokemonSearchParser.RangeValueExprContext) bool {
	val := int64(v.GetValueFromOp(ctx.GetOp()))
	min, _ := strconv.ParseInt(ctx.GetMin().GetText(), 10, 64)
	if val < min {
		return false
	}
	max, _ := strconv.ParseInt(ctx.GetMax().GetText(), 10, 64)
	return val <= max
}

func (v *PokemonEvalVisitor) GetValueFromOp(op antlr.Token) int {
	switch op.GetTokenType() {
	case pokemonSearchParser.AdvancedSearchLexerAtk:
		return v.pokemonData.Atk
	case pokemonSearchParser.AdvancedSearchLexerDef:
		return v.pokemonData.Def
	case pokemonSearchParser.AdvancedSearchLexerSta:
		return v.pokemonData.Sta
	case pokemonSearchParser.AdvancedSearchLexerLevel:
		return v.pokemonData.Level
	default:
		panic("Unknown op")
	}
}
