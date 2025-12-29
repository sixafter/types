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

set -euo pipefail

__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${__dir}/os-type.sh"

if is_windows; then
  echo "[ERROR] Windows is not currently supported." >&2
  exit 1
fi

command -v buf >/dev/null 2>&1 || { echo "[ERROR] buf not found in PATH"; exit 1; }
[[ -f buf.yaml ]] || { echo "[ERROR] Run 'deps' make target."; exit 1; }

if [[ ! -f buf.gen.doc.yaml ]]; then
  echo "[INFO] buf.gen.doc.yaml not found; skipping generation."
  exit 0
fi

echo "[INFO] Running: proto generate docs"
buf generate --clean --template buf.gen.docs.yaml
