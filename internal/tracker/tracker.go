package tracker

import (
	"encoding/base64"
	"fmt"
	"github.com/zaheenadeep/torrgo/internal/metainfo"
	"hash/fnv"
	"net/http"
	"os"
	"time"
	"strconv"
	"io"
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
	req.Header.Add("peer_id", generatePeerID())
	req.Header.Add("port", strconv.Itoa(Port))
	req.Header.Add("uploaded", "0")
	req.Header.Add("downloaded", "0")
	req.Header.Add("left", strconv.Itoa(mi.TotalSize()))
	req.Header.Add("compact", "0")  // TODO: Implement 1
	req.Header.Add("no_peer_id", "0")
	req.Header.Add("event", "started")
	// req.Header.Add("ip")
	// req.Header.Add("numwant")
	// req.Header.Add("key")
	// req.Header.Add("trackerid")
	resp, err := c.Do(req)
	if err != nil {
		return fmt.Errorf("c.Do: %v", err)
	} else {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Printf("%s", body)
		fmt.Println(resp.Header)
		return nil
	}
}

func generatePeerID() string {
	h := fnv.New32()
	h.Write([]byte(strconv.Itoa(os.Getpid())))
	h.Write([]byte(time.Now().String()))
	return fmt.Sprintf("-%s%s-%20d", ClientID, ClientVersion, h.Sum32())
}
