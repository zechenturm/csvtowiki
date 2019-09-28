package csvToWiki

import "bytes"

type Converter struct {

}

func (c Converter) Convert(input []byte) []byte {
	bts := parseRow(input)
	return bytes.ReplaceAll(bts, []byte("\n"), []byte(" |\n| "))
}

func parseRow(input []byte) []byte {
	b := bytes.Buffer{}
	b.Write([]byte("| "))
	b.Write(bytes.ReplaceAll(input, []byte(","), []byte(" |")))
	b.Write([]byte(" |"))
	return b.Bytes()
}
