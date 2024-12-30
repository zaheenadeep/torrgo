package tracker

import (
	"testing"
	"github.com/zaheenadeep/torrgo/internal/metainfo"
)

func TestArtOfWar(t *testing.T) {
	testFile("../testfiles/artofwar.torrent", t)
}

func TestHitchhikers(t *testing.T) {
	testFile("../testfiles/hitchhikersguide.torrent", t)
}

func testFile(name string, t *testing.T) {
	mi, err := metainfo.Load(name)
	if err != nil {
		t.Fatal(err)
	}
	err = Announce(mi)
	if err != nil {
		t.Fatal(err)
	}
}
