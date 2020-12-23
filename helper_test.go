package bitfinex

import (
	"fmt"
	"io/ioutil"
)

func LoadFixture(name string) []byte {
	b, err := ioutil.ReadFile(fixturePath(name))
	if err != nil {
		panic(err)
	}

	return b
}

func fixturePath(name string) string {
	return fmt.Sprintf("fixtures/%s", name)
}
