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
- **FEATURE:** Added [geometry.proto](geometry.proto) for defining geometrical constructs, including:
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

[Unreleased]: https://github.com/sixafter/types/compare/v1.40.0...HEAD
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
