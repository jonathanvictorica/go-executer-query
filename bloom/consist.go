package bloom

import (
	"fmt"
	"strconv"
)

type Concise struct {
	mapBit map[uint64]uint64
}

func NewConcise() *Concise {
	return &Concise{
		mapBit: make(map[uint64]uint64),
	}
}

/*
   3
   5

   10000000 = 8
   00001000 = 8

*/

const numberForRow = 8

func (m *Concise) Add(hash uint64) {
	if m.Contains(hash) {
		return
	}
	positionMapBit := uint64(0)
	if hash > numberForRow {
		positionMapBit = hash / numberForRow
	}
	positionBlock := hash % numberForRow

	valueAdd := 1 << positionBlock
	m.mapBit[positionMapBit] = m.Compress(m.mapBit[positionMapBit] + uint64(valueAdd))
}
func (m *Concise) Remove(hash uint64) {
	if !m.Contains(hash) {
		return
	}
	positionMapBit := uint64(0)
	if hash > numberForRow {
		positionMapBit = hash / numberForRow
	}
	positionBlock := hash % numberForRow
	valueAdd := 1 << positionBlock
	m.mapBit[positionMapBit] = m.mapBit[positionMapBit] - uint64(valueAdd)
}
func (m *Concise) Contains(hash uint64) bool {
	positionMapBit := uint64(0)
	if hash > numberForRow {
		positionMapBit = hash / numberForRow
	}
	positionBlock := hash % numberForRow
	valueBinary := strconv.FormatUint(m.mapBit[positionMapBit], 2)
	if uint64(len(valueBinary)) < positionBlock {
		return false
	}

	return valueBinary[positionBlock-1] == 1
}

func (m *Concise) Compress(u uint64) uint64 {
	representation := strconv.FormatUint(u, 2)
	representation = "11001000"

	//cadena1 := representation[0:4]

	fmt.Sprint(representation)
	return u
}
