package metainfo

import (
	"os"
	"fmt"
	"github.com/anacrolix/torrent/bencode"
)

type SingleFileInfo struct {
	PieceLength int    `bencode:"piece length"`
	Pieces      string `bencode:"pieces"`
	Private     int    `bencode:"private,omitempty"`
	Name        string `bencode:"name"`
	Length      int    `bencode:"length"`
	MD5Sum      string `bencode:"md5sum,omitempty"`
}

type MIFile struct {
	Length      int      `bencode:"length"`
	MD5Sum      string   `bencode:"md5sum,omitempty"`
	Path        []string `bencode:"path"`
}

type MultiFileInfo struct {
	PieceLength int      `bencode:"piece length"`
	Pieces      string   `bencode:"pieces"`
	Private     int      `bencode:"private,omitempty"`
	Name        string   `bencode:"name"`
	Files       []MIFile `bencode:"md5sum"`
}

type Metainfo struct {
	InfoBytes    bencode.Bytes `bencode:"info"`
	Announce     string        `bencode:"announce"`
	AnnounceList [][]string	   `bencode:"announce-list,omitempty"`
	CreationDate int           `bencode:"creation date,omitempty"`
	Comment      string        `bencode:"comment,omitempty"`
	CreatedBy    string        `bencode:"created by,omitempty"`
	Encoding     string        `bencode:"encoding,omitempty"`
}

func UnmarshalMetainfo(name string) (*Metainfo, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile: %v", err)
	}

	var mi Metainfo
	err = bencode.Unmarshal(data, &mi)
	if err != nil {
		return nil, fmt.Errorf("bencode.Unmarshal: %v", err)
	}

	return &mi, nil
}
