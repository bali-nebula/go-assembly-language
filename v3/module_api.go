/*
................................................................................
.    Copyright (c) 2009-2026 Crater Dog Technologies™.  All Rights Reserved.   .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

/*
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│         This "module_api.go" file was automatically generated using:         │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│      Updates to any part of this file—other than the Module Description      │
│             and the Global Functions sections may be overwritten.            │
└──────────────────────────────────────────────────────────────────────────────┘

Package "module" declares type aliases for the commonly used types declared in
the packages contained in this module.  It also provides constructors for each
commonly used class that is exported by the module.  Each constructor delegates
the actual construction process to its corresponding concrete class declared in
the corresponding package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - https://github.com/bali-nebula/go-assembly-language/wiki
*/
package module

import (
	ast "github.com/bali-nebula/go-assembly-language/v3/ast"
	gra "github.com/bali-nebula/go-assembly-language/v3/grammar"
	com "github.com/craterdog/go-essential-composites/v8"
)

// TYPE ALIASES

// Ast

type (
	ActionClassLike      = ast.ActionClassLike
	ArgumentClassLike    = ast.ArgumentClassLike
	CallClassLike        = ast.CallClassLike
	ComponentClassLike   = ast.ComponentClassLike
	ConditionClassLike   = ast.ConditionClassLike
	ConstantClassLike    = ast.ConstantClassLike
	ContextClassLike     = ast.ContextClassLike
	DestinationClassLike = ast.DestinationClassLike
	DropClassLike        = ast.DropClassLike
	HandlerClassLike     = ast.HandlerClassLike
	InstructionClassLike = ast.InstructionClassLike
	JumpClassLike        = ast.JumpClassLike
	LiteralClassLike     = ast.LiteralClassLike
	LoadClassLike        = ast.LoadClassLike
	MethodClassLike      = ast.MethodClassLike
	NoteClassLike        = ast.NoteClassLike
	PrefixClassLike      = ast.PrefixClassLike
	PullClassLike        = ast.PullClassLike
	PushClassLike        = ast.PushClassLike
	SaveClassLike        = ast.SaveClassLike
	SendClassLike        = ast.SendClassLike
	SkipClassLike        = ast.SkipClassLike
	SourceClassLike      = ast.SourceClassLike
	ValueClassLike       = ast.ValueClassLike
)

type (
	ActionLike      = ast.ActionLike
	ArgumentLike    = ast.ArgumentLike
	CallLike        = ast.CallLike
	ComponentLike   = ast.ComponentLike
	ConditionLike   = ast.ConditionLike
	ConstantLike    = ast.ConstantLike
	ContextLike     = ast.ContextLike
	DestinationLike = ast.DestinationLike
	DropLike        = ast.DropLike
	HandlerLike     = ast.HandlerLike
	InstructionLike = ast.InstructionLike
	JumpLike        = ast.JumpLike
	LiteralLike     = ast.LiteralLike
	LoadLike        = ast.LoadLike
	MethodLike      = ast.MethodLike
	NoteLike        = ast.NoteLike
	PrefixLike      = ast.PrefixLike
	PullLike        = ast.PullLike
	PushLike        = ast.PushLike
	SaveLike        = ast.SaveLike
	SendLike        = ast.SendLike
	SkipLike        = ast.SkipLike
	SourceLike      = ast.SourceLike
	ValueLike       = ast.ValueLike
)

// Grammar

type (
	TokenType = gra.TokenType
)

const (
	ErrorToken       = gra.ErrorToken
	DelimiterToken   = gra.DelimiterToken
	DescriptionToken = gra.DescriptionToken
	LabelToken       = gra.LabelToken
	NewlineToken     = gra.NewlineToken
	QuotedToken      = gra.QuotedToken
	SpaceToken       = gra.SpaceToken
	SymbolToken      = gra.SymbolToken
)

type (
	FormatterClassLike = gra.FormatterClassLike
	ParserClassLike    = gra.ParserClassLike
	ProcessorClassLike = gra.ProcessorClassLike
	ScannerClassLike   = gra.ScannerClassLike
	TokenClassLike     = gra.TokenClassLike
	ValidatorClassLike = gra.ValidatorClassLike
	VisitorClassLike   = gra.VisitorClassLike
)

