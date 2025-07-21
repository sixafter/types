// Copyright 2020-2024 SIX AFTER, INC (SIX AFTER)
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: SIX AFTER, INC (SIX AFTER)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUUIDHelpers_RoundTrip(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	orig := uuid.New()
	msg := UUIDToProto(orig)
	is.NotNil(msg, "UUIDToProto should not return nil")
	is.Equal(16, len(msg.Value), "protobuf UUID must be 16 bytes")

	// Validate
	is.NoError(ValidateUUID(msg), "ValidateUUID should succeed on correct UUID")

	// Convert back
	parsed, err := ProtoToUUID(msg)
	is.NoError(err, "ProtoToUUID should not error")
	is.Equal(orig, parsed, "Roundtrip UUIDs should match")

	// To string and back
	str, err := ProtoToString(msg)
	is.NoError(err, "ProtoToString should not error")
	is.Equal(orig.String(), str, "UUID string representation must match")

	msg2, err := StringToProto(str)
	is.NoError(err, "StringToProto should not error")
	is.Equal(16, len(msg2.Value), "protobuf UUID from string must be 16 bytes")
	is.Equal(msg.Value, msg2.Value, "Proto values from original and string must match")
}

func TestUUIDHelpers_InvalidLength(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	// Invalid: nil message
	is.Error(ValidateUUID(nil), "ValidateUUID should fail on nil input")
	_, err := ProtoToUUID(nil)
	is.Error(err, "ProtoToUUID should error on nil input")
	_, err = ProtoToString(nil)
	is.Error(err, "ProtoToString should error on nil input")

	// Invalid: wrong length
	msg := &UUID{Value: []byte{1, 2, 3}}
	is.Error(ValidateUUID(msg), "ValidateUUID should fail on wrong length")
	_, err = ProtoToUUID(msg)
	is.Error(err, "ProtoToUUID should error on wrong length")
	_, err = ProtoToString(msg)
	is.Error(err, "ProtoToString should error on wrong length")
}

func TestUUIDHelpers_StringParsing(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	validStr := "2b7e1516-28ae-4c98-8c3d-9b1636b6033c"
	msg, err := StringToProto(validStr)
	is.NoError(err, "StringToProto should not error on valid string")
	is.Equal(16, len(msg.Value), "StringToProto should produce 16 bytes")

	// Invalid: malformed string
	_, err = StringToProto("not-a-uuid")
	is.Error(err, "StringToProto should error on invalid string")
}

func TestUUIDHelpers_ProtoToStringAndBack(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	// Use a known UUID string to ensure roundtrip correctness.
	u, err := uuid.Parse("f81d4fae-7dec-11d0-a765-00a0c91e6bf6")
	is.NoError(err, "Parse should succeed on valid string")
	msg := UUIDToProto(u)

	// Proto -> String
	str, err := ProtoToString(msg)
	is.NoError(err, "ProtoToString should not error")
	is.Equal(u.String(), str, "String representations should match")

	// String -> Proto
	msg2, err := StringToProto(str)
	is.NoError(err, "StringToProto should not error")
	is.Equal(msg.Value, msg2.Value, "UUID proto values should match for roundtrip")
}

func TestUUIDHelpers_ValidateUUID(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	// Valid
	u := uuid.New()
	msg := UUIDToProto(u)
	is.NoError(ValidateUUID(msg), "should validate correct UUID")

	// Nil message
	is.Error(ValidateUUID(nil), "should error on nil UUID")

	// Invalid length (too short)
	short := &UUID{Value: []byte{1, 2, 3}}
	is.Error(ValidateUUID(short), "should error on short UUID")

	// Invalid length (too long)
	long := &UUID{Value: make([]byte, 20)}
	is.Error(ValidateUUID(long), "should error on long UUID")
}
