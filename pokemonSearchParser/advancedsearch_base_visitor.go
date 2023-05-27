// Code generated from AdvancedSearch.g4 by ANTLR 4.13.0. DO NOT EDIT.

package pokemonSearchParser // AdvancedSearch
import "github.com/antlr4-go/antlr/v4"

type BaseAdvancedSearchVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseAdvancedSearchVisitor) VisitEntry(ctx *EntryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAdvancedSearchVisitor) VisitSimpleExpr(ctx *SimpleExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAdvancedSearchVisitor) VisitCommaExpr(ctx *CommaExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAdvancedSearchVisitor) VisitSingleValueExpr(ctx *SingleValueExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAdvancedSearchVisitor) VisitAndExpr(ctx *AndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAdvancedSearchVisitor) VisitIvValueExpr(ctx *IvValueExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAdvancedSearchVisitor) VisitRangeValueExpr(ctx *RangeValueExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAdvancedSearchVisitor) VisitIvRangeValueExpr(ctx *IvRangeValueExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAdvancedSearchVisitor) VisitNestedExpr(ctx *NestedExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAdvancedSearchVisitor) VisitOrExpr(ctx *OrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAdvancedSearchVisitor) VisitNegateExpr(ctx *NegateExprContext) interface{} {
	return v.VisitChildren(ctx)
}
