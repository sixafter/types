#!/bin/bash
# Copyright 2020-2025 SIX AFTER, INC (SIX AFTER)
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

source "${__dir}"/os-type.sh

# Windows
if is_windows; then
    echo "Windows is not currently supported."
    exit 1
fi

protoc --go_out=. --go_opt=paths=source_relative --proto_path=. ./*.proto
