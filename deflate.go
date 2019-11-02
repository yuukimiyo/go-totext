package totext

import (
    "fmt"
    "bytes"
    "compress/zlib"
    "io"
)


func Deflate(text string) (string, error){
    r, err := compress(bytes.NewBufferString(text))
    if err != nil {
        panic(err)
    }

    b := r.Bytes()
    fmt.Printf("%d bytes: %v\n", len(b), b)

	return "test", nil
}
