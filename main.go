// Copyright 2019 Tamás Gulácsi
//
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	vmetrics "github.com/VictoriaMetrics/metrics"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/tgulacsi/pcoweb-client/pkg/config"
	"github.com/tgulacsi/pcoweb-client/pkg/metrics"

	_ "net/http/pprof"

	"golang.org/x/sync/errgroup"

	"github.com/goburrow/modbus"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("%+v", err)
	}
}

func run() error {
	f := &config.Flags{}
	f.Bind(pflag.CommandLine)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	metricsAddr := ":9112"
	if f.MetricsAddr != "" {
		metricsAddr = f.MetricsAddr
	}

	bus, err := NewBus(f.Host, metrics.GlenDimplexAnalogVariablesMapping)
	if err != nil {
		return err
	}

	defer func(bus *metrics.Bus) {
		err := bus.Close()
		if err != nil {
			log.Printf("error closing bus: %v", err)
		}
	}(bus)

	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		defer cancel()
		sig := <-ch
		log.Println("SIGNAL", sig)
		if p, _ := os.FindProcess(os.Getpid()); p != nil {
			time.Sleep(time.Second)
			_ = p.Signal(sig)
		}
	}()
	errgroup, ctx := errgroup.WithContext(ctx)

	http.HandleFunc("/metrics", func(w http.ResponseWriter, req *http.Request) {
		vmetrics.WritePrometheus(w, true)
	})

	errgroup.Go(func() error {
		return http.ListenAndServe(metricsAddr, nil)
	})

	var mutex sync.Mutex
	measurement := metrics.GlenDimplexAnalogVariablesMapping.NewMeasurement()

	interval := 10 * time.Second
	ticker := time.NewTicker(interval)
	initialRun := true

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		if err = bus.Observe(measurement.AnalogVariables); err != nil {
			return err
		}
		if initialRun {
			for k := range measurement.AnalogVariables {
				vmetrics.NewGauge(fmt.Sprintf("modbus_glendimplex_analog{name=%q}", k),
					func() float64 {
						mutex.Lock()
						v := float64(measurement.AnalogVariables[k]) / 10.0
						mutex.Unlock()

						return v
					})
			}
		}

		if len(measurement.IntegerVariables) > 0 {
			if err = bus.Integers(measurement.IntegerVariables, 0); err != nil {
				return err
			}
			if initialRun {
				for i := range measurement.IntegerVariables {
					i := i
					vmetrics.NewGauge(
						fmt.Sprintf("modbus_glendimplex_integer{index=\"i%03d\"}", i),
						func() float64 {
							mutex.Lock()
							v := measurement.IntegerVariables[i]
							mutex.Unlock()
							return float64(v)
						})
				}
			}
		}

		if len(measurement.DigitialVariables) > 0 {
			if err = bus.Bits(measurement.DigitialVariables); err != nil {
				return err
			}
			if initialRun {
				for i := range measurement.DigitialVariables {
					i := i
					vmetrics.NewGauge(
						fmt.Sprintf("modbus_glendimplex_digital{index=\"b%03d\"}", i),
						func() float64 {
							var j float64
							mutex.Lock()
							b := measurement.DigitialVariables[i]
							mutex.Unlock()
							if b {
								j = 1
							}
							return j
						})
				}
			}
		}

		if initialRun {
			logrus.Printf("exporting %d (analog), %d (integer), %d (digital/bits) metrics",
				len(measurement.AnalogVariables), len(measurement.IntegerVariables), len(measurement.DigitialVariables))
		}
		initialRun = false

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
		}
	}
}

func NewBus(host string, typ metrics.PCOType) (*metrics.Bus, error) {
	handler := modbus.NewTCPClientHandler(host + ":502")
	handler.Timeout = 10 * time.Second
	handler.SlaveId = 0x7F

	if err := handler.Connect(); err != nil {
		return nil, err
	}
	return &metrics.Bus{Handler: handler, Client: modbus.NewClient(handler), PCOType: typ}, nil
}
