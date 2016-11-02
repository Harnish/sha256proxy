This code can do inline sha256 calculations


```golang
package main

import (
    "../shaproxy"
    "fmt"
    "io"
    "os"
)

func main() {
    shain := shaproxy.New()
    file := "shaproxy"
    f, err := os.Open(file)
    if err != nil {
        fmt.Println(err)
    }

    fo, err := os.Create("junk.file")
    if err != nil {
        fmt.Println(err)
    }

    reader := shain.NewProxyReader(f)
    io.Copy(fo, reader)
    shain.Finish()
    fmt.Println(shain.SumHex())

}
```

I used gopkg.in/cheggaaa/pb.v1 as the example for this as I was trying to add progress to my app that sha256s the file before and after the transaction.  
