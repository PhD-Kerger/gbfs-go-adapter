package gbfsversions

import (
	"github.com/phd-kerger/gbfs-go-adapter/models/common"
	"github.com/phd-kerger/gbfs-go-adapter/models/v2.3/header"
)

// Lists all feed endpoints published according to version sof the GBFS documentation. (added in v1.1)
type GbfsVersions struct {
	header.HeaderStruct
	// Response data in the form of name:value pairs.
	Data Data `json:"data"`
}

// Response data in the form of name:value pairs.
type Data struct {
	// Contains one object, as defined below, for each of the available versions of a feed. The
	// array must be sorted by increasing MAJOR and MINOR version number.
	Versions []VersionElement `json:"versions"`
}

type VersionElement struct {
	// URL of the corresponding gbfs.json endpoint
	URL string `json:"url"`
	// The semantic version of the feed in the form X.Y
	Version common.Version `json:"version"`
}
