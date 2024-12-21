package tracker

import (
	"encoding/base64"
	"fmt"
	"github.com/zaheenadeep/torrgo/internal/metainfo"
	"hash/fnv"
	"net/http"
	"os"
	"time"
)

const (
	ClientID      = "TG"
	ClientVersion = "0001"
	Port          = 6881
)

func Announce(mi *metainfo.Metainfo) error {
	c := &http.Client{}
	req, err := http.NewRequest("GET", mi.Announce, nil)
	if err != nil {
		return fmt.Errorf("http.NewRequest: %v", err)
	}
	req.Header.Add("info_hash", base64.URLEncoding.EncodeToString(mi.InfoHash()))
	req.Header.Add("peer_id", GeneratePeerID())
	req.Header.Add("port", string(Port))
	req.Header.Add("uploaded", "0")
	req.Header.Add("downloaded", "0")
	req.Header.Add("left", string(mi.TotalSize()))
	req.Header.Add("compact", "0")  // TODO: Implement 1
	req.Header.Add("no_peer_id", "0")
	req.Header.Add("event", "started")
	// req.Header.Add("ip")
	// req.Header.Add("numwant")
	// req.Header.Add("key")
	// req.Header.Add("trackerid")
	resp, err := c.Do(req)
	return err
}

func GeneratePeerID() string {
	h := fnv.New32()
	h.Write([]byte(string(os.Getpid())))
	h.Write([]byte(time.Now().String()))
	return fmt.Sprintf("-%s%s-%20d", ClientID, ClientVersion, h.Sum32())
}
