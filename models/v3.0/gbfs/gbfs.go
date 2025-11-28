package gbfs

import (
	"github.com/phd-kerger/gbfs-go-adapter/models/common"
	"github.com/phd-kerger/gbfs-go-adapter/models/v3.0/header"
)

// Auto-discovery file that links to all of the other files published by the system.
type Gbfs struct {
	header.HeaderStruct
	Data Data `json:"data"`
}

type Data struct {
	Feeds []common.Feed `json:"feeds"`
}
