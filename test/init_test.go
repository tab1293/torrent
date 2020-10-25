package test

import (
	_ "github.com/anacrolix/envpprof"
	"github.com/tab1293/torrent"
)

func init() {
	torrent.TestingTempDir.Init("torrent-test.test")
}
