package csvToWiki

import "testing"

func TestCreate (t *testing.T) {
	_ = Converter{}
}

func TestSingleCell (t *testing.T) {
	c := Converter{}
	out := c.Convert([]byte("a"))
	compare(t, out, "| a |")
}

func TestSmallRow (t *testing.T) {
	c := Converter{}
	out := c.Convert([]byte("a, b"))
	compare(t, out, "| a | b |")
}

func compare(t *testing.T,  result[]byte, reference string) {
	if string(result) != reference {
		t.Fatalf("got wrong output format: wanted \"%s\", got \"%s\"", reference, string(result))
	}
}