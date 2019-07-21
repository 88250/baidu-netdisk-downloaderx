// Gulu - Golang common utilities for everyone.
// Copyright (c) 2019-present, b3log.org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package gulu implements some common utilities.
package gulu

import "os"

// Version is the version of Gulu.
const Version = "v1.0.0"

// Logger is the logger used in Gulu internally.
var logger = Log.NewLogger(os.Stdout)

type (
	// GuluFile is the receiver of file utilities
	GuluFile byte
	// GuluGo is the receiver of Go utilities
	GuluGo byte
	// GuluNet is the receiver of network utilities
	GuluNet byte
	// GuluOS is the receiver of OS utilities
	GuluOS byte
	// GuluPanic is the receiver of panic utilities
	GuluPanic byte
	// GuluRand is the receiver of random utilities
	GuluRand byte
	// GuluRet is the receiver of result utilities
	GuluRet byte
	// GuluRune is the receiver of rune utilities
	GuluRune byte
	// GuluStr is the receiver of string utilities
	GuluStr byte
	// GuluZip is the receiver of zip utilities
	GuluZip byte
)

var (
	// File utilities
	File GuluFile
	// Go utilities
	Go GuluGo
	// Net utilities
	Net GuluNet
	// OS utilities
	OS GuluOS
	// Panic utilities
	Panic GuluPanic
	// Rand utilities
	Rand GuluRand
	// Ret utilities
	Ret GuluRet
	// Rune utilities
	Rune GuluRune
	// Str utilities
	Str GuluStr
	// Zip utilities
	Zip GuluZip
)
