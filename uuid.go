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

/*
Package types provides helper functions for working with the sixafter.types.UUID
Protobuf message. This message represents a canonical 128-bit UUID (RFC 4122)
using a 16-byte binary encoding.

These helpers allow for conversion between the Protobuf UUID message and the
native github.com/google/uuid.UUID type, as well as conversion to and from the
canonical string representation of a UUID.

# Usage

	import (
	    "github.com/google/uuid"
	    "github.com/sixafter/types" // path to your generated and helper packages
	)

	// Generate a UUID and convert to protobuf
	u := uuid.New()
	pbUUID := types.UUIDToProto(u)

	// Convert back to uuid.UUID
	u2, err := types.ProtoToUUID(pbUUID)
	if err != nil {
	    // handle error
	}

	// Convert from string to protobuf
	pbUUID2, err := types.StringToProto("2b7e1516-28ae-4c98-8c3d-9b1636b6033c")
	if err != nil {
	    // handle error
	}

	// Convert protobuf to string
	s, err := types.ProtoToString(pbUUID2)
	if err != nil {
	    // handle error
	}

	// Validate protobuf message
	if err := types.ValidateUUID(pbUUID2); err != nil {
	    // handle error
	}

# Validation

The sixafter.types.UUID message MUST have a Value field of exactly 16 bytes.
These helpers perform validation and will return errors if the field is not valid.
*/
package types

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

// UUIDToProto converts a uuid.UUID to a sixafter.types.UUID Protobuf message.
//
// The returned UUID message Value field will be exactly 16 bytes, as required by the
// RFC 4122 standard and the sixafter.types.UUID message definition.
//
// Parameters:
//
//	u - uuid.UUID to be encoded as protobuf message.
//
// Returns:
//
//	Pointer to a new UUID protobuf message.
//
// Example:
//
//	u := uuid.New()
//	msg := UUIDToProto(u)
func UUIDToProto(u uuid.UUID) *UUID {
	return &UUID{Value: u[:]}
}

// ProtoToUUID decodes a sixafter.types.UUID Protobuf message to a uuid.UUID.
//
// Returns an error if the Value field is not exactly 16 bytes.
//
// Parameters:
//
//	msg - pointer to the UUID protobuf message.
//
// Returns:
//
//	uuid.UUID decoded from the protobuf message.
//	error if the Value field is nil or not exactly 16 bytes.
//
// Example:
//
//	u, err := ProtoToUUID(msg)
//	if err != nil { ... }
func ProtoToUUID(msg *UUID) (uuid.UUID, error) {
	if msg == nil {
		return uuid.Nil, fmt.Errorf("UUID message is nil")
	}
	if len(msg.Value) != 16 {
		return uuid.Nil, fmt.Errorf("invalid UUID length: got %d, want 16", len(msg.Value))
	}
	var arr [16]byte
	copy(arr[:], msg.Value)
	return uuid.UUID(arr), nil
}

// StringToProto parses a canonical UUID string and returns a sixafter.types.UUID Protobuf message.
//
// The input string must be in the canonical RFC 4122 format (e.g., "2b7e1516-28ae-4c98-8c3d-9b1636b6033c").
// Returns an error if the string is not a valid UUID.
//
// Parameters:
//
//	s - string representation of the UUID.
//
// Returns:
//
//	Pointer to a new UUID protobuf message.
//	error if the string is not a valid UUID.
//
// Example:
//
//	msg, err := StringToProto("2b7e1516-28ae-4c98-8c3d-9b1636b6033c")
func StringToProto(s string) (*UUID, error) {
	u, err := uuid.Parse(s)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID string: %w", err)
	}
	return UUIDToProto(u), nil
}

// ProtoToString returns the canonical RFC 4122 string representation of a sixafter.types.UUID Protobuf message.
//
// Returns an error if the Value field is not exactly 16 bytes.
//
// Parameters:
//
//	msg - pointer to the UUID protobuf message.
//
// Returns:
//
//	string representation of the UUID.
//	error if the Value field is nil or not exactly 16 bytes.
//
// Example:
//
//	s, err := ProtoToString(msg)
func ProtoToString(msg *UUID) (string, error) {
	u, err := ProtoToUUID(msg)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

// ValidateUUID checks that the sixafter.types.UUID protobuf message Value field is exactly 16 bytes.
//
// Returns nil if valid, or an error otherwise.
//
// Parameters:
//
//	msg - pointer to the UUID protobuf message.
//
// Returns:
//
//	error if Value is nil or not exactly 16 bytes.
//
// Example:
//
//	if err := ValidateUUID(msg); err != nil { ... }
func ValidateUUID(msg *UUID) error {
	if msg == nil {
		return fmt.Errorf("UUID message is nil")
	}
	if len(msg.Value) != 16 {
		return fmt.Errorf("UUID must be exactly 16 bytes, got %d", len(msg.Value))
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface for UUID.
// It outputs the UUID as a canonical RFC 4122 string (e.g. "550e8400-e29b-41d4-a716-446655440000")
// instead of the default base64 encoding that protojson uses for bytes fields.
func (m *UUID) MarshalJSON() ([]byte, error) {
	if m == nil || len(m.Value) != 16 {
		return []byte(`""`), nil
	}
	u, err := ProtoToUUID(m)
	if err != nil {
		return nil, err
	}
	return json.Marshal(u.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for UUID.
// It accepts a canonical RFC 4122 UUID string and converts it to binary form.
func (m *UUID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("invalid JSON for UUID: %w", err)
	}
	if s == "" {
		m.Value = nil
		return nil
	}
	u, err := uuid.Parse(s)
	if err != nil {
		return fmt.Errorf("invalid UUID string: %w", err)
	}
	m.Value = u[:]
	return nil
}

// Ensure interface conformance at compile time.
var (
	_ json.Marshaler   = (*UUID)(nil)
	_ json.Unmarshaler = (*UUID)(nil)
)
