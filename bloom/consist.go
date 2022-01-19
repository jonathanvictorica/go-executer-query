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


   111100000


  00100000000
*/

const numberForRow = 64

func (m *Concise) Add(hash uint64) {
	//positionMapBit := uint64(0)
	//if hash > numberForRow {
	positionMapBit := hash / numberForRow
	positionBlock := getPosition(hash)
	// valueAdd := 1 << positionBlock
	m.mapBit[positionMapBit] = m.mapBit[positionMapBit] | positionBlock
}
func (m *Concise) Remove(hash uint64) {
	positionMapBit := hash / numberForRow
	positionBlock := getPosition(hash)
	// valueAdd := 1 << positionBlock
	m.mapBit[positionMapBit] = m.mapBit[positionMapBit] & ^positionBlock
}

func getPosition(hash uint64) uint64 {
	positionBlock := hash % numberForRow

	if positionBlock != 0 {
		return uint64(2) << uint64(positionBlock-1)
	}
	return 1
}
func (m *Concise) Contains(hash uint64) bool {
	positionMapBit := uint64(0)
	if hash >= numberForRow {
		positionMapBit = hash / numberForRow
	}
	// 00000111111000
	// 00001000000000
	// 00000000000000
	positionBlock := getPosition(hash)
	return (m.mapBit[positionMapBit] & positionBlock) == positionBlock
}
