package shaproxy

import (
	"io"
)

type Reader struct {
	io.Reader
	shahash *ShaProxy
}

func (r *Reader) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)
	r.shahash.ShaHash.Write(p[:n])
	return

}

func (r *Reader) Close() (err error) {
	if closer, ok := r.Reader.(io.Closer); ok {
		return closer.Close()
	}
	return
}
