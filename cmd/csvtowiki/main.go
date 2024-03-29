package main

import (
	"fmt"
	"github.com/zechenturm/csvtowiki"
	"io/ioutil"
	"os"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		_, err = fmt.Fprint(os.Stderr, "error reading input:", err)
		panic(err)
	}
	c := csvtowiki.Converter{}
	fmt.Println(string(c.Convert(input)))
}
