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

func TestSmallRow (t *testing.T) {
	c := Converter{}
	out := c.Convert([]byte("a, b"))
	if string(out) != "| a | b |" {
		t.Fatalf("got wrong output format: wanted \"| a | b |\", got \"%s\"", string(out))
	}
}