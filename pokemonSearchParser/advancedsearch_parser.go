// Code generated from AdvancedSearch.g4 by ANTLR 4.13.0. DO NOT EDIT.

package pokemonSearchParser // AdvancedSearch
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type AdvancedSearchParser struct {
	*antlr.BaseParser
}

var AdvancedSearchParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func advancedsearchParserInit() {
	staticData := &AdvancedSearchParserStaticData
	staticData.LiteralNames = []string{
		"", "','", "'!'", "'('", "')'", "'&'", "'|'", "'-'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "Atk", "Def", "Sta", "Level", "Number",
		"WS",
	}
	staticData.RuleNames = []string{
		"entry", "expr", "singleExpr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 13, 47, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 15, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		3, 2, 34, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 42, 8, 2, 10,
		2, 12, 2, 45, 9, 2, 1, 2, 0, 1, 4, 3, 0, 2, 4, 0, 1, 1, 0, 8, 11, 51, 0,
		6, 1, 0, 0, 0, 2, 14, 1, 0, 0, 0, 4, 33, 1, 0, 0, 0, 6, 7, 3, 2, 1, 0,
		7, 8, 5, 0, 0, 1, 8, 1, 1, 0, 0, 0, 9, 15, 3, 4, 2, 0, 10, 11, 3, 4, 2,
		0, 11, 12, 5, 1, 0, 0, 12, 13, 3, 2, 1, 0, 13, 15, 1, 0, 0, 0, 14, 9, 1,
		0, 0, 0, 14, 10, 1, 0, 0, 0, 15, 3, 1, 0, 0, 0, 16, 17, 6, 2, -1, 0, 17,
		18, 5, 2, 0, 0, 18, 34, 3, 4, 2, 8, 19, 20, 5, 3, 0, 0, 20, 21, 3, 4, 2,
		0, 21, 22, 5, 4, 0, 0, 22, 34, 1, 0, 0, 0, 23, 24, 7, 0, 0, 0, 24, 25,
		5, 12, 0, 0, 25, 26, 5, 7, 0, 0, 26, 34, 5, 12, 0, 0, 27, 28, 7, 0, 0,
		0, 28, 34, 5, 12, 0, 0, 29, 30, 5, 12, 0, 0, 30, 31, 5, 7, 0, 0, 31, 34,
		5, 12, 0, 0, 32, 34, 5, 12, 0, 0, 33, 16, 1, 0, 0, 0, 33, 19, 1, 0, 0,
		0, 33, 23, 1, 0, 0, 0, 33, 27, 1, 0, 0, 0, 33, 29, 1, 0, 0, 0, 33, 32,
		1, 0, 0, 0, 34, 43, 1, 0, 0, 0, 35, 36, 10, 6, 0, 0, 36, 37, 5, 5, 0, 0,
		37, 42, 3, 4, 2, 7, 38, 39, 10, 5, 0, 0, 39, 40, 5, 6, 0, 0, 40, 42, 3,
		4, 2, 6, 41, 35, 1, 0, 0, 0, 41, 38, 1, 0, 0, 0, 42, 45, 1, 0, 0, 0, 43,
		41, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 5, 1, 0, 0, 0, 45, 43, 1, 0, 0,
		0, 4, 14, 33, 41, 43,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// AdvancedSearchParserInit initializes any static state used to implement AdvancedSearchParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAdvancedSearchParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AdvancedSearchParserInit() {
	staticData := &AdvancedSearchParserStaticData
	staticData.once.Do(advancedsearchParserInit)
}

// NewAdvancedSearchParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAdvancedSearchParser(input antlr.TokenStream) *AdvancedSearchParser {
	AdvancedSearchParserInit()
	this := new(AdvancedSearchParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &AdvancedSearchParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "AdvancedSearch.g4"

	return this
}

// AdvancedSearchParser tokens.
const (
	AdvancedSearchParserEOF    = antlr.TokenEOF
	AdvancedSearchParserT__0   = 1
	AdvancedSearchParserT__1   = 2
	AdvancedSearchParserT__2   = 3
	AdvancedSearchParserT__3   = 4
	AdvancedSearchParserT__4   = 5
	AdvancedSearchParserT__5   = 6
	AdvancedSearchParserT__6   = 7
	AdvancedSearchParserAtk    = 8
	AdvancedSearchParserDef    = 9
	AdvancedSearchParserSta    = 10
	AdvancedSearchParserLevel  = 11
	AdvancedSearchParserNumber = 12
	AdvancedSearchParserWS     = 13
)

// AdvancedSearchParser rules.
const (
	AdvancedSearchParserRULE_entry      = 0
	AdvancedSearchParserRULE_expr       = 1
	AdvancedSearchParserRULE_singleExpr = 2
)

// IEntryContext is an interface to support dynamic dispatch.
type IEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	EOF() antlr.TerminalNode

	// IsEntryContext differentiates from other interfaces.
	IsEntryContext()
}

type EntryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntryContext() *EntryContext {
	var p = new(EntryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AdvancedSearchParserRULE_entry
	return p
}

func InitEmptyEntryContext(p *EntryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AdvancedSearchParserRULE_entry
}

func (*EntryContext) IsEntryContext() {}

func NewEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntryContext {
	var p = new(EntryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdvancedSearchParserRULE_entry

	return p
}

func (s *EntryContext) GetParser() antlr.Parser { return s.parser }

func (s *EntryContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EntryContext) EOF() antlr.TerminalNode {
	return s.GetToken(AdvancedSearchParserEOF, 0)
}

func (s *EntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AdvancedSearchVisitor:
		return t.VisitEntry(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AdvancedSearchParser) Entry() (localctx IEntryContext) {
	localctx = NewEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AdvancedSearchParserRULE_entry)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(6)
		p.Expr()
	}
	{
		p.SetState(7)
		p.Match(AdvancedSearchParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AdvancedSearchParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AdvancedSearchParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdvancedSearchParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SimpleExprContext struct {
	ExprContext
}

func NewSimpleExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleExprContext {
	var p = new(SimpleExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *SimpleExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleExprContext) SingleExpr() ISingleExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleExprContext)
}

func (s *SimpleExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AdvancedSearchVisitor:
		return t.VisitSimpleExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CommaExprContext struct {
	ExprContext
	lhs ISingleExprContext
	rhs IExprContext
}

func NewCommaExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CommaExprContext {
	var p = new(CommaExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CommaExprContext) GetLhs() ISingleExprContext { return s.lhs }

func (s *CommaExprContext) GetRhs() IExprContext { return s.rhs }

func (s *CommaExprContext) SetLhs(v ISingleExprContext) { s.lhs = v }

func (s *CommaExprContext) SetRhs(v IExprContext) { s.rhs = v }

func (s *CommaExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommaExprContext) SingleExpr() ISingleExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleExprContext)
}

func (s *CommaExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CommaExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AdvancedSearchVisitor:
		return t.VisitCommaExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AdvancedSearchParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AdvancedSearchParserRULE_expr)
	p.SetState(14)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSimpleExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(9)
			p.singleExpr(0)
		}

	case 2:
		localctx = NewCommaExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(10)

			var _x = p.singleExpr(0)

			localctx.(*CommaExprContext).lhs = _x
		}
		{
			p.SetState(11)
			p.Match(AdvancedSearchParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(12)

			var _x = p.Expr()

			localctx.(*CommaExprContext).rhs = _x
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISingleExprContext is an interface to support dynamic dispatch.
type ISingleExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSingleExprContext differentiates from other interfaces.
	IsSingleExprContext()
}

type SingleExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleExprContext() *SingleExprContext {
	var p = new(SingleExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AdvancedSearchParserRULE_singleExpr
	return p
}

func InitEmptySingleExprContext(p *SingleExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AdvancedSearchParserRULE_singleExpr
}

func (*SingleExprContext) IsSingleExprContext() {}

func NewSingleExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleExprContext {
	var p = new(SingleExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdvancedSearchParserRULE_singleExpr

	return p
}

func (s *SingleExprContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleExprContext) CopyAll(ctx *SingleExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SingleExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SingleValueExprContext struct {
	SingleExprContext
	op  antlr.Token
	val antlr.Token
}

func NewSingleValueExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleValueExprContext {
	var p = new(SingleValueExprContext)

	InitEmptySingleExprContext(&p.SingleExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleExprContext))

	return p
}

func (s *SingleValueExprContext) GetOp() antlr.Token { return s.op }

func (s *SingleValueExprContext) GetVal() antlr.Token { return s.val }

func (s *SingleValueExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *SingleValueExprContext) SetVal(v antlr.Token) { s.val = v }

func (s *SingleValueExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleValueExprContext) Number() antlr.TerminalNode {
	return s.GetToken(AdvancedSearchParserNumber, 0)
}

func (s *SingleValueExprContext) Atk() antlr.TerminalNode {
	return s.GetToken(AdvancedSearchParserAtk, 0)
}

func (s *SingleValueExprContext) Def() antlr.TerminalNode {
	return s.GetToken(AdvancedSearchParserDef, 0)
}

func (s *SingleValueExprContext) Sta() antlr.TerminalNode {
	return s.GetToken(AdvancedSearchParserSta, 0)
}

func (s *SingleValueExprContext) Level() antlr.TerminalNode {
	return s.GetToken(AdvancedSearchParserLevel, 0)
}

func (s *SingleValueExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AdvancedSearchVisitor:
		return t.VisitSingleValueExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndExprContext struct {
	SingleExprContext
	lhs ISingleExprContext
	rhs ISingleExprContext
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	InitEmptySingleExprContext(&p.SingleExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleExprContext))

	return p
}

func (s *AndExprContext) GetLhs() ISingleExprContext { return s.lhs }

func (s *AndExprContext) GetRhs() ISingleExprContext { return s.rhs }

func (s *AndExprContext) SetLhs(v ISingleExprContext) { s.lhs = v }

func (s *AndExprContext) SetRhs(v ISingleExprContext) { s.rhs = v }

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AllSingleExpr() []ISingleExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleExprContext); ok {
			len++
		}
	}

	tst := make([]ISingleExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleExprContext); ok {
			tst[i] = t.(ISingleExprContext)
			i++
		}
	}

	return tst
}

