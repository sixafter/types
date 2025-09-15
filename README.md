# types: Common Protocol Buffer Types

This module is a set of common types expressed as [Google Protocol Buffers](https://developers.google.com/protocol-buffers/).

[![Go Report Card](https://goreportcard.com/badge/github.com/sixafter/types)](https://goreportcard.com/report/github.com/sixafter/types)
[![License: Apache 2.0](https://img.shields.io/badge/license-Apache%202.0-blue?style=flat-square)](LICENSE)
[![Go](https://img.shields.io/github/go-mod/go-version/sixafter/types)](https://img.shields.io/github/go-mod/go-version/sixafter/types)
[![Go Reference](https://pkg.go.dev/badge/github.com/sixafter/types.svg)](https://pkg.go.dev/github.com/sixafter/types)

---
## Status

### Build & Test

[![ci](https://github.com/sixafter/types/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/sixafter/types/actions/workflows/ci.yaml)
[![GitHub issues](https://img.shields.io/github/issues/sixafter/types)](https://img.shields.io/github/issues/sixafter/types)
![GitHub last commit](https://img.shields.io/github/last-commit/sixafter/types)

### Package and Deploy

[![Release](https://github.com/sixafter/types/workflows/release/badge.svg)](https://github.com/sixafter/types/actions)

---

## Verify with Cosign

[Cosign](https://github.com/sigstore/cosign) is used to sign releases for integrity verification.

To verify the integrity of the `types` source, run the following commands:

```sh
# Fetch the latest release tag from GitHub API (e.g., "v1.47.0")
TAG=$(curl -s https://api.github.com/repos/sixafter/types/releases/latest | jq -r .tag_name)

# Remove leading "v" for filenames (e.g., "v1.47.0" -> "1.47.0")
VERSION=${TAG#v}

# Verify the release tarball
cosign verify-blob \
  --key https://raw.githubusercontent.com/sixafter/types/main/cosign.pub \
  --signature nanoid-${VERSION}.tar.gz.sig \
  nanoid-${VERSION}.tar.gz

# Download checksums.txt and its signature from the latest release assets
curl -LO https://github.com/sixafter/types/releases/download/${TAG}/checksums.txt
curl -LO https://github.com/sixafter/types/releases/download/${TAG}/checksums.txt.sig

# Verify checksums.txt with cosign
cosign verify-blob \
  --key https://raw.githubusercontent.com/sixafter/types/main/cosign.pub \
  --signature checksums.txt.sig \
  checksums.txt
```

If valid, Cosign will output:

```shell
Verified OK
```

---
### Overview

To use this module, execute the following command:

```shell
go get -u github.com/sixafter/types
```

In your Go code, you can import the module as follows.

```go
package main

import (
    "github.com/sixafter/types"
)
```

The use of Go [vendoring](https://golang.org/ref/mod#vendoring) is [RECOMMENDED].

## Contributing

Contributions are welcome. See [CONTRIBUTING](CONTRIBUTING.md)

### License

This project is licensed under the [Apache 2.0 License](https://choosealicense.com/licenses/apache-2.0/). See [LICENSE](LICENSE) file.

[MUST]: https://datatracker.ietf.org/doc/html/rfc2119
[MUST NOT]: https://datatracker.ietf.org/doc/html/rfc2119
[SHOULD]: https://datatracker.ietf.org/doc/html/rfc2119
[SHOULD NOT]: https://datatracker.ietf.org/doc/html/rfc2119
[MAY]: https://datatracker.ietf.org/doc/html/rfc2119
[SHALL]: https://datatracker.ietf.org/doc/html/rfc2119
[SHALL NOT]: https://datatracker.ietf.org/doc/html/rfc2119
[REQUIRED]: https://datatracker.ietf.org/doc/html/rfc2119
[RECOMMENDED]: https://datatracker.ietf.org/doc/html/rfc2119
[NOT RECOMMENDED]: https://datatracker.ietf.org/doc/html/rfc2119
