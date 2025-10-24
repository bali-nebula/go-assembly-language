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

package module_test

import (
	fmt "fmt"
	not "github.com/bali-nebula/go-assembly-language/v3"
	uti "github.com/craterdog/go-missing-utilities/v8"
	ass "github.com/stretchr/testify/assert"
	sts "strings"
	tes "testing"
)

const testDirectory = "./test/"

func TestParsingRoundtrips(t *tes.T) {
	var filenames = uti.ReadDirectory(testDirectory)
	for _, filename := range filenames {
		if sts.HasSuffix(filename, ".basm") {
			filename = testDirectory + filename
			fmt.Println(filename)
			var source = uti.ReadFile(filename)
			var assembly = not.ParseSource(source)
			not.ValidateAssembly(assembly)
			var formatted = not.FormatAssembly(assembly)
			ass.Equal(t, source, formatted)
		}
	}
}
