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

UNAME=$( command -v uname)

function detect_os() {
  OS=$( "${UNAME}" | tr '[:upper:]' '[:lower:]')
  case $OS in
    linux*)
      OS='Linux'
      ;;
    msys*|cygwin*|mingw*)
      # or possible 'bash on windows'
      OS='Windows'
      ;;
    nt|win*)
      OS='Windows'
      ;;
    darwin*)
      OS='macOS'
      ;;
    *) ;;
  esac

  echo $OS
}

function is_linux() {
  local OS=$(detect_os)
  if [[ $OS == 'Linux' ]]; then
    return $(true)
  fi

  return $(false)
}

function is_macos() {
  local OS=$(detect_os)
  if [[ $OS == 'macOS' ]]; then
    return $(true)
  fi

  return $(false)
}

function is_macos_arm() {
  local OS=$(detect_os)
  if [[ $OS == 'macOS' && $(uname -p) == 'arm64' ]]; then
    return $(true)
  fi

  return $(false)
}

function is_macos_amd() {
  local OS=$(detect_os)
  if [[ $OS == 'macOS' && $(uname -p) == 'i386' ]]; then
    return $(true)
  fi

  return $(false)
}

function is_windows() {
  local OS=$(detect_os)
  if [[ $OS == 'Windows' ]]; then
    return $(true)
  fi

  return $(false)
}

function goos() {
  case "$(detect_os)" in
    Linux)  echo "linux" ;;
    macOS)  echo "darwin" ;;
    Windows) echo "windows" ;;
    *) echo "unsupported"; return 1 ;;
  esac
}

function goarch() {
  local ARCH
  ARCH=$(uname -m)

  case "$ARCH" in
    x86_64|amd64)  echo "amd64" ;;
    arm64|aarch64) echo "arm64" ;;
    armv6l)        echo "armv6" ;;
    armv7l)        echo "armv7" ;;
    i386|i686)     echo "386" ;;
    *) echo "unsupported"; return 1 ;;
  esac
}