func (s *AndExprContext) SingleExpr(i int) ISingleExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleExprContext)
}

func (s *AndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AdvancedSearchVisitor:
		return t.VisitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IvValueExprContext struct {
	SingleExprContext
	val antlr.Token
}

func NewIvValueExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IvValueExprContext {
	var p = new(IvValueExprContext)

	InitEmptySingleExprContext(&p.SingleExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleExprContext))

	return p
}

func (s *IvValueExprContext) GetVal() antlr.Token { return s.val }

func (s *IvValueExprContext) SetVal(v antlr.Token) { s.val = v }

func (s *IvValueExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IvValueExprContext) Number() antlr.TerminalNode {
	return s.GetToken(AdvancedSearchParserNumber, 0)
}

func (s *IvValueExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AdvancedSearchVisitor:
		return t.VisitIvValueExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type RangeValueExprContext struct {
	SingleExprContext
	op  antlr.Token
	min antlr.Token
	max antlr.Token
}

func NewRangeValueExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RangeValueExprContext {
	var p = new(RangeValueExprContext)

	InitEmptySingleExprContext(&p.SingleExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleExprContext))

	return p
}

func (s *RangeValueExprContext) GetOp() antlr.Token { return s.op }

func (s *RangeValueExprContext) GetMin() antlr.Token { return s.min }

func (s *RangeValueExprContext) GetMax() antlr.Token { return s.max }

func (s *RangeValueExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *RangeValueExprContext) SetMin(v antlr.Token) { s.min = v }

func (s *RangeValueExprContext) SetMax(v antlr.Token) { s.max = v }

func (s *RangeValueExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeValueExprContext) AllNumber() []antlr.TerminalNode {
	return s.GetTokens(AdvancedSearchParserNumber)
}

func (s *RangeValueExprContext) Number(i int) antlr.TerminalNode {
	return s.GetToken(AdvancedSearchParserNumber, i)
}

func (s *RangeValueExprContext) Atk() antlr.TerminalNode {
	return s.GetToken(AdvancedSearchParserAtk, 0)
}

func (s *RangeValueExprContext) Def() antlr.TerminalNode {
	return s.GetToken(AdvancedSearchParserDef, 0)
}

func (s *RangeValueExprContext) Sta() antlr.TerminalNode {
	return s.GetToken(AdvancedSearchParserSta, 0)
}

func (s *RangeValueExprContext) Level() antlr.TerminalNode {
	return s.GetToken(AdvancedSearchParserLevel, 0)
}

func (s *RangeValueExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AdvancedSearchVisitor:
		return t.VisitRangeValueExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IvRangeValueExprContext struct {
	SingleExprContext
	min antlr.Token
	max antlr.Token
}

func NewIvRangeValueExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IvRangeValueExprContext {
	var p = new(IvRangeValueExprContext)

	InitEmptySingleExprContext(&p.SingleExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleExprContext))

	return p
}

func (s *IvRangeValueExprContext) GetMin() antlr.Token { return s.min }

func (s *IvRangeValueExprContext) GetMax() antlr.Token { return s.max }

func (s *IvRangeValueExprContext) SetMin(v antlr.Token) { s.min = v }

func (s *IvRangeValueExprContext) SetMax(v antlr.Token) { s.max = v }

func (s *IvRangeValueExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IvRangeValueExprContext) AllNumber() []antlr.TerminalNode {
	return s.GetTokens(AdvancedSearchParserNumber)
}

func (s *IvRangeValueExprContext) Number(i int) antlr.TerminalNode {
	return s.GetToken(AdvancedSearchParserNumber, i)
}

