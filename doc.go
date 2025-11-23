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

// Package types provides protocol buffer definitions for common, reusable types
// used across SIX AFTER services and applications.
//
// These types include geographic coordinates, temporal ranges, entity metadata,
// and other well-known types that can be imported and used in other protobuf
// definitions or Go code.
//
// To import these proto files in your own .proto files:
//
//	import "sixafter/types/uuid.proto";
//	import "sixafter/types/geospatial_coordinate.proto";
//
// To use the generated Go types:
//
//	import "github.com/sixafter/types/sixafter/types"
package types
