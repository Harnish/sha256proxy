package shaproxy

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func (me *ShaProxy) TestSha(t *testing.T) {
	testdata := "This is my test data.  It needs to be long enough to get a few pages worth of data."
	testdatasha256 := "e1a0facc7cd1ba737a0c96b625442ebffac97487021658c42634e97f4eabfb55"
	var output bytes.Buffer
	testobj := me.New()
	b := bytes.NewBufferString(testdata)
	reader := testobj.NewProxyReader(b)
	io.Copy(os.Stdout, reader)
	testobj.Finish()
	if testobj.SumHex() != testdatasha256 {
		t.Error("Expected ", testdatasha256, " got", testobj.SumHex())
	}
}
