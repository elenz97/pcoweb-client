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
	"bytes"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	vmetrics "github.com/VictoriaMetrics/metrics"
	"github.com/spf13/pflag"
	"github.com/tgulacsi/pcoweb-client/pkg/config"
	"github.com/tgulacsi/pcoweb-client/pkg/metrics"

	_ "net/http/pprof"

	"golang.org/x/sync/errgroup"

	"github.com/goburrow/modbus"
)

var hostname string

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

	bus, err := NewBus(f.Host, GlenDimplexMapping)
	if err != nil {
		return err
	}
	if hostname, err = os.Hostname(); err != nil {
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
	grp, ctx := errgroup.WithContext(ctx)

	http.HandleFunc("/metrics", func(w http.ResponseWriter, req *http.Request) {
		vmetrics.WritePrometheus(w, true)
	})
	grp.Go(func() error {
		return http.ListenAndServe(metricsAddr, nil)
	})

	var mu sync.Mutex
	act := GlenDimplexMapping.NewMeasurement()
	pre := GlenDimplexMapping.NewMeasurement()

	interval := 10 * time.Second
	tick := time.NewTicker(interval)
	first := true

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		if err = bus.Observe(act.Map); err != nil {
			return err
		}
		if first {
			for k := range act.Map {
				k := k
				vmetrics.NewGauge(fmt.Sprintf("modbus_aqua11c_analogue{name=%q}", k),
					func() float64 {
						mu.Lock()
						v := float64(act.Map[k]) / 10.0
						mu.Unlock()
						return v
					})
			}
		}

		if len(act.Ints) > 0 {
			if err = bus.Integers(act.Ints, 0); err != nil {
				return err
			}
			if first {
				for i := range act.Ints {
					i := i
					vmetrics.NewGauge(
						fmt.Sprintf("modbus_aqua11c_integer{index=\"i%03d\"}", i),
						func() float64 {
							mu.Lock()
							v := act.Ints[i]
							mu.Unlock()
							return float64(v)
						})
				}
			}
		}

		if len(act.Bits) > 0 {
			if err = bus.Bits(act.Bits); err != nil {
				return err
			}
			if first {
				for i := range act.Bits {
					i := i
					vmetrics.NewGauge(
						fmt.Sprintf("modbus_aqua11c_bit{index=\"b%03d\"}", i),
						func() float64 {
							var j float64
							mu.Lock()
							b := act.Bits[i]
							mu.Unlock()
							if b {
								j = 1
							}
							return j
						})
				}
			}
		}

		if i := pre.Ints.DiffIndex(act.Ints); first || i >= 0 {
			log.Printf("Ints[%d]: %v", i, act.Ints)
		}
		if i := pre.Bits.DiffIndex(act.Bits); first || i >= 0 {
			log.Printf("Bits[%02d]: %v", i, act.Bits)
			if f.AlertTo != "" {
				var alert []string
				for _, ab := range bus.AlertBits {
					for j := i; j < len(act.Bits); j++ {
						if j == int(ab) && act.Bits[j] {
							alert = append(alert, bus.PCOType.Bits[ab])
						}
					}
				}
				if len(alert) != 0 {
					if err = sendAlert(f.AlertTo, alert); err != nil {
						log.Printf("alert to %q: %+v", f.AlertTo, alert)
					}
				}
			}
		}
		if k := pre.Map.DiffIndex(act.Map, 3); first || k != "" {
			log.Printf("Map[%q]=%d: %v", k, act.Map[k], act.Map)
		}
		first = false

		mu.Lock()
		pre, act = act, pre
		mu.Unlock()
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-tick.C:
		}
	}
}

func sendAlert(to string, alerts []string) error {
	var buf bytes.Buffer
	buf.WriteString("Subject: ALERT\r\n\r\n")
	for _, alert := range alerts {
		buf.WriteString(alert)
		buf.WriteString("\r\n")
	}
	log.Printf("connecting to %q", hostname)
	return smtp.SendMail(hostname+":25", nil, "pcosweb-client@"+hostname, []string{to}, buf.Bytes())
}

var GlenDimplexMapping = metrics.PCOType{
	Names: map[uint16]string{
		1:  "Outside temp",
		2:  "House temp",
		3:  "Hot water temp",
		5:  "Flow (in) temp",
		8:  "High pressure sensor (bar)",
		29: "Heating Setpoint",
		58: "Hot water Setpoint",
		96: "Heating Power Level (unsure)",
		71: "Additional Pump (operating hours)",
		72: "Compressor 1 (operating hours)",
		73: "Compressor 2 (operating hours)",
		74: "Fan (operating hours)",
		76: "Heating Pump (operating hours)",
		77: "Hot water Pump (operating hours)",
	},
}

func NewBus(host string, typ metrics.PCOType) (*metrics.Bus, error) {
	handler := modbus.NewTCPClientHandler(host + ":502")
	handler.Timeout = 10 * time.Second
	handler.SlaveId = 0x7F
	//handler.Logger = log.New(os.Stdout, "test: ", log.LstdFlags)
	if err := handler.Connect(); err != nil {
		return nil, err
	}
	return &metrics.Bus{Handler: handler, Client: modbus.NewClient(handler), PCOType: typ}, nil
}
