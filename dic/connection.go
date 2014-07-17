package dic

import (
        "bytes"
        "encoding/gob"

        "github.com/soundTricker/kagome/data"
)

type ConnectionTable struct {
	RowSize, ColSize int16
	Vec              []int16
}

func (this *ConnectionTable) At(a_row, a_col int) (cost int16, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	pos := int(this.ColSize)*a_row + a_col
	return this.Vec[pos], nil
}

var Connection *ConnectionTable

func init() {
        vec, err := data.Asset("data/connection.dic")
        if err != nil {
                panic(err)
        }
        decorder := gob.NewDecoder(bytes.NewBuffer(vec))
        if err = decorder.Decode(&Connection); err != nil {
		panic(err)
	}
}
