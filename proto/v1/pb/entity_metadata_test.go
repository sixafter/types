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
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateTags(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		tags  []string
		valid bool
	}{
		{"nil", nil, true},
		{"empty list", []string{}, true},
		{"valid", []string{"finance", "report2024", "a_b-C"}, true},
		{"max length", []string{strings.Repeat("a", MaxTagLength)}, true},
		{"case-sensitive distinct", []string{"Finance", "finance"}, true},
		{"empty tag", []string{"finance", ""}, false},
		{"too long", []string{strings.Repeat("a", MaxTagLength+1)}, false},
		{"space", []string{"annual report"}, false},
		{"dot", []string{"v1.0"}, false},
		{"non-ASCII", []string{"café"}, false},
		{"duplicate", []string{"finance", "report", "finance"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			err := ValidateTags(tt.tags)
			if tt.valid {
				is.NoError(err)
			} else {
				is.Error(err)
			}
		})
	}
}
