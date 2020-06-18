package main;

import (
  "bufio"
  enc "encoding/ascii85"
  "os"
)

func main() {
  onionReader := bufio.NewReader(os.Stdin);
  onionDecoder := enc.NewDecoder(onionReader);
  bufDecoder := bufio.NewReader(onionDecoder);
  bufDecoder.WriteTo(os.Stdout);
}


