package csvToWiki

import "testing"

func TestCreate (t *testing.T) {
	_ = Converter{}
}

func TestSingleCell (t *testing.T) {
	c := Converter{}
	out := c.Convert([]byte("a"))
	if string(out) != "| a |" {
		t.Fatalf("got wrong output format: wanted \"| a |\", got \"%s\"", string(out))
	}
}