type (
	FormatterLike = gra.FormatterLike
	ParserLike    = gra.ParserLike
	ProcessorLike = gra.ProcessorLike
	ScannerLike   = gra.ScannerLike
	TokenLike     = gra.TokenLike
	ValidatorLike = gra.ValidatorLike
	VisitorLike   = gra.VisitorLike
)

type (
	Methodical = gra.Methodical
)

// CLASS ACCESSORS

// Ast

func ActionClass() ActionClassLike {
	return ast.ActionClass()
}

func Action(
	any_ any,
) ActionLike {
	return ActionClass().Action(
		any_,
	)
}

func ArgumentClass() ArgumentClassLike {
	return ast.ArgumentClass()
}

func Argument(
	delimiter string,
	symbol string,
) ArgumentLike {
	return ArgumentClass().Argument(
		delimiter,
		symbol,
	)
}

func CallClass() CallClassLike {
	return ast.CallClass()
}

func Call(
	delimiter string,
	symbol string,
	optionalContext ast.ContextLike,
) CallLike {
	return CallClass().Call(
		delimiter,
		symbol,
		optionalContext,
	)
}

func ComponentClass() ComponentClassLike {
	return ast.ComponentClass()
}

func Component(
	any_ any,
) ComponentLike {
	return ComponentClass().Component(
		any_,
	)
}

func ConditionClass() ConditionClassLike {
	return ast.ConditionClass()
}

func Condition(
	any_ any,
) ConditionLike {
	return ConditionClass().Condition(
		any_,
	)
}

func ConstantClass() ConstantClassLike {
	return ast.ConstantClass()
}

func Constant(
	delimiter string,
	symbol string,
) ConstantLike {
	return ConstantClass().Constant(
		delimiter,
		symbol,
	)
}

func ContextClass() ContextClassLike {
	return ast.ContextClass()
}

func Context(
	any_ any,
) ContextLike {
	return ContextClass().Context(
		any_,
	)
}

func DestinationClass() DestinationClassLike {
	return ast.DestinationClass()
}

func Destination(
	any_ any,
) DestinationLike {
	return DestinationClass().Destination(
		any_,
	)
}

func DropClass() DropClassLike {
	return ast.DropClass()
}

func Drop(
	delimiter string,
	component ast.ComponentLike,
	symbol string,
) DropLike {
	return DropClass().Drop(
		delimiter,
		component,
		symbol,
	)
}

func HandlerClass() HandlerClassLike {
	return ast.HandlerClass()
}

func Handler(
	delimiter string,
	label string,
) HandlerLike {
	return HandlerClass().Handler(
		delimiter,
		label,
	)
}

func InstructionClass() InstructionClassLike {
	return ast.InstructionClass()
}

func Instruction(
	optionalPrefix ast.PrefixLike,
	action ast.ActionLike,
) InstructionLike {
	return InstructionClass().Instruction(
		optionalPrefix,
		action,
	)
}

func JumpClass() JumpClassLike {
	return ast.JumpClass()
}

func Jump(
	delimiter string,
	label string,
	optionalCondition ast.ConditionLike,
) JumpLike {
	return JumpClass().Jump(
		delimiter,
		label,
		optionalCondition,
	)
}

func LiteralClass() LiteralClassLike {
	return ast.LiteralClass()
}

func Literal(
	delimiter string,
	quoted string,
) LiteralLike {
	return LiteralClass().Literal(
		delimiter,
		quoted,
	)
}

func LoadClass() LoadClassLike {
	return ast.LoadClass()
}

func Load(
	delimiter string,
	component ast.ComponentLike,
	symbol string,
) LoadLike {
	return LoadClass().Load(
		delimiter,
		component,
		symbol,
	)
}

func MethodClass() MethodClassLike {
	return ast.MethodClass()
}

func Method(
	instructions com.Sequential[ast.InstructionLike],
) MethodLike {
	return MethodClass().Method(
		instructions,
	)
}

