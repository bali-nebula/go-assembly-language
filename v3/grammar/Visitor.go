/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies.  All Rights Reserved.    .
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
│              This class file was automatically generated using:              │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package grammar

import (
	ast "github.com/bali-nebula/go-assembly-language/v3/ast"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func VisitorClass() VisitorClassLike {
	return visitorClass()
}

// Constructor Methods

func (c *visitorClass_) Visitor(
	processor Methodical,
) VisitorLike {
	if uti.IsUndefined(processor) {
		panic("The \"processor\" attribute is required by this class.")
	}
	var instance = &visitor_{
		// Initialize the instance attributes.
		processor_: processor,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *visitor_) GetClass() VisitorClassLike {
	return visitorClass()
}

func (v *visitor_) VisitAssembly(
	assembly ast.AssemblyLike,
) {
	v.processor_.PreprocessAssembly(
		assembly,
		1,
		1,
	)
	v.visitAssembly(assembly)
	v.processor_.PostprocessAssembly(
		assembly,
		1,
		1,
	)
}

// PROTECTED INTERFACE

// Private Methods

func (v *visitor_) visitAction(
	action ast.ActionLike,
) {
	// Visit the possible action rule types.
	switch actual := action.GetAny().(type) {
	case ast.NoteLike:
		v.processor_.PreprocessNote(
			actual,
			1,
			1,
		)
		v.visitNote(actual)
		v.processor_.PostprocessNote(
			actual,
			1,
			1,
		)
	case ast.SkipLike:
		v.processor_.PreprocessSkip(
			actual,
			1,
			1,
		)
		v.visitSkip(actual)
		v.processor_.PostprocessSkip(
			actual,
			1,
			1,
		)
	case ast.JumpLike:
		v.processor_.PreprocessJump(
			actual,
			1,
			1,
		)
		v.visitJump(actual)
		v.processor_.PostprocessJump(
			actual,
			1,
			1,
		)
	case ast.PushLike:
		v.processor_.PreprocessPush(
			actual,
			1,
			1,
		)
		v.visitPush(actual)
		v.processor_.PostprocessPush(
			actual,
			1,
			1,
		)
	case ast.PullLike:
		v.processor_.PreprocessPull(
			actual,
			1,
			1,
		)
		v.visitPull(actual)
		v.processor_.PostprocessPull(
			actual,
			1,
			1,
		)
	case ast.LoadLike:
		v.processor_.PreprocessLoad(
			actual,
			1,
			1,
		)
		v.visitLoad(actual)
		v.processor_.PostprocessLoad(
			actual,
			1,
			1,
		)
	case ast.SaveLike:
		v.processor_.PreprocessSave(
			actual,
			1,
			1,
		)
		v.visitSave(actual)
		v.processor_.PostprocessSave(
			actual,
			1,
			1,
		)
	case ast.DropLike:
		v.processor_.PreprocessDrop(
			actual,
			1,
			1,
		)
		v.visitDrop(actual)
		v.processor_.PostprocessDrop(
			actual,
			1,
			1,
		)
	case ast.CallLike:
		v.processor_.PreprocessCall(
			actual,
			1,
			1,
		)
		v.visitCall(actual)
		v.processor_.PostprocessCall(
			actual,
			1,
			1,
		)
	case ast.SendLike:
		v.processor_.PreprocessSend(
			actual,
			1,
			1,
		)
		v.visitSend(actual)
		v.processor_.PostprocessSend(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitArgument(
	argument ast.ArgumentLike,
) {
	var delimiter = argument.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessArgumentSlot(
		argument,
		1,
	)

	var symbol = argument.GetSymbol()
	v.processor_.ProcessSymbol(symbol)
}

func (v *visitor_) visitAssembly(
	assembly ast.AssemblyLike,
) {
	var instructionsIndex uint
	var instructions = assembly.GetInstructions().GetIterator()
	var instructionsCount = uint(instructions.GetSize())
	for instructions.HasNext() {
		instructionsIndex++
		var rule = instructions.GetNext()
		v.processor_.PreprocessInstruction(
			rule,
			instructionsIndex,
			instructionsCount,
		)
		v.visitInstruction(rule)
		v.processor_.PostprocessInstruction(
			rule,
			instructionsIndex,
			instructionsCount,
		)
	}
}

func (v *visitor_) visitCall(
	call ast.CallLike,
) {
	var delimiter = call.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessCallSlot(
		call,
		1,
	)

	var symbol = call.GetSymbol()
	v.processor_.ProcessSymbol(symbol)
	// Visit slot 2 between terms.
	v.processor_.ProcessCallSlot(
		call,
		2,
	)

	var optionalCardinality = call.GetOptionalCardinality()
	if uti.IsDefined(optionalCardinality) {
		v.processor_.PreprocessCardinality(
			optionalCardinality,
			1,
			1,
		)
		v.visitCardinality(optionalCardinality)
		v.processor_.PostprocessCardinality(
			optionalCardinality,
			1,
			1,
		)
	}
}

func (v *visitor_) visitCardinality(
	cardinality ast.CardinalityLike,
) {
	// Visit the possible cardinality literal values.
	var actual = cardinality.GetAny().(string)
	switch actual {
	case "WITH 1 ARGUMENT":
		v.processor_.ProcessDelimiter("WITH 1 ARGUMENT")
	case "WITH 2 ARGUMENTS":
		v.processor_.ProcessDelimiter("WITH 2 ARGUMENTS")
	case "WITH 3 ARGUMENTS":
		v.processor_.ProcessDelimiter("WITH 3 ARGUMENTS")
	}
}

func (v *visitor_) visitComponent(
	component ast.ComponentLike,
) {
	// Visit the possible component literal values.
	var actual = component.GetAny().(string)
	switch actual {
	case "CONTRACT":
		v.processor_.ProcessDelimiter("CONTRACT")
	case "DRAFT":
		v.processor_.ProcessDelimiter("DRAFT")
	case "VARIABLE":
		v.processor_.ProcessDelimiter("VARIABLE")
	case "MESSAGE":
		v.processor_.ProcessDelimiter("MESSAGE")
	}
}

func (v *visitor_) visitConditionally(
	conditionally ast.ConditionallyLike,
) {
	// Visit the possible conditionally literal values.
	var actual = conditionally.GetAny().(string)
	switch actual {
	case "ON EMPTY":
		v.processor_.ProcessDelimiter("ON EMPTY")
	case "ON NONE":
		v.processor_.ProcessDelimiter("ON NONE")
	case "ON FALSE":
		v.processor_.ProcessDelimiter("ON FALSE")
	}
}

func (v *visitor_) visitConstant(
	constant ast.ConstantLike,
) {
	var delimiter = constant.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessConstantSlot(
		constant,
		1,
	)

	var symbol = constant.GetSymbol()
	v.processor_.ProcessSymbol(symbol)
}

func (v *visitor_) visitDestination(
	destination ast.DestinationLike,
) {
	// Visit the possible destination literal values.
	var actual = destination.GetAny().(string)
	switch actual {
	case "CONTRACT":
		v.processor_.ProcessDelimiter("CONTRACT")
	case "COMPONENT":
		v.processor_.ProcessDelimiter("COMPONENT")
	}
}

func (v *visitor_) visitDrop(
	drop ast.DropLike,
) {
	var delimiter = drop.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessDropSlot(
		drop,
		1,
	)

	var component = drop.GetComponent()
	v.processor_.PreprocessComponent(
		component,
		1,
		1,
	)
	v.visitComponent(component)
	v.processor_.PostprocessComponent(
		component,
		1,
		1,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessDropSlot(
		drop,
		2,
	)

	var symbol = drop.GetSymbol()
	v.processor_.ProcessSymbol(symbol)
}

func (v *visitor_) visitHandler(
	handler ast.HandlerLike,
) {
	var delimiter = handler.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessHandlerSlot(
		handler,
		1,
	)

	var label = handler.GetLabel()
	v.processor_.ProcessLabel(label)
}

func (v *visitor_) visitInstruction(
	instruction ast.InstructionLike,
) {
	var optionalPrefix = instruction.GetOptionalPrefix()
	if uti.IsDefined(optionalPrefix) {
		v.processor_.PreprocessPrefix(
			optionalPrefix,
			1,
			1,
		)
		v.visitPrefix(optionalPrefix)
		v.processor_.PostprocessPrefix(
			optionalPrefix,
			1,
			1,
		)
	}
	// Visit slot 1 between terms.
	v.processor_.ProcessInstructionSlot(
		instruction,
		1,
	)

	var action = instruction.GetAction()
	v.processor_.PreprocessAction(
		action,
		1,
		1,
	)
	v.visitAction(action)
	v.processor_.PostprocessAction(
		action,
		1,
		1,
	)
}

func (v *visitor_) visitJump(
	jump ast.JumpLike,
) {
	var delimiter = jump.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessJumpSlot(
		jump,
		1,
	)

	var label = jump.GetLabel()
	v.processor_.ProcessLabel(label)
	// Visit slot 2 between terms.
	v.processor_.ProcessJumpSlot(
		jump,
		2,
	)

	var optionalConditionally = jump.GetOptionalConditionally()
	if uti.IsDefined(optionalConditionally) {
		v.processor_.PreprocessConditionally(
			optionalConditionally,
			1,
			1,
		)
		v.visitConditionally(optionalConditionally)
		v.processor_.PostprocessConditionally(
			optionalConditionally,
			1,
			1,
		)
	}
}

func (v *visitor_) visitLiteral(
	literal ast.LiteralLike,
) {
	var delimiter = literal.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessLiteralSlot(
		literal,
		1,
	)

	var quoted = literal.GetQuoted()
	v.processor_.ProcessQuoted(quoted)
}

func (v *visitor_) visitLoad(
	load ast.LoadLike,
) {
	var delimiter = load.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessLoadSlot(
		load,
		1,
	)

	var component = load.GetComponent()
	v.processor_.PreprocessComponent(
		component,
		1,
		1,
	)
	v.visitComponent(component)
	v.processor_.PostprocessComponent(
		component,
		1,
		1,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessLoadSlot(
		load,
		2,
	)

	var symbol = load.GetSymbol()
	v.processor_.ProcessSymbol(symbol)
}

func (v *visitor_) visitNote(
	note ast.NoteLike,
) {
	var delimiter = note.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessNoteSlot(
		note,
		1,
	)

	var comment = note.GetComment()
	v.processor_.ProcessComment(comment)
}

func (v *visitor_) visitParameterized(
	parameterized ast.ParameterizedLike,
) {
	var delimiter = parameterized.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
}

func (v *visitor_) visitPrefix(
	prefix ast.PrefixLike,
) {
	var label = prefix.GetLabel()
	v.processor_.ProcessLabel(label)
	// Visit slot 1 between terms.
	v.processor_.ProcessPrefixSlot(
		prefix,
		1,
	)

	var delimiter = prefix.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
}

func (v *visitor_) visitPull(
	pull ast.PullLike,
) {
	var delimiter = pull.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessPullSlot(
		pull,
		1,
	)

	var value = pull.GetValue()
	v.processor_.PreprocessValue(
		value,
		1,
		1,
	)
	v.visitValue(value)
	v.processor_.PostprocessValue(
		value,
		1,
		1,
	)
}

func (v *visitor_) visitPush(
	push ast.PushLike,
) {
	var delimiter = push.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessPushSlot(
		push,
		1,
	)

	var source = push.GetSource()
	v.processor_.PreprocessSource(
		source,
		1,
		1,
	)
	v.visitSource(source)
	v.processor_.PostprocessSource(
		source,
		1,
		1,
	)
}

func (v *visitor_) visitSave(
	save ast.SaveLike,
) {
	var delimiter = save.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessSaveSlot(
		save,
		1,
	)

	var component = save.GetComponent()
	v.processor_.PreprocessComponent(
		component,
		1,
		1,
	)
	v.visitComponent(component)
	v.processor_.PostprocessComponent(
		component,
		1,
		1,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessSaveSlot(
		save,
		2,
	)

	var symbol = save.GetSymbol()
	v.processor_.ProcessSymbol(symbol)
}

func (v *visitor_) visitSend(
	send ast.SendLike,
) {
	var delimiter1 = send.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessSendSlot(
		send,
		1,
	)

	var symbol = send.GetSymbol()
	v.processor_.ProcessSymbol(symbol)
	// Visit slot 2 between terms.
	v.processor_.ProcessSendSlot(
		send,
		2,
	)

	var delimiter2 = send.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessSendSlot(
		send,
		3,
	)

	var destination = send.GetDestination()
	v.processor_.PreprocessDestination(
		destination,
		1,
		1,
	)
	v.visitDestination(destination)
	v.processor_.PostprocessDestination(
		destination,
		1,
		1,
	)
	// Visit slot 4 between terms.
	v.processor_.ProcessSendSlot(
		send,
		4,
	)

	var optionalParameterized = send.GetOptionalParameterized()
	if uti.IsDefined(optionalParameterized) {
		v.processor_.PreprocessParameterized(
			optionalParameterized,
			1,
			1,
		)
		v.visitParameterized(optionalParameterized)
		v.processor_.PostprocessParameterized(
			optionalParameterized,
			1,
			1,
		)
	}
}

func (v *visitor_) visitSkip(
	skip ast.SkipLike,
) {
	var delimiter = skip.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
}

func (v *visitor_) visitSource(
	source ast.SourceLike,
) {
	// Visit the possible source rule types.
	switch actual := source.GetAny().(type) {
	case ast.HandlerLike:
		v.processor_.PreprocessHandler(
			actual,
			1,
			1,
		)
		v.visitHandler(actual)
		v.processor_.PostprocessHandler(
			actual,
			1,
			1,
		)
	case ast.LiteralLike:
		v.processor_.PreprocessLiteral(
			actual,
			1,
			1,
		)
		v.visitLiteral(actual)
		v.processor_.PostprocessLiteral(
			actual,
			1,
			1,
		)
	case ast.ConstantLike:
		v.processor_.PreprocessConstant(
			actual,
			1,
			1,
		)
		v.visitConstant(actual)
		v.processor_.PostprocessConstant(
			actual,
			1,
			1,
		)
	case ast.ArgumentLike:
		v.processor_.PreprocessArgument(
			actual,
			1,
			1,
		)
		v.visitArgument(actual)
		v.processor_.PostprocessArgument(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitValue(
	value ast.ValueLike,
) {
	// Visit the possible value literal values.
	var actual = value.GetAny().(string)
	switch actual {
	case "HANDLER":
		v.processor_.ProcessDelimiter("HANDLER")
	case "EXCEPTION":
		v.processor_.ProcessDelimiter("EXCEPTION")
	case "COMPONENT":
		v.processor_.ProcessDelimiter("COMPONENT")
	case "RESULT":
		v.processor_.ProcessDelimiter("RESULT")
	}
}

// Instance Structure

type visitor_ struct {
	// Declare the instance attributes.
	processor_ Methodical
}

// Class Structure

type visitorClass_ struct {
	// Declare the class constants.
}

// Class Reference

func visitorClass() *visitorClass_ {
	return visitorClassReference_
}

var visitorClassReference_ = &visitorClass_{
	// Initialize the class constants.
}
