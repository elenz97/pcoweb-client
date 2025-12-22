package modbus

import (
	"encoding/binary"
	"time"

	"github.com/elenz97/pcoweb-client/pkg/metrics"
	"github.com/goburrow/modbus"
)

func NewBus(host string, typ metrics.GlenDimplexAnalogVariables) (*Bus, error) {
	handler := modbus.NewTCPClientHandler(host + ":502")
	handler.Timeout = 10 * time.Second
	handler.SlaveId = 0x7F

	if err := handler.Connect(); err != nil {
		return nil, err
	}

	return &Bus{Handler: handler, Client: modbus.NewClient(handler), GlenDimplexAnalogVariables: typ}, nil
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

func (bus *Bus) Observe(metrics map[metrics.GlenDimplexAnalogVariableName]int16) error {
	results, err := bus.ReadInputRegisters(1, 125)
	for i := 0; i+1 < len(results); i += 2 {
		nm := bus.Names[uint16(i/2)+1]
		if nm == "" {
			continue
		}
		metrics[nm] = int16(binary.BigEndian.Uint16(results[i : i+2]))
	}
	return err
}
