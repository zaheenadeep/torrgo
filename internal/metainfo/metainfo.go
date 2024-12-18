package metainfo

import (
	"os"
	"fmt"
	"strings"
	"crypto/sha1"
	"github.com/anacrolix/torrent/bencode"
)

type SingleFileInfo struct {
	PieceLength int    `bencode:"piece length"`
	Pieces      []byte `bencode:"pieces"`
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
	Pieces      []byte   `bencode:"pieces"`
	Private     int      `bencode:"private,omitempty"`
	Name        string   `bencode:"name"`
	Files       []MIFile `bencode:"files"`
}

type FileInfo interface{}

type Metainfo struct {
	Info         FileInfo
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

	if strings.Contains(mi.InfoBytes.GoString(), ":pathl") {
		var m MultiFileInfo
		err = bencode.Unmarshal(mi.InfoBytes, &m)
		if err != nil {
			return nil, fmt.Errorf("bencode.Unmarshal: %v", err)
		}
		mi.Info = m
	} else {
		var s SingleFileInfo
		err = bencode.Unmarshal(mi.InfoBytes, &s)
		if err != nil {
			return nil, fmt.Errorf("bencode.Unmarshal: %v", err)
		}
		mi.Info = s
	}

	return &mi, nil
}

func (mi *Metainfo) InfoHash() []byte {
	b := sha1.Sum(mi.InfoBytes)
	return b[:]
}
