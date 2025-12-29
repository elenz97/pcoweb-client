package modbus

import (
	"sync"

	"github.com/elenz97/dimplex-pcoweb-exporter/pkg/metrics"
	"github.com/goburrow/modbus"
)

type Bus struct {
	metrics.GlenDimplexAnalogVariables

	mu sync.Mutex
	modbus.Client
	Handler *modbus.TCPClientHandler
}
