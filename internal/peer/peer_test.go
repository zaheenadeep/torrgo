package peer

import (
	"io"
	"net"
	"testing"
	"time"
	"sync"
	"bytes"
)

func TestHandshake(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go serve(&wg, t)
	time.Sleep(time.Second)

	p := Peer{
		PeerId:      "aaaaaaaaaaaaaaaaaaaa",
		PeerAddress: "127.0.0.1",
		Port:        int64(8080),
	}
	h := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	err := p.Handshake(h)
	if err != nil {
		t.Fatal(err)
	}

	wg.Wait()
}

func serve(wg *sync.WaitGroup, t *testing.T) {
	defer wg.Done()

	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()
	
	conn, err := ln.Accept()
	if err != nil {
		t.Fatal(err)
	}
	ln.Close()
	defer conn.Close()
	
	// recv
	var buf bytes.Buffer
	n64, err := io.CopyN(&buf, conn, 48)
	if err != nil {
		t.Fatalf("io.CopyN %v", err)
	} else if n64 != 48 {
		t.Fatalf("Number of bytes: %d", n64)
	} else {
		t.Log("Bytes received:", buf.Bytes())
	}

	// send
	resp := []byte{
		0x13,  // pstrlen
		0x42, 0x69, 0x74, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,  // pstr
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,  // reserved
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10, 0x11, 0x12, 0x13, 0x14, // info_hash
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10, 0x11, 0x12, 0x13, 0x14, // peer_id
	}
	n, err := conn.Write(resp)
	if err != nil {
		t.Fatal(err)
	} else if n != len(resp) {
		t.Fatalf("Number of bytes sent: %x", n)
	}

	// recv peer_id
	b, err := io.ReadAll(conn)
	if err != nil {
		t.Fatalf("io.CopyN %v", err)
	} else if len(b) != 20 {
		t.Fatalf("%d bytes received instead of 20: % x", len(b), b)
	} else {
		t.Log("Bytes received:", b)
	}
}