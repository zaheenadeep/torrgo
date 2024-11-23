package main

import (
	"os"
	"fmt"
	"github.com/zaheenadeep/torrgo/internal/metainfo"
	"github.com/zaheenadeep/torrgo/internal/peer"
	"github.com/zaheenadeep/torrgo/internal/tracker"
)

type State struct {
	InfoHash   string
	PeerID     string
	Port       int
	Uploaded   int
	Downloaded int
	Left       int
}

func main() {
	name := os.Args[1]
	mi, err := metainfo.Load(name)
	if err != nil {
		fmt.Errorf("metainfo.UnmarshalMetainfo: %v", err)
	}
	tracker.Initiate(mi)
	// peer.Handshake(mi.InfoHash())
}
