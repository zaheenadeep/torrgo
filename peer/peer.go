package peer

import (
	"bufio"
	"io"
	"fmt"
	"net"
	"bytes"
	"strings"
)

// TODO change Peer to a suitable type
type Peer struct {
	PeerId      string
	PeerAddress string
	Port        int64
}

const (
	PSTR         = "BitTorrent protocol"
	NINFOHASH    = 20
	NRESERVED    = 8
)

func (p *Peer) Handshake(infoHash []byte) error {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", p.PeerAddress, p.Port))
	if err != nil {
		return fmt.Errorf("net.Dial: %v", err)
	}

	defer conn.Close()

	// prepare and send
	var b bytes.Buffer
	b.WriteByte(byte(len(PSTR)))  // 19
	b.WriteString(PSTR)
	b.Write(make([]byte, NRESERVED))  // four 0 bytes
	b.Write(infoHash)
	n, err := b.WriteTo(conn)  // send
	if err != nil {
		return fmt.Errorf("b.WriteTo: %v", err)
	}
	if n != 48 {
		return fmt.Errorf("b.WriteTo wrote %d bytes", n)
	}

	// recv and check
	r := bufio.NewReader(conn)

	pstrlen, err := r.ReadByte()
	if err != nil {
		return fmt.Errorf("w.ReadByte: %v", err)
	}

	var pstr strings.Builder
	_, err = io.CopyN(&pstr, r, int64(pstrlen))
	if err != nil {
		return fmt.Errorf("io.CopyN: %v", err)
	}
	if pstr.String() != PSTR {
		return fmt.Errorf("pstrb is %s", pstr.String())
	}

	for i := 0; i < NRESERVED; i++ {
		by, err := r.ReadByte()
		if err != nil {
			return fmt.Errorf("r.ReadByte: %v", err)
		}
		if by != byte(0) {
			return fmt.Errorf("by is %b", by)
		}
	}

	h := make([]byte, NINFOHASH)
	_, err = io.ReadFull(r, h)
	if err != nil {
		return fmt.Errorf("io.ReadFull: %v", err)
	}
	if !bytes.Equal(h, infoHash) {
		return fmt.Errorf("infoHashes do not match:\n%v\n%v", h, infoHash)
	}

	// send peer_id
	m, err := fmt.Fprint(conn, p.PeerId)
	if err != nil {
		return fmt.Errorf("fmt.Fprint: %v", err)
	}
	if m != NINFOHASH {
		return fmt.Errorf("fmt.Fprint: %d bytes sent", m)
	}

	return nil
}
