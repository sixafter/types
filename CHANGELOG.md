# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Date format: `YYYY-MM-DD`

---
## [Unreleased]

### Added
### Changed
### Deprecated
### Removed
### Fixed
### Security

---

## [1.57.2] - 2025-12-29

### Added
### Changed
- **defect:** Fix incorrect package names in .proto files.

### Deprecated
### Removed
### Fixed
### Security

---

## [1.57.1] - 2025-12-29

### Added
### Changed
- **defect:** Corrected import paths in `.proto` files.

### Deprecated
### Removed
### Fixed
### Security

package sixafter.types.proto.v1;
---

## [1.57.0] - 2025-12-29

### Added
### Changed
- **defect:** Added missing Go file to proto folder so `get get` works as expected.

### Deprecated
### Removed
### Fixed
### Security

---

## [1.56.0] - 2025-12-29

### Added
### Changed
- **risk:** Refactored structure to avoid naming collisions when vendored by consumers.
- **debt:** Buf is now used for linting and code generation.

### Deprecated
### Removed
### Fixed
### Security

---

## [1.55.0] - 2025-12-13

### Added
### Changed
- **debt:** Upgraded dependencies to their latest stable versions.

### Deprecated
### Removed
### Fixed
### Security

---

## [1.54.0] - 2025-11-23

### Added
### Changed
### Deprecated
### Removed
### Fixed
- **defect:** Corrected `go_package` option in all proto files to use full import path `github.com/sixafter/types/sixafter/types;types`, eliminating import cycles and undefined type errors in generated code. This fix removes the requirement for consumers to specify manual import path mappings (`M` flags) in their `buf.gen.yaml` or protoc configurations.
### Security

---

## [1.53.0] - 2025-11-23

### Added
### Changed
- **risk:** Restructured proto files to `sixafter/types` subdirectory to prevent naming collisions when vendored by consumers.

### Deprecated
### Removed
### Fixed
### Security

---

## [1.52.0] - 2025-11-20

### Added
- **risk:** Added `signature-verify` make target to verify latest release's digital signatures for the current GOOS and GOARCH combination.

### Changed
- **debt:** Upgraded dependencies to their latest stable versions.

### Deprecated
### Removed
### Fixed
- **defect:** Fixed `README.md` instructions for verifying module checksums.

### Security
- **risk:** Upgraded `golang.org/x/crypto` to `v0.45.0` to address vulnerabilities.

---
## [1.51.3] - 2025-11-07

### Added
### Changed
- **debt:** Upgraded dependencies to latest stable versions.

### Deprecated
### Removed
### Fixed
### Security
- **risk:** Upgraded to the new GoReleaser Sigstore bundles verification method to enhance supply chain security.

---
## [1.50.0] - 2025-10-15

### Added
- Add JSON `marshal/unmarshal` support for `UUID` with tests ensuring canonical string encoding, roundtrip accuracy, and error handling.

### Changed
### Deprecated
### Removed
### Fixed
### Security

---
## [1.49.0] - 2025-10-14

### Added
### Changed
- **debt:** Changed `entity_metadata.proto` to use `_at` suffix for timestamp fields to align with common conventions.

### Deprecated
### Removed
### Fixed
### Security

---
## [1.48.0] - 2025-10-06

### Added
### Changed
- **debt:** Upgraded dependencies to latest stable versions.

### Deprecated
### Removed
### Fixed
### Security

---
## [1.47.0] - 2025-09-15

### Added
- **debt:** Added digitally signed Software Bill of Materials (SBOM) to the release artifacts for enhanced supply chain security and transparency.

### Changed
- **debt:** Upgraded dependencies to latest stable versions.

### Deprecated
### Removed
### Fixed
### Security

---
## [1.46.0] - 2025-09-14

### Added
### Changed
- **debt:** Upgraded dependencies to latest stable versions.

### Deprecated
### Removed
### Fixed
### Security

---
## [1.45.0] - 2025-09-01

### Added
- **feature:** Added GoReleaser support.

### Changed
- **debt:** Upgraded dependencies to latest stable versions.

### Deprecated
### Removed
### Fixed
### Security

---
## [1.44.0] - 2025-08-24

### Added
- **risk:** Added linting to the CI pipeline.
### Changed
- **defect:** Fixed `entity_metadata.proto` to correctly use `google/protobuf/struct.proto`.
### Deprecated
### Removed
### Fixed
### Security

---
## [1.43.0] - 2025-08-21

### Added
**feature:** Added support for ISO 639 [Language](sixafter/types/language.proto) message, which includes BCP 47 language tags.

### Changed
- **debt:** Upgraded dependencies to their latest stable versions; e.g. protoc-gen-go `v1.36.8`.

