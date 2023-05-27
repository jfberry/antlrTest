// Code generated from AdvancedSearch.g4 by ANTLR 4.13.0. DO NOT EDIT.

package pokemonSearchParser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type AdvancedSearchLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var AdvancedSearchLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func advancedsearchlexerLexerInit() {
	staticData := &AdvancedSearchLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "','", "'!'", "'('", "')'", "'&'", "'|'", "'-'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "Atk", "Def", "Sta", "Level", "Number",
		"WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "Atk", "Def",
		"Sta", "Level", "Number", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 13, 61, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 4, 11, 51, 8, 11, 11, 11, 12, 11, 52,
		1, 12, 4, 12, 56, 8, 12, 11, 12, 12, 12, 57, 1, 12, 1, 12, 0, 0, 13, 1,
		1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 1, 0, 5, 2, 0, 65, 65, 97, 97, 2, 0, 68, 68, 100, 100,
		2, 0, 83, 83, 115, 115, 2, 0, 76, 76, 108, 108, 3, 0, 9, 10, 12, 13, 32,
		32, 62, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 1, 27, 1, 0, 0, 0, 3, 29, 1, 0, 0, 0,
		5, 31, 1, 0, 0, 0, 7, 33, 1, 0, 0, 0, 9, 35, 1, 0, 0, 0, 11, 37, 1, 0,
		0, 0, 13, 39, 1, 0, 0, 0, 15, 41, 1, 0, 0, 0, 17, 43, 1, 0, 0, 0, 19, 45,
		1, 0, 0, 0, 21, 47, 1, 0, 0, 0, 23, 50, 1, 0, 0, 0, 25, 55, 1, 0, 0, 0,
		27, 28, 5, 44, 0, 0, 28, 2, 1, 0, 0, 0, 29, 30, 5, 33, 0, 0, 30, 4, 1,
		0, 0, 0, 31, 32, 5, 40, 0, 0, 32, 6, 1, 0, 0, 0, 33, 34, 5, 41, 0, 0, 34,
		8, 1, 0, 0, 0, 35, 36, 5, 38, 0, 0, 36, 10, 1, 0, 0, 0, 37, 38, 5, 124,
		0, 0, 38, 12, 1, 0, 0, 0, 39, 40, 5, 45, 0, 0, 40, 14, 1, 0, 0, 0, 41,
		42, 7, 0, 0, 0, 42, 16, 1, 0, 0, 0, 43, 44, 7, 1, 0, 0, 44, 18, 1, 0, 0,
		0, 45, 46, 7, 2, 0, 0, 46, 20, 1, 0, 0, 0, 47, 48, 7, 3, 0, 0, 48, 22,
		1, 0, 0, 0, 49, 51, 2, 48, 57, 0, 50, 49, 1, 0, 0, 0, 51, 52, 1, 0, 0,
		0, 52, 50, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 24, 1, 0, 0, 0, 54, 56,
		7, 4, 0, 0, 55, 54, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0,
		57, 58, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 60, 6, 12, 0, 0, 60, 26, 1,
		0, 0, 0, 3, 0, 52, 57, 1, 6, 0, 0,
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

// AdvancedSearchLexerInit initializes any static state used to implement AdvancedSearchLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewAdvancedSearchLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func AdvancedSearchLexerInit() {
	staticData := &AdvancedSearchLexerLexerStaticData
	staticData.once.Do(advancedsearchlexerLexerInit)
}

// NewAdvancedSearchLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewAdvancedSearchLexer(input antlr.CharStream) *AdvancedSearchLexer {
	AdvancedSearchLexerInit()
	l := new(AdvancedSearchLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &AdvancedSearchLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "AdvancedSearch.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AdvancedSearchLexer tokens.
const (
	AdvancedSearchLexerT__0   = 1
	AdvancedSearchLexerT__1   = 2
	AdvancedSearchLexerT__2   = 3
	AdvancedSearchLexerT__3   = 4
	AdvancedSearchLexerT__4   = 5
	AdvancedSearchLexerT__5   = 6
	AdvancedSearchLexerT__6   = 7
	AdvancedSearchLexerAtk    = 8
	AdvancedSearchLexerDef    = 9
	AdvancedSearchLexerSta    = 10
	AdvancedSearchLexerLevel  = 11
	AdvancedSearchLexerNumber = 12
	AdvancedSearchLexerWS     = 13
)
