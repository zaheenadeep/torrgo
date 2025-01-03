package metainfo

import (
	"testing"
	"fmt"
)

func TestArtOfWar(t *testing.T) {
	testFile("../testfiles/artofwar.torrent", t)
}

func TestHitchhikers(t *testing.T) {
	testFile("../testfiles/hitchhikersguide.torrent", t)
}

func testFile(name string, t *testing.T) {
	mi, err := Load(name)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v\n", mi)
}
