// Code generated from AdvancedSearch.g4 by ANTLR 4.13.0. DO NOT EDIT.

package pokemonSearchParser // AdvancedSearch
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by AdvancedSearchParser.
type AdvancedSearchVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by AdvancedSearchParser#entry.
	VisitEntry(ctx *EntryContext) interface{}

	// Visit a parse tree produced by AdvancedSearchParser#SimpleExpr.
	VisitSimpleExpr(ctx *SimpleExprContext) interface{}

	// Visit a parse tree produced by AdvancedSearchParser#CommaExpr.
	VisitCommaExpr(ctx *CommaExprContext) interface{}

	// Visit a parse tree produced by AdvancedSearchParser#SingleValueExpr.
	VisitSingleValueExpr(ctx *SingleValueExprContext) interface{}

	// Visit a parse tree produced by AdvancedSearchParser#AndExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by AdvancedSearchParser#IvValueExpr.
	VisitIvValueExpr(ctx *IvValueExprContext) interface{}

	// Visit a parse tree produced by AdvancedSearchParser#RangeValueExpr.
	VisitRangeValueExpr(ctx *RangeValueExprContext) interface{}

	// Visit a parse tree produced by AdvancedSearchParser#IvRangeValueExpr.
	VisitIvRangeValueExpr(ctx *IvRangeValueExprContext) interface{}

	// Visit a parse tree produced by AdvancedSearchParser#NestedExpr.
	VisitNestedExpr(ctx *NestedExprContext) interface{}

	// Visit a parse tree produced by AdvancedSearchParser#OrExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by AdvancedSearchParser#NegateExpr.
	VisitNegateExpr(ctx *NegateExprContext) interface{}
}
