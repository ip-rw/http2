package maybe_later

import (
	"bufio"
	"bytes"
	"io"
	"testing"

	"github.com/ip-rw/http2"
	"github.com/ip-rw/http2/http2utils"
)

const (
	testStr = "make fasthttp great again"
)

func TestFrameWrite(t *testing.T) {
	fr := http2.AcquireFrameHeader()
	defer http2.ReleaseFrameHeader(fr)

	n, err := io.WriteString(fr, testStr)
	if err != nil {
		t.Fatal(err)
	}
	if nn := len(testStr); n != nn {
		t.Fatalf("unexpected size %d<>%d", n, nn)
	}

	var bf = bytes.NewBuffer(nil)
	var bw = bufio.NewWriter(bf)
	fr.WriteTo(bw)
	bw.Flush()

	b := bf.Bytes()
	if str := string(b[9:]); str != testStr {
		t.Fatalf("mismatch %s<>%s", str, testStr)
	}
}

func TestFrameRead(t *testing.T) {
	var h [9]byte
	bf := bytes.NewBuffer(nil)
	br := bufio.NewReader(bf)

	http2utils.Uint24ToBytes(h[:3], uint32(len(testStr)))

	n, err := bf.Write(h[:9])
	if err != nil {
		t.Fatal(err)
	}
	if n != 9 {
		t.Fatalf("unexpected written bytes %d<>9", n)
	}

	n, err = io.WriteString(bf, testStr)
	if err != nil {
		t.Fatal(err)
	}
	if n != len(testStr) {
		t.Fatalf("unexpected written bytes %d<>%d", n, len(testStr))
	}

	fr := http2.AcquireFrame()
	defer http2.ReleaseFrame(fr)

	nn, err := fr.ReadFrom(br)
	if err != nil {
		t.Fatal(err)
	}
	n = int(nn)
	if n != len(testStr)+9 {
		t.Fatalf("unexpected readed bytes %d<>%d", n, len(testStr)+9)
	}

	if str := string(fr.Payload()); str != testStr {
		t.Fatalf("mismatch %s<>%s", str, testStr)
	}
}

// TODO: continue
