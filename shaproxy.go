// Package shaproxy provides a stream reader object to allow calculating mulitple sha256s
// inline.  This is useful when capturing the sha256 prior to transformation of data and
// after.
//
package shaproxy

import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"io"
)

type ShaProxy struct {
	// ShaHash is the Hash object that is added to as the stream is read
	ShaHash hash.Hash
	// sumBytes is an internal value to take the output of the final checksum, in bytes.
	sumBytes []byte
	// sumHex is an internal value to store the final checksum in Hex as a string.
	sumHex string
	// finished is an internal bool to signify that you have called finish and calculated the final sha256.
	finished bool
}

// New creates and returns a shaproxy object.
func New() *ShaProxy {
	sha256hash := sha256.New()
	me := &ShaProxy{
		ShaHash: sha256hash,
	}
	return me
}

// This takes an io.Reader and returns an io.Reader.  It does not modify the data in the stream but adds
// to the object's hash.
func (me *ShaProxy) NewProxyReader(r io.Reader) *Reader {
	return &Reader{r, me}
}

// Returns the value from the private sumBytes variable in a byteslice representation.
// Requires Finish to be called before the value is populated.
func (me *ShaProxy) SumBytes() []byte {
	return me.sumBytes
}

// Retruns the private variable sumHex, which is the hex representation in a string.
// Requires Finish to be called before the value is populated.
func (me *ShaProxy) SumHex() string {
	return me.sumHex
}

// This executes the final calculation of the sha256 sum and sets the values.
func (me *ShaProxy) Finish() {
	if !me.finished {
		me.sumBytes = me.ShaHash.Sum(nil)
		me.sumHex = hex.EncodeToString(me.sumBytes)
		me.finished = true
	}
}
