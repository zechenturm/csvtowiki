package csvtowiki

import "testing"

type tst struct {
	in string
	ref string
}

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

func TestLongerRows (t *testing.T) {
	data := []tst{
		tst{"a, b, c", "| a | b | c |"},
		tst{"1, 2, 3, a, v, 4", "| 1 | 2 | 3 | a | v | 4 |"},
		tst{"hello, world", "| hello | world |"},
		tst{"this is a test, 1.23, abcde", "| this is a test | 1.23 | abcde |"},
	}

	for _, d := range data {
		c := Converter{}
		compare(t, c.Convert([]byte(d.in)), d.ref)
	}
}

func TestMultipleRows (t *testing.T) {
	data := []tst{
		tst{"a, b, c\nd, e, f", "| a | b | c |\n| d | e | f |"},
		tst{"1\n2\n3\na\nv\n4", "| 1 |\n| 2 |\n| 3 |\n| a |\n| v |\n| 4 |"},
		tst{"hello\nworld", "| hello |\n| world |"},
		tst{"this is a test\n1.23, abcde", "| this is a test |\n| 1.23 | abcde |"},
	}

	for _, d := range data {
		c := Converter{}
		compare(t, c.Convert([]byte(d.in)), d.ref)
	}
}

func compare(t *testing.T,  result[]byte, reference string) {
	if string(result) != reference {
		t.Fatalf("got wrong output format: wanted \"%s\", got \"%s\"", reference, string(result))
	}
}
