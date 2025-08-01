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

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func ParameterizedClass() ParameterizedClassLike {
	return parameterizedClass()
}

// Constructor Methods

func (c *parameterizedClass_) Parameterized(
	delimiter string,
) ParameterizedLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	var instance = &parameterized_{
		// Initialize the instance attributes.
		delimiter_: delimiter,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *parameterized_) GetClass() ParameterizedClassLike {
	return parameterizedClass()
}

// Attribute Methods

func (v *parameterized_) GetDelimiter() string {
	return v.delimiter_
}

// PROTECTED INTERFACE

// Instance Structure

type parameterized_ struct {
	// Declare the instance attributes.
	delimiter_ string
}

// Class Structure

type parameterizedClass_ struct {
	// Declare the class constants.
}

// Class Reference

func parameterizedClass() *parameterizedClass_ {
	return parameterizedClassReference_
}

var parameterizedClassReference_ = &parameterizedClass_{
	// Initialize the class constants.
}
