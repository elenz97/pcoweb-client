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
	"flag"
	"net/http"
	"time"

	"github.com/elenz97/dimplex-pcoweb-exporter/pkg/metrics"
	"github.com/elenz97/dimplex-pcoweb-exporter/pkg/modbus"
	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Config struct {
	Host        string
	MetricsAddr string
}

func (c *Config) Bind(fs *flag.FlagSet) {
	fs.StringVar(&c.Host, "host", "", "Modbus host")
	fs.StringVar(&c.MetricsAddr, "metrics-addr", ":9112", "Metrics address")
}

func main() {
	var cfg Config
	cfg.Bind(flag.CommandLine)
	flag.Parse()

	metricsAddr := ":9112"
	if cfg.MetricsAddr != "" {
		metricsAddr = cfg.MetricsAddr
	}

	reg := prom.NewRegistry()
	m := metrics.NewMetrics(reg)

	bus, err := modbus.NewBus(cfg.Host, metrics.GlenDimplexAnalogVariablesMapping)
	if err != nil {
		panic(err)
	}

	defer func(bus *modbus.Bus) {
		err := bus.Close()
		if err != nil {
			panic(err)
		}
	}(bus)

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))

	go func() {
		if err := http.ListenAndServe(metricsAddr, nil); err != nil {
			panic(err)
		}
	}()

	measurement := metrics.GlenDimplexAnalogVariablesMapping.NewMeasurement()

	interval := 10 * time.Second
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	if err = bus.Observe(measurement.AnalogVariables); err != nil {
		panic(err)
	}

	for analogVar := range measurement.AnalogVariables {
		metrics.RecordAnalogMetrics(analogVar, m, measurement)
	}

	for range ticker.C {
		if err = bus.Observe(measurement.AnalogVariables); err != nil {
			panic(err)
		}

		for analogVar := range measurement.AnalogVariables {
			metrics.RecordAnalogMetrics(analogVar, m, measurement)
		}
	}
}