func NoteClass() NoteClassLike {
	return ast.NoteClass()
}

func Note(
	delimiter string,
	description string,
) NoteLike {
	return NoteClass().Note(
		delimiter,
		description,
	)
}

func PrefixClass() PrefixClassLike {
	return ast.PrefixClass()
}

func Prefix(
	label string,
	delimiter string,
) PrefixLike {
	return PrefixClass().Prefix(
		label,
		delimiter,
	)
}

func PullClass() PullClassLike {
	return ast.PullClass()
}

func Pull(
	delimiter string,
	value ast.ValueLike,
) PullLike {
	return PullClass().Pull(
		delimiter,
		value,
	)
}

func PushClass() PushClassLike {
	return ast.PushClass()
}

func Push(
	delimiter string,
	source ast.SourceLike,
) PushLike {
	return PushClass().Push(
		delimiter,
		source,
	)
}

func SaveClass() SaveClassLike {
	return ast.SaveClass()
}

func Save(
	delimiter string,
	component ast.ComponentLike,
	symbol string,
) SaveLike {
	return SaveClass().Save(
		delimiter,
		component,
		symbol,
	)
}

func SendClass() SendClassLike {
	return ast.SendClass()
}

func Send(
	delimiter1 string,
	symbol string,
	delimiter2 string,
	destination ast.DestinationLike,
) SendLike {
	return SendClass().Send(
		delimiter1,
		symbol,
		delimiter2,
		destination,
	)
}

func SkipClass() SkipClassLike {
	return ast.SkipClass()
}

func Skip(
	delimiter string,
) SkipLike {
	return SkipClass().Skip(
		delimiter,
	)
}

func SourceClass() SourceClassLike {
	return ast.SourceClass()
}

func Source(
	any_ any,
) SourceLike {
	return SourceClass().Source(
		any_,
	)
}

func ValueClass() ValueClassLike {
	return ast.ValueClass()
}

func Value(
	any_ any,
) ValueLike {
	return ValueClass().Value(
		any_,
	)
}

// Grammar

func FormatterClass() FormatterClassLike {
	return gra.FormatterClass()
}

func Formatter() FormatterLike {
	return FormatterClass().Formatter()
}

func ParserClass() ParserClassLike {
	return gra.ParserClass()
}

func Parser() ParserLike {
	return ParserClass().Parser()
}

func ProcessorClass() ProcessorClassLike {
	return gra.ProcessorClass()
}

func Processor() ProcessorLike {
	return ProcessorClass().Processor()
}

func ScannerClass() ScannerClassLike {
	return gra.ScannerClass()
}

func Scanner(
	source string,
	tokens com.QueueLike[gra.TokenLike],
) ScannerLike {
	return ScannerClass().Scanner(
		source,
		tokens,
	)
}

func TokenClass() TokenClassLike {
	return gra.TokenClass()
}

func Token(
	line uint,
	position uint,
	type_ gra.TokenType,
	value string,
) TokenLike {
	return TokenClass().Token(
		line,
		position,
		type_,
		value,
	)
}

func ValidatorClass() ValidatorClassLike {
	return gra.ValidatorClass()
}

func Validator() ValidatorLike {
	return ValidatorClass().Validator()
}

func VisitorClass() VisitorClassLike {
	return gra.VisitorClass()
}

func Visitor(
	processor gra.Methodical,
) VisitorLike {
	return VisitorClass().Visitor(
		processor,
	)
}

// GLOBAL FUNCTIONS

func FormatMethod(
	method MethodLike,
) string {
	var formatter = Formatter()
	return formatter.FormatMethod(method)
}

func MatchesType(
	tokenValue string,
	tokenType TokenType,
) bool {
	var scannerClass = gra.ScannerClass()
	return scannerClass.MatchesType(tokenValue, tokenType)
}

func ParseSource(
	source string,
) MethodLike {
	var parser = Parser()
	return parser.ParseSource(source)
}

func ValidateMethod(
	method MethodLike,
) {
	var validator = Validator()
	validator.ValidateMethod(method)
}
