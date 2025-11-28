package gbfs

import (
	"github.com/phd-kerger/gbfs-go-adapter/models/common"
	"github.com/phd-kerger/gbfs-go-adapter/models/v2.3/header"
)

// Auto-discovery file that links to all of the other files published by the system.
type Gbfs struct {
	header.HeaderStruct
	// Response data in the form of name:value pairs.
	Data Data `json:"data"`
}

// The language that will be used throughout the rest of the files. It MUST match the value in the system_information.json file.
type Data map[common.Language]LanguageFeeds

type LanguageFeeds struct {
	// An array of all of the feeds that are published by this auto-discovery file. Each element in the array is an object with the keys below.
	Feeds []common.Feed `json:"feeds"`
}
