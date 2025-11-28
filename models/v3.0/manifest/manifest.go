package manifest

import (
	"github.com/phd-kerger/gbfs-go-adapter/common"
	"github.com/phd-kerger/gbfs-go-adapter/models/v3.0/header"
)

// An index of gbfs.json URLs for each GBFS data set produced by a publisher. A single
// instance of this file should be published at a single stable URL, for example:
// https://example.com/gbfs/manifest.json.
type Manifest struct {
	header.HeaderStruct
	Data Data `json:"data"`
}

type Data struct {
	// An array of objects, each containing the keys below.
	Datasets []Dataset `json:"datasets"`
}

type Dataset struct {
	// The system_id from system_information.json for the corresponding data set(s).
	SystemID common.ID `json:"system_id"`
	// Contains one object, as defined below, for each of the available versions of a feed. The
	// array MUST be sorted by increasing MAJOR and MINOR version number.
	Versions []common.VersionElement `json:"versions"`
}
