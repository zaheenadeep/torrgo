package metainfo

import (
	"os"
	"fmt"
	"crypto/sha1"
	"github.com/anacrolix/torrent/bencode"
)

type FileInfo struct {
	Length      int      `bencode:"length"`
	MD5Sum      string   `bencode:"md5sum,omitempty"`
	Path        []string `bencode:"path"`
}

type Info struct {
	PieceLength int    `bencode:"piece length"`
	Pieces      []byte `bencode:"pieces"`
	Private     int    `bencode:"private,omitempty"`
	Name        string `bencode:"name"`

	// Single File Info Only
	Length      int    `bencode:"length,omitempty"`  // mandatory
	MD5Sum      string `bencode:"md5sum,omitempty"`  // optional

	// Multi File Info Only
	Files       []FileInfo `bencode:"files,omitempty"` // mandatory
}

type Metainfo struct {
	Info         Info
	InfoBytes    bencode.Bytes `bencode:"info"`
	Announce     string        `bencode:"announce"`
	AnnounceList [][]string	   `bencode:"announce-list,omitempty"`
	CreationDate int           `bencode:"creation date,omitempty"`
	Comment      string        `bencode:"comment,omitempty"`
	CreatedBy    string        `bencode:"created by,omitempty"`
	Encoding     string        `bencode:"encoding,omitempty"`
}

func Load(name string) (*Metainfo, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile: %v", err)
	}

	var mi Metainfo
	err = bencode.Unmarshal(data, &mi)
	if err != nil {
		return nil, fmt.Errorf("bencode.Unmarshal: %v", err)
	}

	err = bencode.Unmarshal(mi.InfoBytes, &mi.Info)
	if err != nil {
		return nil, fmt.Errorf("bencode.Unmarshal: %v", err)
	} else if mi.Info.Files != nil && mi.Info.Length != 0 { // TODO: 0 byte file?
		return nil, fmt.Errorf("Info dict has both files and length keys")
	} else {
		return &mi, nil
	}
}

func (mi *Metainfo) InfoHash() []byte {
	b := sha1.Sum(mi.InfoBytes)
	return b[:]
}

func (mi *Metainfo) TotalSize() int {
	return mi.Info.PieceLength * len(mi.Info.Pieces)
}