### Deprecated
### Removed
### Fixed
### Security

---
## [1.42.0] - 2025-08-16

### Added
### Changed
- **debt:** Upgraded dependencies to their latest stable versions; e.g. protobuf.

### Deprecated
### Removed
### Fixed
### Security

---
## [1.41.0] - 2025-08-15

### Added
### Changed
- **debt:** Upgraded dependencies to their latest stable versions; e.g. protobuf.

### Deprecated
### Removed
### Fixed
### Security

---
## [1.40.0] - 2025-08-13

### Added
### Changed
- **debt:** Upgraded to Go 1.25 to take advantage of the latest features and improvements.

### Deprecated
### Removed
### Fixed
### Security

---
## [1.39.0] - 2025-08-08

### Added
### Changed
- **debt:** Upgraded dependencies to their latest stable versions.

### Deprecated
### Removed
### Fixed
### Security

---
## [1.38.0] - 2025-07-21

### Added
- **feature**: Added canonical UUID Protobuf message type (`sixafter.types.UUID`) with formal RFC 4122 documentation and 16-byte wire encoding. Provided Go helper functions for safe conversion, validation, and string handling, with full unit test coverage.

### Changed
### Deprecated
### Removed
### Fixed
### Security

---
## [1.37.0] - 2025-06-19

### Added
- **feature:* Added 'country' to the `country_subdivision.proto` file to represent the country to which the subdivision belongs.
- **feature:** Added `time_zone.central_coordinate` to the `time_zone.proto` file to represent the central coordinate of the time zone.

### Changed
- **debt:** The `help` make target now sorts the output.
- **debt:** Renamed `entity_metadata.extensions` to `entity_metadata.attributes` to avoid confusion with protocol buffer extensions.

### Deprecated
### Removed
### Fixed
### Security

---
## [1.36.0] - 2025-05-12

### Added
### Changed
### Deprecated
### Removed
### Fixed
- **defect:** Fixed bug in 'CHANGELOG.md' where the path was incorrect.
### Security

---
## [1.34.0] - 2025-05-07

### Added
### Changed
- **debt:** Upgraded all dependencies to the latest versions to ensure compatibility and security.

### Deprecated
### Removed
### Fixed
### Security

---
## [1.33.0] - 2025-03-25

### Added
### Changed
- **debt:** Upgraded all dependencies to the latest versions to ensure compatibility and security.

### Deprecated
### Removed
### Fixed
### Security

---
## [1.32.0] - 2025-02-13

### Added
### Changed
- **debt:** Upgraded to Go 1.24 to take advantage of the latest features and improvements.
- **debt:** Upgraded all dependencies to the latest versions to ensure compatibility and security.

### Deprecated
### Removed
### Fixed
### Security

---
## [1.31.0] - 2025-02-06

### Added
### Changed
- **debt:** Upgraded all dependencies to the latest versions to ensure compatibility and security.

### Deprecated
### Removed
### Fixed
### Security

---
## [1.30.0] - 2025-01-24

### Added
### Changed
- **debt:** Upgraded all dependencies to the latest versions to ensure compatibility and security.

### Deprecated
### Removed
### Fixed
### Security

---
## [1.29.0] - 2025-01-16

### Added
### Changed
- **debt:** Upgraded all dependencies to the latest versions to ensure compatibility and security.

### Deprecated
### Removed
### Fixed
### Security

---
## [1.28.0] - 2025-01-08

### Added
### Changed
- **debt:** Upgraded all dependencies to the latest versions to ensure compatibility and security.

### Deprecated
### Removed
### Fixed
### Security

---
## [1.27.0] - 2024-12-26

### Added
### Changed
- **debt:** Upgraded all dependencies to the latest versions to ensure compatibility and security.

### Deprecated
### Removed
### Fixed
### Security

---
## [1.24.0] - 2024-11-17

### Added
- **FEATURE:** Added [geometry.proto](sixafter/types/geometry.proto) for defining geometrical constructs, including:
  - Scalar: Represents a scalar number using an unscaled integer value and a scale for fixed-point arithmetic. 
  - CoordinateSystem: Defines the coordinate system used to interpret geometric entities (e.g., Cartesian, Polar), with optional parameters. 
  - GeometryType: Specifies the type of geometry abstracting the mathematical space (e.g., Euclidean, Hyperbolic), with optional parameters. 
  - Coordinate: Represents an n-dimensional coordinate within a specified geometry and coordinate system, using Scalar values for precision. 
  - Point: Defines a point in a mathematical space, characterized by its Coordinate. 
  - Line: Represents a line or geodesic in a mathematical space, defined by starting and ending Points. 
  - Polygon: Describes a polygon in a mathematical space, defined by a series of Point vertices and an is_closed flag.

