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
command -v git >/dev/null 2>&1 || { echo "[ERROR] git not found in PATH"; exit 1; }

# Allow running from service root (has buf.yaml v2) or module dir
if [[ ! -f buf.yaml && ! -f buf.work.yaml ]]; then
  echo "[ERROR] Neither buf.yaml nor buf.work.yaml found in: $(pwd)"
  exit 1
fi

# Resolve repo root and service subdir
GIT_ROOT="$(git rev-parse --show-toplevel)"
REL_SUBDIR="$(git rev-parse --show-prefix)"; REL_SUBDIR="${REL_SUBDIR%/}"

BRANCH="${BUF_BREAK_BRANCH:-main}"
AGAINST="${BUF_BREAK_AGAINST:-${GIT_ROOT}/.git#branch=${BRANCH},subdir=${REL_SUBDIR}}"

# Strip accidental quotes
AGAINST="${AGAINST%\"}"; AGAINST="${AGAINST#\"}"
AGAINST="${AGAINST%\'}"; AGAINST="${AGAINST#\'}"

if [[ -n "${BUF_BREAK_PATHS:-}" ]]; then
  echo "[INFO] Running: proto breaking --against ${AGAINST} --path ${BUF_BREAK_PATHS}"
  buf breaking --against "${AGAINST}" --path "${BUF_BREAK_PATHS}"
else
  echo "[INFO] Running: proto breaking --against ${AGAINST}"
  buf breaking --against "${AGAINST}"
fi
