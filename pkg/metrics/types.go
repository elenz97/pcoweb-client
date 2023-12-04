package metrics

import (
	"github.com/goburrow/modbus"
	"sync"
)

type Ints []uint16

type Map map[string]int16

type Bits []bool

type Measurement struct {
	Map  Map
	Ints Ints
	Bits Bits
}

type Bus struct {
	PCOType

	mu sync.Mutex
	modbus.Client
	Handler *modbus.TCPClientHandler
}

type PCOType struct {
	Length      uint16
	Names, Bits map[uint16]string
	AlertBits   []uint16
	AlertNames  []string
}
