package bloom

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

const numberForRow = 18446744073709551615

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
	m.mapBit[positionMapBit] = m.mapBit[positionMapBit] + uint64(valueAdd)
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

	return false
}
