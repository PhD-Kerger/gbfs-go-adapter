package header

import (
	"github.com/phd-kerger/gbfs-go-adapter/common"
)

type HeaderStruct struct {
	// Last time the data in the feed was updated in POSIX time (epoch seconds).
	LastUpdated common.PosixTimestamp `json:"last_updated"`
	// Number of seconds before the data in the feed will be updated again (0 if the data should always be refreshed).
	TTL uint64 `json:"ttl"`
	// GBFS version number to which the feed conforms, according to the versioning framework
	// (added in v1.1).
	Version common.Version `json:"version"`
}