func (s *IvRangeValueExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AdvancedSearchVisitor:
		return t.VisitIvRangeValueExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NestedExprContext struct {
	SingleExprContext
}

func NewNestedExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NestedExprContext {
	var p = new(NestedExprContext)

	InitEmptySingleExprContext(&p.SingleExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleExprContext))

	return p
}

func (s *NestedExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedExprContext) SingleExpr() ISingleExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleExprContext)
}

func (s *NestedExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AdvancedSearchVisitor:
		return t.VisitNestedExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrExprContext struct {
	SingleExprContext
	lhs ISingleExprContext
	rhs ISingleExprContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	InitEmptySingleExprContext(&p.SingleExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleExprContext))

	return p
}

func (s *OrExprContext) GetLhs() ISingleExprContext { return s.lhs }

func (s *OrExprContext) GetRhs() ISingleExprContext { return s.rhs }

func (s *OrExprContext) SetLhs(v ISingleExprContext) { s.lhs = v }

func (s *OrExprContext) SetRhs(v ISingleExprContext) { s.rhs = v }

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) AllSingleExpr() []ISingleExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleExprContext); ok {
			len++
		}
	}

	tst := make([]ISingleExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleExprContext); ok {
			tst[i] = t.(ISingleExprContext)
			i++
		}
	}

	return tst
}

func (s *OrExprContext) SingleExpr(i int) ISingleExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleExprContext)
}

func (s *OrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AdvancedSearchVisitor:
		return t.VisitOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NegateExprContext struct {
	SingleExprContext
}

func NewNegateExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegateExprContext {
	var p = new(NegateExprContext)

	InitEmptySingleExprContext(&p.SingleExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleExprContext))

	return p
}

func (s *NegateExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegateExprContext) SingleExpr() ISingleExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleExprContext)
}

func (s *NegateExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AdvancedSearchVisitor:
		return t.VisitNegateExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AdvancedSearchParser) SingleExpr() (localctx ISingleExprContext) {
	return p.singleExpr(0)
}

func (p *AdvancedSearchParser) singleExpr(_p int) (localctx ISingleExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewSingleExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ISingleExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, AdvancedSearchParserRULE_singleExpr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNegateExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(17)
			p.Match(AdvancedSearchParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(18)
			p.singleExpr(8)
		}

	case 2:
		localctx = NewNestedExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(19)
			p.Match(AdvancedSearchParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(20)
			p.singleExpr(0)
		}
		{
			p.SetState(21)
			p.Match(AdvancedSearchParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewRangeValueExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(23)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*RangeValueExprContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3840) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*RangeValueExprContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(24)

			var _m = p.Match(AdvancedSearchParserNumber)

			localctx.(*RangeValueExprContext).min = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(25)
			p.Match(AdvancedSearchParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(26)

			var _m = p.Match(AdvancedSearchParserNumber)

			localctx.(*RangeValueExprContext).max = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewSingleValueExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(27)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*SingleValueExprContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3840) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*SingleValueExprContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(28)

			var _m = p.Match(AdvancedSearchParserNumber)

			localctx.(*SingleValueExprContext).val = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewIvRangeValueExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(29)

			var _m = p.Match(AdvancedSearchParserNumber)

			localctx.(*IvRangeValueExprContext).min = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(30)
			p.Match(AdvancedSearchParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(31)

			var _m = p.Match(AdvancedSearchParserNumber)

			localctx.(*IvRangeValueExprContext).max = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewIvValueExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(32)

			var _m = p.Match(AdvancedSearchParserNumber)

			localctx.(*IvValueExprContext).val = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(41)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndExprContext(p, NewSingleExprContext(p, _parentctx, _parentState))
				localctx.(*AndExprContext).lhs = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AdvancedSearchParserRULE_singleExpr)
				p.SetState(35)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(36)
					p.Match(AdvancedSearchParserT__4)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(37)

					var _x = p.singleExpr(7)

					localctx.(*AndExprContext).rhs = _x
				}

			case 2:
				localctx = NewOrExprContext(p, NewSingleExprContext(p, _parentctx, _parentState))
				localctx.(*OrExprContext).lhs = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AdvancedSearchParserRULE_singleExpr)
				p.SetState(38)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(39)
					p.Match(AdvancedSearchParserT__5)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(40)

					var _x = p.singleExpr(6)

					localctx.(*OrExprContext).rhs = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *AdvancedSearchParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *SingleExprContext = nil
		if localctx != nil {
			t = localctx.(*SingleExprContext)
		}
		return p.SingleExpr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *AdvancedSearchParser) SingleExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
