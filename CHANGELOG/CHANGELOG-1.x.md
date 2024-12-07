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

[Unreleased]: https://github.com/sixafter/nanoid/compare/v1.24.0...HEAD
[1.24.0]: https://github.com/sixafter/nanoid/compare/v1.23.0...v1.24.0
[1.23.0]: https://github.com/sixafter/nanoid/compare/v1.22.0...v1.23.0
[1.22.0]: https://github.com/sixafter/nanoid/compare/v1.0.1...v1.22.0

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
