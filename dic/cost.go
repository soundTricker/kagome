package dic

import (
        "bytes"
        "encoding/gob"

        "github.com/soundTricker/kagome/data"
)

type Cost struct {
	Left, Right, Weight int16
}

var Costs []Cost

func init() {
        vec, err := data.Asset("data/costs.dic")
        if err != nil {
                panic(err)
        }
        decorder := gob.NewDecoder(bytes.NewBuffer(vec))
        if err = decorder.Decode(&Costs); err != nil {
		panic(err)
	}
}
