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

[Unreleased]: https://github.com/sixafter/nanoid/compare/v1.23.0..HEAD
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
