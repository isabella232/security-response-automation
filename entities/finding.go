// Package entities holds commonly used methods used in security automation.
package entities

// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
	"github.com/pkg/errors"
)

var (
	// ErrUnmarshal thrown when unable to unmarshal.
	ErrUnmarshal = errors.New("failed to unmarshal")
	// ErrParsing thrown when unable to parse.
	ErrParsing = errors.New("not a valid log")
	// ErrValueNotFound thrown when a value is requested but not found.
	ErrValueNotFound = errors.New("value not found")
)

// Life of a finding
//
// Findings are deserialized with the Cloud Function that uses them. For example, the
// CreateSnapshot Cloud Function will attempt to deserialize ETD bad_ip findings. In an effort to
// reuse methods where possible (and without inheritence) we have a ETD base finding that all ETD
// findings will satisfy then a separate structure for each finding. Since Go does not support
// inheritence and we do not currenty have a published schema for each finding this means we
// deserialize several times. `NewBadIP` will attempt to deserialize it as a bad_ip finding as well
// as a base ETD finding.
//
// In order to get common values from findings we have helper methods associated with the base struct
// `Finding` and on the specific types such as `BadIP`. In an effort to make deferencing these values
// easily we wrap in a method without error checking. To prevent accesor failures each finding will
// implement a `validate` method that ensure all 'getter' method calls will succeed.

// StackDriverLog struct fits StackDriver logs.
type StackDriverLog struct {
	InsertID string `json:"insertId"`
	LogName  string `json:"logName"`
}

// badNetworkFinding contains any finding based off VPC flow logs.
type badNetworkFinding struct {
	JSONPayload struct {
		Properties struct {
			Location       string
			SourceInstance string
			IP             []string
		}
	}
}