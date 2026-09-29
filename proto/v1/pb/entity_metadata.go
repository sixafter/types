// Copyright 2026 SIX AFTER, INC (SIX AFTER)
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
	"fmt"
)

// MaxTagLength is the maximum length, in characters, of a sixafter.types.proto.v1.EntityMetadata tag.
const MaxTagLength = 256

// ValidateTags checks that tags satisfy the sixafter.types.proto.v1.EntityMetadata tags constraints.
//
// Each tag must be 1 to MaxTagLength characters long and contain only ASCII
// letters, digits, underscores, and hyphens. Tags must be unique; comparison is
// case-sensitive, so "Finance" and "finance" are different tags.
//
// Returns nil if valid, or an error describing the first invalid tag otherwise.
// An empty or nil list is valid.
//
// Parameters:
//
//	tags - the tags to validate, typically msg.GetTags().
//
// Returns:
//
//	error if any tag is empty, too long, contains a disallowed character, or is a duplicate.
//
// Example:
//
//	if err := ValidateTags(msg.GetTags()); err != nil { ... }
func ValidateTags(tags []string) error {
	seen := make(map[string]struct{}, len(tags))
	for i, tag := range tags {
		if tag == "" {
			return fmt.Errorf("tag %d is empty", i)
		}
		if len(tag) > MaxTagLength {
			return fmt.Errorf("tag %d is %d characters, exceeds maximum of %d", i, len(tag), MaxTagLength)
		}
		for j := 0; j < len(tag); j++ {
			if !isTagByte(tag[j]) {
				return fmt.Errorf("tag %d %q contains disallowed character %q at position %d", i, tag, tag[j], j)
			}
		}
		if _, dup := seen[tag]; dup {
			return fmt.Errorf("tag %d %q is a duplicate", i, tag)
		}
		seen[tag] = struct{}{}
	}
	return nil
}

// isTagByte reports whether c matches the tag character class [a-zA-Z0-9_-].
func isTagByte(c byte) bool {
	return c >= 'a' && c <= 'z' ||
		c >= 'A' && c <= 'Z' ||
		c >= '0' && c <= '9' ||
		c == '_' || c == '-'
}
