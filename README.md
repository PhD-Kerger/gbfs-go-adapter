# GBFS Go JSON Schema and Language Bindings

This repository provides JSON Schema definitions and Go language bindings for the General Bikeshare Feed Specification (GBFS). It extends the original GBFS JSON Schema maintained by MobilityData to support all versions of the GBFS specification.

## Features

- Comprehensive JSON Schema definitions for all GBFS versions.
- Go language bindings for easy integration with Go applications.

## Installation

To use the GBFS JSON Schema and Go bindings in your project, you can either clone this repository:

```bash
git clone https://github.com/PhDKerger/gbfs-json-schema.git
cd gbfs-json-schema
```

or you can use our Go module directly by adding it to your `go.mod` file:

```bashgo
require github.com/PhDKerger/gbfs-json-schema v0.1.0
```

## Usage

You can use the provided JSON Schema files to validate GBFS feeds in your applications. The Go bindings can be imported into your Go projects for seamless integration.

## Related Projects

The original code basis of this repository was forked from the official [gbfs-json-schema](https://github.com/MobilityData/gbfs-json-schema) repository maintained by MobilityData, which contains the authoritative JSON Schema files for GBFS feeds. As the official Go connector for GBFS does not support all versions of the GBFS specification, this repository extends the original schemas to provide comprehensive support for all GBFS versions.
