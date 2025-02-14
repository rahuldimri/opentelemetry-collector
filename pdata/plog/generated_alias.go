// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "model/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "go run model/internal/cmd/pdatagen/main.go".

package plog

import "go.opentelemetry.io/collector/pdata/internal"

// ResourceLogsSlice logically represents a slice of ResourceLogs.
//
// This is a reference type. If passed by value and callee modifies it, the
// caller will see the modification.
//
// Must use NewResourceLogsSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ResourceLogsSlice = internal.ResourceLogsSlice

// NewResourceLogsSlice creates a ResourceLogsSlice with 0 elements.
// Can use "EnsureCapacity" to initialize with a given capacity.
var NewResourceLogsSlice = internal.NewResourceLogsSlice

// ResourceLogs is a collection of logs from a Resource.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewResourceLogs function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ResourceLogs = internal.ResourceLogs

// NewResourceLogs is an alias for a function to create a new empty ResourceLogs.
var NewResourceLogs = internal.NewResourceLogs

// ScopeLogsSlice logically represents a slice of ScopeLogs.
//
// This is a reference type. If passed by value and callee modifies it, the
// caller will see the modification.
//
// Must use NewScopeLogsSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ScopeLogsSlice = internal.ScopeLogsSlice

// NewScopeLogsSlice creates a ScopeLogsSlice with 0 elements.
// Can use "EnsureCapacity" to initialize with a given capacity.
var NewScopeLogsSlice = internal.NewScopeLogsSlice

// ScopeLogs is a collection of logs from a LibraryInstrumentation.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewScopeLogs function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ScopeLogs = internal.ScopeLogs

// NewScopeLogs is an alias for a function to create a new empty ScopeLogs.
var NewScopeLogs = internal.NewScopeLogs

// LogRecordSlice logically represents a slice of LogRecord.
//
// This is a reference type. If passed by value and callee modifies it, the
// caller will see the modification.
//
// Must use NewLogRecordSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type LogRecordSlice = internal.LogRecordSlice

// NewLogRecordSlice creates a LogRecordSlice with 0 elements.
// Can use "EnsureCapacity" to initialize with a given capacity.
var NewLogRecordSlice = internal.NewLogRecordSlice

// LogRecord are experimental implementation of OpenTelemetry Log Data Model.

// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewLogRecord function to create new instances.
// Important: zero-initialized instance is not valid for use.
type LogRecord = internal.LogRecord

// NewLogRecord is an alias for a function to create a new empty LogRecord.
var NewLogRecord = internal.NewLogRecord
