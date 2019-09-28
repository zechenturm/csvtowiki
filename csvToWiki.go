package csvToWiki

type Converter struct {

}

func (c Converter) Convert(input []byte) []byte {
	return []byte("| a |")
}
