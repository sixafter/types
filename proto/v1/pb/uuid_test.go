// Copyright 2020-2025 SIX AFTER, INC (SIX AFTER)
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
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUUIDHelpers_RoundTrip(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	orig := uuid.New()
	msg := UUIDToProto(orig)
	is.NotNil(msg)
	is.Equal(16, len(msg.Value))
	is.NoError(ValidateUUID(msg))

	parsed, err := ProtoToUUID(msg)
	is.NoError(err)
	is.Equal(orig, parsed)

	str, err := ProtoToString(msg)
	is.NoError(err)
	is.Equal(orig.String(), str)

	msg2, err := StringToProto(str)
	is.NoError(err)
	is.Equal(msg.Value, msg2.Value)
}

func TestUUIDHelpers_InvalidLength(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Error(ValidateUUID(nil))
	_, err := ProtoToUUID(nil)
	is.Error(err)
	_, err = ProtoToString(nil)
	is.Error(err)

	msg := &UUID{Value: []byte{1, 2, 3}}
	is.Error(ValidateUUID(msg))
	_, err = ProtoToUUID(msg)
	is.Error(err)
	_, err = ProtoToString(msg)
	is.Error(err)
}

func TestUUIDHelpers_StringParsing(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	validStr := "2b7e1516-28ae-4c98-8c3d-9b1636b6033c"
	msg, err := StringToProto(validStr)
	is.NoError(err)
	is.Equal(16, len(msg.Value))

	_, err = StringToProto("not-a-uuid")
	is.Error(err)
}

func TestUUIDHelpers_ProtoToStringAndBack(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	u, err := uuid.Parse("f81d4fae-7dec-11d0-a765-00a0c91e6bf6")
	is.NoError(err)
	msg := UUIDToProto(u)

	str, err := ProtoToString(msg)
	is.NoError(err)
	is.Equal(u.String(), str)

	msg2, err := StringToProto(str)
	is.NoError(err)
	is.Equal(msg.Value, msg2.Value)
}

func TestUUIDHelpers_ValidateUUID(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	u := uuid.New()
	msg := UUIDToProto(u)
	is.NoError(ValidateUUID(msg))

	is.Error(ValidateUUID(nil))
	short := &UUID{Value: []byte{1, 2, 3}}
	is.Error(ValidateUUID(short))
	long := &UUID{Value: make([]byte, 20)}
	is.Error(ValidateUUID(long))
}

// --- JSON Tests ---

func TestUUID_JSON_MarshalUnmarshal(t *testing.T) {
	is := assert.New(t)

	u := uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")
	msg := UUIDToProto(u)

	// Marshal to JSON
	b, err := json.Marshal(msg)
	is.NoError(err)
	is.Contains(string(b), u.String())

	// Unmarshal back
	var decoded UUID
	err = json.Unmarshal(b, &decoded)
	is.NoError(err)
	is.Equal(msg.Value, decoded.Value)
}

func TestUUID_JSON_EmptyAndInvalid(t *testing.T) {
	is := assert.New(t)

	// Empty string
	var m UUID
	err := json.Unmarshal([]byte(`""`), &m)
	is.NoError(err)
	is.Nil(m.Value)

	// Invalid JSON
	err = json.Unmarshal([]byte(`123`), &m)
	is.Error(err)

	// Invalid UUID string
	err = json.Unmarshal([]byte(`"not-a-uuid"`), &m)
	is.Error(err)
}

func TestUUID_JSON_InterfaceConformance(t *testing.T) {
	var _ json.Marshaler = (*UUID)(nil)
	var _ json.Unmarshaler = (*UUID)(nil)
}
