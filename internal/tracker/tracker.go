package tracker

import (
	"os"
	"fmt"
	"time"
	"net/http"
	"hash/fnv"
	"encoding/base64"
	"github.com/zaheenadeep/torrgo/internal/metainfo"
)

const (
	ClientID = "TG"
	ClientVersion  = "0001"
)

func Announce(mi *metainfo.Metainfo) error {
	c := &http.Client{}
	req, err := http.NewRequest("GET", mi.Announce, nil)
	if err != nil {
		return fmt.Errorf("http.NewRequest: %v", err)
	}
	req.Header.Add("info_hash", base64.URLEncoding.EncodeToString(mi.InfoHash()))
	req.Header.Add("peer_id", GeneratePeerID())
	req.Header.Add("port", )
	req.Header.Add("uploaded", )
	req.Header.Add("downloaded", )
	req.Header.Add("left", )
	req.Header.Add("compact", )
	req.Header.Add("no_peer_id", )
	req.Header.Add("event", )
	req.Header.Add("ip", )
	req.Header.Add("numwant", )
	req.Header.Add("key", )
	req.Header.Add("trackerid", )
	resp, err := c.Do(req)
	return err
}

func GeneratePeerID() string {
	h := fnv.New32()
	h.Write([]byte(string(os.Getpid())))
	h.Write([]byte(time.Now().String()))
	return fmt.Sprintf("-%s%s-%20d", ClientID, ClientVersion, h.Sum32())
}
