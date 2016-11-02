package shaproxy

import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"io"
)

type ShaProxy struct {
	ShaHash  hash.Hash
	sumBytes []byte
	sumHex   string
	finished bool
}

func New() *ShaProxy {
	sha256hash := sha256.New()
	me := &ShaProxy{
		ShaHash: sha256hash,
	}
	return me
}

func (me *ShaProxy) NewProxyReader(r io.Reader) *Reader {
	return &Reader{r, me}
}

func (me *ShaProxy) SumBytes() []byte {
	return me.sumBytes
}

func (me *ShaProxy) SumHex() string {
	return me.sumHex
}

func (me *ShaProxy) Finish() {
	if !me.finished {
		me.sumBytes = me.ShaHash.Sum(nil)
		me.sumHex = hex.EncodeToString(me.sumBytes)
		me.finished = true
	}
}
