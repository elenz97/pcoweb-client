package metrics

import (
	"encoding/binary"
	"fmt"
	"strings"
)

var GlenDimplexAnalogVariablesMapping = PCOType{
	Names: map[uint16]string{
		1:   "Outside temp",
		2:   "Return / House temp",
		3:   "Hot water temp",
		5:   "Flow (in) temp",
		8:   "High pressure sensor (bar)",
		29:  "Heating Setpoint",
		53:  "Heating Goal Temperature",
		58:  "Hot water Setpoint",
		96:  "Heating Power Level (unsure)",
		71:  "Additional Pump (operating hours)",
		72:  "Compressor 1 (operating hours)",
		73:  "Compressor 2 (operating hours)",
		74:  "Fan (operating hours)",
		76:  "Heating Pump (operating hours)",
		77:  "Hot water Pump (operating hours)",
		101: "Low-Pressure Sensor (bar)",
	},
}

func (bs DigitalVariables) String() string {
	var buf strings.Builder
	for _, b := range bs {
		if b {
			buf.WriteByte('1')
		} else {
			buf.WriteByte('0')
		}
	}
	return buf.String()
}

func (bs DigitalVariables) DiffIndex(as DigitalVariables) int {
	if len(as) != len(bs) {
		return 0
	}
	for i, b := range as {
		if b != bs[i] {
			return i
		}
	}
	return -1
}

func (is IntegerVariables) String() string {
	var buf strings.Builder
	buf.WriteByte('[')
	for i, u := range is {
		if i != 0 {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", u)
	}
	buf.WriteByte(']')
	return buf.String()
}

func (is IntegerVariables) DiffIndex(js IntegerVariables) int {
	if len(is) != len(js) {
		return 0
	}
	for i, v := range is {
		if v != js[i] {
			return i
		}
	}
	return -1
}

func (typ PCOType) NewMeasurement() *Measurement {
	return &Measurement{
		AnalogVariables:   make(AnalogVariables, len(typ.Names)),
		IntegerVariables:  make(IntegerVariables, typ.Length),
		DigitialVariables: make(DigitalVariables, typ.Length),
	}
}

func (bus *Bus) Close() error {
	bus.mu.Lock()
	handler := bus.Handler
	bus.Handler = nil
	bus.mu.Unlock()
	if handler != nil {
		return handler.Close()
	}
	return nil
}

func (bus *Bus) Integers(dest []uint16, offset int) error {
	results, err := bus.ReadDiscreteInputs(uint16(offset)+1, uint16(len(dest)))
	for i := 0; i+1 < len(results); i += 2 {
		dest[i/2] = binary.BigEndian.Uint16(results[i : i+2])
	}
	return err
}

func (bus *Bus) Coils(bits []bool, offset int) error {
	results, err := bus.ReadCoils(uint16(offset)+1, uint16(len(bits)))
	for i, b := range results {
		for j := uint(0); j < 8; j++ {
			bits[uint(i)*8+j] = b&(1<<j) != 0
		}
	}
	return err
}

func (bus *Bus) Observe(m map[string]int16) error {
	results, err := bus.ReadInputRegisters(1, 125)
	for i := 0; i+1 < len(results); i += 2 {
		nm := bus.Names[uint16(i/2)+1]
		if nm == "" {
			continue
		}
		m[nm] = int16(binary.BigEndian.Uint16(results[i : i+2]))
	}
	return err
}

func (bus *Bus) Bits(bits []bool) error {
	results, err := bus.ReadCoils(1, bus.Length)
	n := 0
	for i := 0; i+1 < len(results); i += 2 {
		for k := range []int{1, 0} {
			res := results[i+k]
			for j := uint8(0); j < 8; j++ {
				bits[n] = false
				if res&(1<<j) != 0 {
					bits[n] = true
				}
				n++
				if n == len(bits) {
					return err
				}
			}
		}
	}
	return err
}