### Changed
- **DEBT:** Renamed `polygon.proto` to `map_polygon.proto` to better reflect the purpose of the file.

### Deprecated
### Removed
### Fixed
### Security

---
## [1.23.0] - 2024-11-17

### Added
- **FEATURE:** Added detailed documentation to all `.proto` files to describe the purpose and expected use of each message.
- **FEATURE:** Added `csharp_namespace` option to the `protoc` command to specify the C# namespace for generated files.
- **FEATURE:** Added `map<string, string> extensions` to the `entity_metadata.proto` file to allow for custom metadata to be added to entities.

### Changed
- **DEBT:** Changed `option objc_class_prefix` to `TPB` to match the project name.

### Deprecated
### Removed
### Fixed
### Security

---
## [1.22.0] - 2024-11-17

### Added
- **FEATURE:** Added a [CHANGELOG](CHANGELOG.md) to track all notable changes to the project.

### Changed
### Deprecated
### Removed
### Fixed
### Security

[Unreleased]: https://github.com/sixafter/types/compare/v1.57.2...HEAD
[1.57.1]: https://github.com/sixafter/types/compare/v1.57.1...v1.57.2
[1.57.1]: https://github.com/sixafter/types/compare/v1.57.0...v1.57.1
[1.57.0]: https://github.com/sixafter/types/compare/v1.56.0...v1.57.0
[1.56.0]: https://github.com/sixafter/types/compare/v1.55.0...v1.56.0
[1.55.0]: https://github.com/sixafter/types/compare/v1.54.0...v1.55.0
[1.54.0]: https://github.com/sixafter/types/compare/v1.53.0...v1.54.0
[1.53.0]: https://github.com/sixafter/types/compare/v1.52.0...v1.53.0
[1.52.0]: https://github.com/sixafter/types/compare/v1.51.3...v1.52.0
[1.51.3]: https://github.com/sixafter/types/compare/v1.50.0...v1.51.3
[1.50.0]: https://github.com/sixafter/types/compare/v1.49.0...v1.50.0
[1.49.0]: https://github.com/sixafter/types/compare/v1.48.0...v1.49.0
[1.48.0]: https://github.com/sixafter/types/compare/v1.47.0...v1.48.0
[1.47.0]: https://github.com/sixafter/types/compare/v1.46.0...v1.47.0
[1.46.0]: https://github.com/sixafter/types/compare/v1.45.0...v1.46.0
[1.45.0]: https://github.com/sixafter/types/compare/v1.44.0...v1.45.0
[1.44.0]: https://github.com/sixafter/types/compare/v1.43.0...v1.44.0
[1.43.0]: https://github.com/sixafter/types/compare/v1.42.0...v1.43.0
[1.42.0]: https://github.com/sixafter/types/compare/v1.41.0...v1.42.0
[1.41.0]: https://github.com/sixafter/types/compare/v1.40.0...v1.41.0
[1.40.0]: https://github.com/sixafter/types/compare/v1.39.0...v1.40.0
[1.39.0]: https://github.com/sixafter/types/compare/v1.38.0...v1.39.0
[1.38.0]: https://github.com/sixafter/types/compare/v1.37.0...v1.38.0
[1.37.0]: https://github.com/sixafter/types/compare/v1.36.0...v1.37.0
[1.36.0]: https://github.com/sixafter/types/compare/v1.34.0...v1.36.0
[1.34.0]: https://github.com/sixafter/types/compare/v1.33.0...v1.34.0
[1.33.0]: https://github.com/sixafter/types/compare/v1.32.0...v1.33.0
[1.32.0]: https://github.com/sixafter/types/compare/v1.31.0...v1.32.0
[1.31.0]: https://github.com/sixafter/types/compare/v1.30.0...v1.31.0
[1.30.0]: https://github.com/sixafter/types/compare/v1.29.0...v1.30.0
[1.29.0]: https://github.com/sixafter/types/compare/v1.28.0...v1.29.0
[1.28.0]: https://github.com/sixafter/types/compare/v1.27.0...v1.28.0
[1.27.0]: https://github.com/sixafter/types/compare/v1.26.0...v1.27.0
[1.26.0]: https://github.com/sixafter/types/compare/v1.25.0...v1.26.0
[1.25.0]: https://github.com/sixafter/types/compare/v1.24.0...v1.25.0
[1.24.0]: https://github.com/sixafter/types/compare/v1.23.0...v1.24.0
[1.23.0]: https://github.com/sixafter/types/compare/v1.22.0...v1.23.0
[1.22.0]: https://github.com/sixafter/types/compare/v1.0.1...v1.22.0

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
