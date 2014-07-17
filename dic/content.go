package dic

import (
	"fmt"
        "bytes"
        "encoding/gob"

        "github.com/soundTricker/kagome/data"
)


type Content struct {
	Pos, Pos1, Pos2, Pos3,
	Katuyougata, Katuyoukei, Kihonkei,
	Yomi, Pronunciation string
}

func (this Content) String() string {
	return fmt.Sprintf("%v, %v, %v, %v, %v, %v, %v, %v, %v",
		this.Pos, this.Pos1, this.Pos2, this.Pos3,
		this.Katuyougata, this.Katuyoukei, this.Kihonkei,
		this.Yomi, this.Pronunciation)
}

var Contents []Content

func init() {
        vec, err := data.Asset("data/contents.dic")
        if err != nil {
                panic(err)
        }
        decorder := gob.NewDecoder(bytes.NewBuffer(vec))
        if err = decorder.Decode(&Contents); err != nil {
		panic(err)
	}
}
