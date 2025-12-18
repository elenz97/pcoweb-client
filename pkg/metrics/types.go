package metrics

import (
	"sync"

	"github.com/goburrow/modbus"
)

type IntegerVariables []uint16

type AnalogVariables map[string]int16

type DigitalVariables []bool

type Measurement struct {
	AnalogVariables  AnalogVariables
	IntegerVariables IntegerVariables
	// aka "Bits"
	DigitialVariables DigitalVariables
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
