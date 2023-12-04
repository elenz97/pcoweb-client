package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

const (
	FlagHost         string = "host"
	FlagMetricsPort  string = "metrics-port"
	FlagAlertTo      string = "alert-to"
	FlagPollInterval string = "poll-interval"
)

type Flags struct {
	Host        string
	MetricsAddr string
	AlertTo     string
}

func (f *Flags) Bind(set *pflag.FlagSet) {
	set.StringVar(&f.Host, FlagHost, "", "modbus host to connect to")
	set.StringVar(&f.MetricsAddr, FlagMetricsPort, ":9112", "metrics address to listen on")
	set.StringVar(&f.AlertTo, FlagAlertTo, "", "")
}

func PrintFlagValue(flag, value string) {
	logrus.Infof("--%s=%s", flag, value)
}
