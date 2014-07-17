package dic

import (
        "bytes"
        "encoding/gob"

        "github.com/soundTricker/kagome/data"
	"github.com/soundTricker/kagome/trie/da"
)

var (
	Counts map[int]int
	Index *da.DoubleArray
)

func init() {
	vec, err := data.Asset("data/index.dic")
        if err != nil {
                panic(err)
        }
        decorder := gob.NewDecoder(bytes.NewBuffer(vec))
        if err = decorder.Decode(&Counts); err != nil {
		panic(err)
	}
	if err = decorder.Decode(&Index); err != nil {
		panic(err)
	}
}
