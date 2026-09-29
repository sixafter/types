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


## [1.61.0] - 2026-09-29

### Added
- **feature:** Added `archived_at` field (11) to `EntityMetadata`.
- **feature:** Added `ValidateTags` and `MaxTagLength` Go helpers for `EntityMetadata.tags`.
- **feature:** Added `path` field (7) to `Uri`.
- **feature:** Added `Space` message and `AngleUnit` enum to `geometry.proto`; every shape carries one `Space`.
- **feature:** Added `LineSegment` and `Polyline` messages to `geometry.proto`.

### Changed
- **feature:** `ValidateUUID` now requires an RFC 9562-compliant value (Nil, Max, or the RFC 9562 variant with version 1–8), not just 16 bytes. Conversion helpers still check only length.
- **debt:** Renumbered `Uri.user_info` from 7 to 6. **Breaking (wire):** `Uri` messages serialized by earlier versions decode incorrectly.
- **debt:** Changed `Uri.port` from `int32` to `uint32`; 0 means no port was given. **Breaking (API).**
- **debt:** Changed `GeometryType` and `CoordinateSystem` from free-form string messages to enums. **Breaking.**
- **debt:** Changed `Coordinate` to carry only `values`; `Point` and `Polygon` now carry a `Space` and `Coordinate`s. `Polygon` is always closed, with counter-clockwise vertices in 2D. **Breaking.**
- **defect:** Changed `MapPoint.x` and `MapPoint.y` from `float` to `double`; `float` limited Web Mercator positions to 0.5–2 m resolution. **Breaking (wire):** values serialized by earlier versions are dropped as unknown fields.
- **defect:** Reshaped `TimeZone` to describe the zone in effect at an instant: `name` (1), `abbreviation` (2), `utc_offset` (3, `google.protobuf.Duration` with a single ISO 8601 sign), `daylight_saving` (4), and `central_coordinate` (5). Removed `utc_offset_std`, `utc_offset_dst`, and `TimeZone.TimeOffset`. **Breaking (wire):** an old `utc_offset_std` decodes as `utc_offset` with the wrong value (hours read as seconds, minutes as nanoseconds); an old `utc_offset_dst` is dropped as an unknown field.
- **debt:** Renamed repeated fields to plural names: `Language.bcp47_tag` to `bcp47_tags` and `Geofence.polygon` to `polygons`. Field numbers are unchanged, so the binary wire format is compatible. **Breaking (JSON and API).**
- **debt:** Changed `csharp_namespace` from `SixAfter.Types.V1.WellKnownTypes` to `SixAfter.Types.V1` in all `.proto` files; `WellKnownTypes` is Google's namespace for protobuf's built-in types. **Breaking (C# API).**

### Deprecated
### Removed
- **debt:** Removed `option cc_enable_arenas = true` from all `.proto` files; arenas have been enabled by default since protobuf 3.14.
- **debt:** Removed the `google.api.field_info` `UUID4` annotation from `UUID.value`; it applied only to version 4 and to string fields. Any RFC 9562 version is now permitted.
- **debt:** Removed the `buf.build/googleapis/googleapis` and `google.golang.org/genproto/googleapis/api` dependencies.
- **debt:** Removed `authority` field from `Uri`; it duplicated `user_info`, `host`, and `port`. **Breaking.**
- **debt:** Removed `Url` message (`url.proto`); use `Uri`, which covers every URL component. **Breaking.**
- **debt:** Removed `Line` (use `LineSegment`) and `Polygon.is_closed` (use `Polyline` for open shapes) from `geometry.proto`. **Breaking.**

### Fixed
- **defect:** Fixed `make proto-format`, which failed because it ran `buf format` from `proto/v1` instead of the module root. Added `make proto-format CHECK=true`, which reports formatting differences and fails instead of rewriting files.
- **defect:** Fixed `sbin/proto-docs.sh`, which checked for `buf.gen.doc.yaml` instead of `buf.gen.docs.yaml` and so never generated documentation. Generated docs under `docs/sdk/_generated/` are now ignored by git.
- **defect:** Corrected documentation for `CompassHeading`, `GeospatialLocation`, `Version` examples (replaced non-ASCII dashes with the SemVer spec's hyphens), `Country.numeric_code` (leading zeros), and `EmailAddress` (split at the last "@").
- **defect:** Clarified that `EntityMetadata.version` versions released content under SemVer 2.0.0 and is not a revision counter for optimistic concurrency.
- **defect:** Clarified `EntityMetadata.tags` constraints: tags MUST be unique and are case-sensitive.
- **defect:** Corrected `EntityMetadata.attributes` documentation: values may be any JSON value, not only strings, as `google.protobuf.Struct` allows.
- **defect:** Corrected `TimeZone.central_coordinate` documentation: it is the principal city from the IANA `zone1970.tab`, not a geometric center, and does not by itself determine a location's time zone.
- **debt:** Applied `buf format` to `geospatial_location.proto` (import order).
- **defect:** Documented that a negative `CompassHeading.true_heading`, `CompassHeading.heading_accuracy`, `GeospatialLocation.course`, or `GeospatialLocation.speed` means invalid or unknown, matching Core Location.
- **defect:** Documented that an unset `GeospatialCoordinate.elevation`, `TimeZone.utc_offset`, `TimeZone.central_coordinate`, `GeospatialLocation.coordinate`, or `MapPoint.coordinate` means unknown.
- **defect:** Updated UUID helper doc comments to cite RFC 9562, the Go standard library `uuid` package, and the full message name `sixafter.types.proto.v1.UUID`.
- **defect:** `UUID.MarshalJSON` now returns an error for a non-empty `Value` that is not 16 bytes instead of silently emitting `""`; doc comments on `MarshalJSON`/`UnmarshalJSON` now state they apply to `encoding/json` only, not protojson.
- **defect:** Updated `UUID` to cite RFC 9562, which obsoletes RFC 4122, and to require an RFC 9562-compliant value.
- **defect:** Updated `Uri` to cite RFC 3986, which obsoletes RFC 2396 and RFC 2732.

### Security

---

## [1.60.2] - 2026-09-28

### Added
### Changed
### Deprecated
### Removed
### Fixed
- **risk:** Modified copyright to reflect current year.

### Security

---

## [1.60.1] - 2026-08-22

### Added
### Changed
### Deprecated
- **debt:** Deprecated use of `github.com/google/uuid` in favor of Go 1.27 intrinsic [uuid](https://github.com/golang/go/tree/master/src/uuid) support.

### Removed
### Fixed
### Security

---

## [1.60.0] - 2026-08-22

### Added
### Changed
- **debt:** Upgraded to [Go 1.27](https://go.dev/doc/go1.27).
- **debt:** Upgraded dependencies to their latest stable versions.

### Deprecated
### Removed
### Fixed
### Security

---

## [1.59.3] - 2026-08-13

### Added
### Changed
- **debt:** Upgraded dependencies to their latest stable versions.

### Deprecated
### Removed
### Fixed
### Security

---

## [1.59.2] - 2026-06-29

### Added
### Changed
- **debt:** Upgraded dependencies to their latest stable versions.

### Deprecated
### Removed
### Fixed
### Security

---

## [1.59.1] - 2026-04-08

### Added
### Changed
- **debt:** Upgraded dependencies to their latest stable versions.

### Deprecated
### Removed
### Fixed
### Security

---

## [1.59.0] - 2026-02-12

### Added
### Changed
- **debt:** Upgraded dependencies to their latest stable versions including module to Go 1.26.

### Deprecated
### Removed
### Fixed
### Security

---

## [1.58.0] - 2026-01-05

### Added
- **feature:** Added `google.api.field_info` format annotation to UUID value field with UUID4 format specification per [AIP-202](https://google.aip.dev/202), enabling machine-readable format validation and enhancing API documentation. See [protocolbuffers/protobuf#2224](https://github.com/protocolbuffers/protobuf/issues/2224#issuecomment-3711636308) for more details.

### Changed
### Deprecated
### Removed
### Fixed
### Security

---

## [1.57.4] - 2025-12-29

### Added
### Changed
- **defect:** Removed the duplicate (nested) proto directory.

### Deprecated
### Removed
### Fixed
### Security

---

## [1.57.3] - 2025-12-29

### Added
### Changed
- **defect:** Fix nested enum by moving `GeodeticDatum` outside of `GeospatialElevation`.

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

[Unreleased]: https://github.com/sixafter/types/compare/v1.61.0...HEAD
[1.61.0]: https://github.com/sixafter/types/compare/v1.60.2...v1.61.0
[1.60.2]: https://github.com/sixafter/types/compare/v1.60.1...v1.60.2
[1.60.1]: https://github.com/sixafter/types/compare/v1.60.0...v1.60.1
[1.60.0]: https://github.com/sixafter/types/compare/v1.59.3...v1.60.0
[1.59.3]: https://github.com/sixafter/types/compare/v1.59.2...v1.59.3
[1.59.2]: https://github.com/sixafter/types/compare/v1.59.1...v1.59.2
[1.59.1]: https://github.com/sixafter/types/compare/v1.59.0...v1.59.1
[1.59.0]: https://github.com/sixafter/types/compare/v1.58.0...v1.59.0
[1.58.0]: https://github.com/sixafter/types/compare/v1.57.4...v1.58.0
[1.57.4]: https://github.com/sixafter/types/compare/v1.57.3...v1.57.4
[1.57.3]: https://github.com/sixafter/types/compare/v1.57.2...v1.57.3
[1.57.2]: https://github.com/sixafter/types/compare/v1.57.1...v1.57.2
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
