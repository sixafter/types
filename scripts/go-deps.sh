#!/bin/bash
# Copyright 2020 SIX AFTER, LLC (SIX AFTER)
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

# Protobuf
go get -u google.golang.org/protobuf
go get -u google.golang.org/protobuf/proto
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

"${__dir}"/go-install.sh
