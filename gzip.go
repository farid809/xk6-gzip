package gzip

import (
    "compress/gzip"
    "bytes"
    "io/ioutil"
    "go.k6.io/k6/js/modules"
)

// Register the extension
func init() {
    modules.Register("k6/x/gzip", new(Gzip))
}

type Gzip struct{}

// Compress data
func (g *Gzip) Compress(data string) (string, error) {
    var buf bytes.Buffer
    writer := gzip.NewWriter(&buf)
    _, err := writer.Write([]byte(data))
    if err != nil {
        return "", err
    }
    writer.Close()
    return buf.String(), nil
}

// Decompress data
func (g *Gzip) Decompress(data string) (string, error) {
    reader, err := gzip.NewReader(bytes.NewBufferString(data))
    if err != nil {
        return "", err
    }
    result, err := ioutil.ReadAll(reader)
    if err != nil {
        return "", err
    }
    return string(result), nil
}