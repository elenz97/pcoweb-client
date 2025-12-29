package metrics

import (
	prom "github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	OutsideTemperature           prom.Gauge
	ReturnTemperature            prom.Gauge
	HotWaterTemperature          prom.Gauge
	FlowInTemperature            prom.Gauge
	HighPressureSensorBar        prom.Gauge
	HeatingSetpoint              prom.Gauge
	HeatingGoalTemperature       prom.Gauge
	HotWaterSetpoint             prom.Gauge
	HeatingPowerLevel            prom.Gauge
	AdditionalPumpOperatingHours prom.Gauge
	Compressor1OperatingHours    prom.Gauge
	Compressor2OperatingHours    prom.Gauge
	FanOperatingHours            prom.Gauge
	HeatingPumpOperatingHours    prom.Gauge
	HotWaterPumpOperatingHours   prom.Gauge
	LowPressureSensorBar         prom.Gauge
	HeatingSetpoint2             prom.Gauge
	HeatingSetpoint3             prom.Gauge
	SecondHeaterOperatingHours   prom.Gauge
	FlangeHeaterOperatingHours   prom.Gauge
}

func (vars GlenDimplexAnalogVariables) NewMeasurement() *Measurement {
	return &Measurement{
		AnalogVariables: make(AnalogVariables, len(vars.Names)),
	}
}

type GlenDimplexAnalogVariableName string

const (
	GlenDimplexAnalogVariableOutsideTemperature           GlenDimplexAnalogVariableName = "OutsideTemperature"
	GlenDimplexAnalogVariableReturnTemperature            GlenDimplexAnalogVariableName = "ReturnTemperature"
	GlenDimplexAnalogVariableHotWaterTemperature          GlenDimplexAnalogVariableName = "HotWaterTemperature"
	GlenDimplexAnalogVariableFlowInTemperature            GlenDimplexAnalogVariableName = "FlowInTemperature"
	GlenDimplexAnalogVariableHighPressureSensorBar        GlenDimplexAnalogVariableName = "HighPressureSensorBar"
	GlenDimplexAnalogVariableHeatingSetpoint              GlenDimplexAnalogVariableName = "HeatingSetpoint"
	GlenDimplexAnalogVariableHeatingSetpoint2             GlenDimplexAnalogVariableName = "HeatingSetpoint2"
	GlenDimplexAnalogVariableHeatingSetpoint3             GlenDimplexAnalogVariableName = "HeatingSetpoint3"
	GlenDimplexAnalogVariableHeatingGoalTemperature       GlenDimplexAnalogVariableName = "HeatingGoalTemperature"
	GlenDimplexAnalogVariableHotWaterSetpoint             GlenDimplexAnalogVariableName = "HotWaterSetpoint"
	GlenDimplexAnalogVariableHeatingPowerLevel            GlenDimplexAnalogVariableName = "HeatingPowerLevel"
	GlenDimplexAnalogVariableAdditionalPumpOperatingHours GlenDimplexAnalogVariableName = "AdditionalPumpOperatingHours"
	GlenDimplexAnalogVariableCompressor1OperatingHours    GlenDimplexAnalogVariableName = "Compressor1OperatingHours"
	GlenDimplexAnalogVariableCompressor2OperatingHours    GlenDimplexAnalogVariableName = "Compressor2OperatingHours"
	GlenDimplexAnalogVariableFanOperatingHours            GlenDimplexAnalogVariableName = "FanOperatingHours"
	GlenDimplexAnalogVariableHeatingPumpOperatingHours    GlenDimplexAnalogVariableName = "HeatingPumpOperatingHours"
	GlenDimplexAnalogVariableHotWaterPumpOperatingHours   GlenDimplexAnalogVariableName = "HotWaterPumpOperatingHours"
	GlenDimplexAnalogVariableLowPressureSensorBar         GlenDimplexAnalogVariableName = "LowPressureSensorBar"
	GlenDimplexAnalogVariableSecondHeaterOperatingHours   GlenDimplexAnalogVariableName = "SecondHeaterOperatingHours"
	GlenDimplexAnalogVariableFlangeHeaterOperatingHours   GlenDimplexAnalogVariableName = "FlangeHeaterOperatingHours"
)

var GlenDimplexAnalogVariablesMapping = GlenDimplexAnalogVariables{
	Names: map[uint16]GlenDimplexAnalogVariableName{
		// Service Data / Betriebsdaten https://dimplex.atlassian.net/wiki/x/zYHox
		// (R1) Außentemperatur / Outside temperature
		1: GlenDimplexAnalogVariableOutsideTemperature,
		// (R2) Temperatur Ruecklauf / Return temperature
		2: GlenDimplexAnalogVariableReturnTemperature,
		// (R3) Temperatur Warmwasser / Hot water temperature
		3: GlenDimplexAnalogVariableHotWaterTemperature,
		// (R9) Temperatur Vorlauf / Flow in temperature
		5:  GlenDimplexAnalogVariableFlowInTemperature,
		8:  GlenDimplexAnalogVariableHighPressureSensorBar,
		29: GlenDimplexAnalogVariableHeatingSetpoint,
		53: GlenDimplexAnalogVariableHeatingGoalTemperature,
		// Solltemperatur 2.Heizkreis / Desired temperature 2nd heating circuit
		54: GlenDimplexAnalogVariableHeatingSetpoint2,
		// Solltemperatur 3.Heizkreis / Desired temperature 3rd heating circuit
		55: GlenDimplexAnalogVariableHeatingSetpoint3,
		// Temperatur Warmwassersoll
		58: GlenDimplexAnalogVariableHotWaterSetpoint,
		96: GlenDimplexAnalogVariableHeatingPowerLevel,
		// Operating hours / Betriebsstunden
		// https://dimplex.atlassian.net/wiki/x/AYAqxw
		// (M16) Zusatzumwälzpumpe / Additional water pump
		71: GlenDimplexAnalogVariableAdditionalPumpOperatingHours,
		// Verdichter 1 / Compressor 1
		72: GlenDimplexAnalogVariableCompressor1OperatingHours,
		// Verdichter 2 / Compressor 2
		73: GlenDimplexAnalogVariableCompressor2OperatingHours,
		// (M11) Primärpumpe - Ventilator / Primary pump - fan
		74: GlenDimplexAnalogVariableFanOperatingHours,
		// (E10) 2. Wärmeerzeuger / 2nd heater
		75: GlenDimplexAnalogVariableSecondHeaterOperatingHours,
		// (M13) Heizungspumpe / Heat pump
		76: GlenDimplexAnalogVariableHeatingPumpOperatingHours,
		// (M18) Warmwasserpumpe / Hot water pump
		77: GlenDimplexAnalogVariableHotWaterPumpOperatingHours,
		// (E9) Flanschheizung / Flange heater
		78:  GlenDimplexAnalogVariableFlangeHeaterOperatingHours,
		101: GlenDimplexAnalogVariableLowPressureSensorBar,
	},
}

func NewMetrics(reg prom.Registerer) *Metrics {
	subsystemAnalog := "analog"
	namespace := "glendimplex"
	outsideTemperature := prom.NewGauge(prom.GaugeOpts{
		Name:      "outside_temperature",
		Namespace: namespace,
		Subsystem: subsystemAnalog,
		Help:      "Outside temperature in Celsius",
		ConstLabels: prom.Labels{
			"domain": "sensors",
		},
	})

	returnTemperature := prom.NewGauge(prom.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystemAnalog,
		Name:      "return_temperature",
		Help:      "Return temperature in Celsius",
		ConstLabels: prom.Labels{
			"domain": "sensors",
		},
	})

	hotWaterTemperature := prom.NewGauge(prom.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystemAnalog,
		Name:      "hot_water_temperature",
		Help:      "Hot water temperature in Celsius",
		ConstLabels: prom.Labels{
			"domain": "sensors",
		},
	})

	flowInTemperature := prom.NewGauge(prom.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystemAnalog,
		Name:      "flow_in_temperature",
		Help:      "Flow in temperature in Celsius",
		ConstLabels: prom.Labels{
			"domain": "sensors",
		},
	})

	highPressureSensorBar := prom.NewGauge(prom.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystemAnalog,
		Name:      "high_pressure_sensor_bar",
		Help:      "High pressure sensor bar in bar",
		ConstLabels: prom.Labels{
			"domain": "pressures",
		},
	})

	heatingSetpoint := prom.NewGauge(prom.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystemAnalog,
		Name:      "heating_setpoint",
		Help:      "Heating setpoint in Celsius",
		ConstLabels: prom.Labels{
			"domain": "heating",
		},
	})

	heatingGoalTemperature := prom.NewGauge(prom.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystemAnalog,
		Name:      "heating_goal_temperature",
		Help:      "Heating goal temperature in Celsius",
		ConstLabels: prom.Labels{
			"domain": "heating",
		},
	})

	hotWaterSetpoint := prom.NewGauge(prom.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystemAnalog,
		Name:      "hot_water_setpoint",
		Help:      "Hot water setpoint in Celsius",
		ConstLabels: prom.Labels{
			"domain": "heating",
		},
	})

	heatingPowerLevel := prom.NewGauge(prom.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystemAnalog,
		Name:      "heating_power_level",
		Help:      "Heating power level in percent",
		ConstLabels: prom.Labels{
			"domain": "heating",
		},
	})

	additionalPumpOperatingHours := prom.NewGauge(prom.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystemAnalog,
		Name:      "additional_pump_operating_hours",
		Help:      "Additional pump operating hours in hours",
		ConstLabels: prom.Labels{
			"domain": "operational",
		},
	})

	compressor1OperatingHours := prom.NewGauge(prom.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystemAnalog,
		Name:      "compressor1_operating_hours",
		Help:      "Compressor 1 operating hours in hours",
		ConstLabels: prom.Labels{
			"domain": "operational",
		},
	})

	compressor2OperatingHours := prom.NewGauge(prom.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystemAnalog,
		Name:      "compressor2_operating_hours",
		Help:      "Compressor 2 operating hours in hours",
		ConstLabels: prom.Labels{
			"domain": "operational",
		},
	})

	fanOperatingHours := prom.NewGauge(prom.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystemAnalog,
		Name:      "fan_operating_hours",
		Help:      "Fan operating hours in hours",
		ConstLabels: prom.Labels{
			"domain": "operational",
		},
	})

	heatingPumpOperatingHours := prom.NewGauge(prom.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystemAnalog,
		Name:      "heating_pump_operating_hours",
		Help:      "Heating pump operating hours in hours",
		ConstLabels: prom.Labels{
			"domain": "operational",
		},
	})

	hotWaterPumpOperatingHours := prom.NewGauge(prom.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystemAnalog,
		Name:      "hot_water_pump_operating_hours",
		Help:      "Hot water pump operating hours in hours",
		ConstLabels: prom.Labels{
			"domain": "operational",
		},
	})

	lowPressureSensorBar := prom.NewGauge(prom.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystemAnalog,
		Name:      "low_pressure_sensor_bar",
		Help:      "Low pressure sensor bar in bar",
		ConstLabels: prom.Labels{
			"domain": "pressures",
		},
	})

	heatingSetpoint2 := prom.NewGauge(prom.GaugeOpts{
		Namespace:   namespace,
		Subsystem:   subsystemAnalog,
		Name:        "heating_setpoint2",
		Help:        "Heating setpoint 2 in degrees Celsius",
		ConstLabels: prom.Labels{},
	})

	heatingSetpoint3 := prom.NewGauge(prom.GaugeOpts{
		Namespace:   namespace,
		Subsystem:   subsystemAnalog,
		Name:        "heating_setpoint3",
		Help:        "Heating setpoint 3 in degrees Celsius",
		ConstLabels: prom.Labels{},
	})

	secondHeaterOperatingHours := prom.NewGauge(prom.GaugeOpts{
		Namespace:   namespace,
		Subsystem:   subsystemAnalog,
		Name:        "second_heater_operating_hours",
		Help:        "Second heater operating hours in hours",
		ConstLabels: prom.Labels{},
	})

	flangeHeaterOperatingHours := prom.NewGauge(prom.GaugeOpts{
		Namespace:   namespace,
		Subsystem:   subsystemAnalog,
		Name:        "flange_heater_operating_hours",
		Help:        "Flange heater operating hours in hours",
		ConstLabels: prom.Labels{},
	})

	reg.MustRegister(outsideTemperature,
		returnTemperature,
		hotWaterTemperature,
		flowInTemperature,
		highPressureSensorBar,
		heatingSetpoint,
		heatingGoalTemperature,
		hotWaterSetpoint,
		heatingPowerLevel,
		additionalPumpOperatingHours,
		compressor1OperatingHours,
		compressor2OperatingHours,
		fanOperatingHours,
		heatingPumpOperatingHours,
		hotWaterPumpOperatingHours,
		lowPressureSensorBar,
		heatingSetpoint2,
		heatingSetpoint3,
		secondHeaterOperatingHours,
		flangeHeaterOperatingHours)

	return &Metrics{
		OutsideTemperature:           outsideTemperature,
		ReturnTemperature:            returnTemperature,
		HotWaterTemperature:          hotWaterTemperature,
		FlowInTemperature:            flowInTemperature,
		HighPressureSensorBar:        highPressureSensorBar,
		HeatingSetpoint:              heatingSetpoint,
		HeatingGoalTemperature:       heatingGoalTemperature,
		HotWaterSetpoint:             hotWaterSetpoint,
		HeatingPowerLevel:            heatingPowerLevel,
		AdditionalPumpOperatingHours: additionalPumpOperatingHours,
		Compressor1OperatingHours:    compressor1OperatingHours,
		Compressor2OperatingHours:    compressor2OperatingHours,
		FanOperatingHours:            fanOperatingHours,
		HeatingPumpOperatingHours:    heatingPumpOperatingHours,
		HotWaterPumpOperatingHours:   hotWaterPumpOperatingHours,
		LowPressureSensorBar:         lowPressureSensorBar,
		HeatingSetpoint2:             heatingSetpoint2,
		HeatingSetpoint3:             heatingSetpoint3,
		SecondHeaterOperatingHours:   secondHeaterOperatingHours,
		FlangeHeaterOperatingHours:   flangeHeaterOperatingHours,
	}
}

func RecordAnalogMetrics(analogVar GlenDimplexAnalogVariableName, m *Metrics, measurement *Measurement) {
	switch analogVar {
	case GlenDimplexAnalogVariableHeatingSetpoint2:
		m.HeatingSetpoint2.Set(float64(measurement.AnalogVariables[analogVar]) / 10.0)
	case GlenDimplexAnalogVariableHeatingSetpoint3:
		m.HeatingSetpoint3.Set(float64(measurement.AnalogVariables[analogVar]) / 10.0)
	case GlenDimplexAnalogVariableSecondHeaterOperatingHours:
		m.SecondHeaterOperatingHours.Set(float64(measurement.AnalogVariables[analogVar]))
	case GlenDimplexAnalogVariableFlangeHeaterOperatingHours:
		m.FlangeHeaterOperatingHours.Set(float64(measurement.AnalogVariables[analogVar]))
	case GlenDimplexAnalogVariableOutsideTemperature:
		m.OutsideTemperature.Set(float64(measurement.AnalogVariables[analogVar]) / 10.0)
	case GlenDimplexAnalogVariableReturnTemperature:
		m.ReturnTemperature.Set(float64(measurement.AnalogVariables[analogVar]) / 10.0)
	case GlenDimplexAnalogVariableHotWaterTemperature:
		m.HotWaterTemperature.Set(float64(measurement.AnalogVariables[analogVar]) / 10.0)
	case GlenDimplexAnalogVariableFlowInTemperature:
		m.FlowInTemperature.Set(float64(measurement.AnalogVariables[analogVar]) / 10.0)
	case GlenDimplexAnalogVariableHighPressureSensorBar:
		m.HighPressureSensorBar.Set(float64(measurement.AnalogVariables[analogVar]) / 10.0)
	case GlenDimplexAnalogVariableHeatingSetpoint:
		m.HeatingSetpoint.Set(float64(measurement.AnalogVariables[analogVar]) / 10.0)
	case GlenDimplexAnalogVariableHeatingGoalTemperature:
		m.HeatingGoalTemperature.Set(float64(measurement.AnalogVariables[analogVar]) / 10.0)
	case GlenDimplexAnalogVariableHotWaterSetpoint:
		m.HotWaterSetpoint.Set(float64(measurement.AnalogVariables[analogVar]) / 10.0)
	case GlenDimplexAnalogVariableHeatingPowerLevel:
		m.HeatingPowerLevel.Set(float64(measurement.AnalogVariables[analogVar]) / 10.0)
	case GlenDimplexAnalogVariableAdditionalPumpOperatingHours:
		m.AdditionalPumpOperatingHours.Set(float64(measurement.AnalogVariables[analogVar]))
	case GlenDimplexAnalogVariableCompressor1OperatingHours:
		m.Compressor1OperatingHours.Set(float64(measurement.AnalogVariables[analogVar]))
	case GlenDimplexAnalogVariableCompressor2OperatingHours:
		m.Compressor2OperatingHours.Set(float64(measurement.AnalogVariables[analogVar]))
	case GlenDimplexAnalogVariableFanOperatingHours:
		m.FanOperatingHours.Set(float64(measurement.AnalogVariables[analogVar]))
	case GlenDimplexAnalogVariableHeatingPumpOperatingHours:
		m.HeatingPumpOperatingHours.Set(float64(measurement.AnalogVariables[analogVar]))
	case GlenDimplexAnalogVariableHotWaterPumpOperatingHours:
		m.HotWaterPumpOperatingHours.Set(float64(measurement.AnalogVariables[analogVar]))
	case GlenDimplexAnalogVariableLowPressureSensorBar:
		m.LowPressureSensorBar.Set(float64(measurement.AnalogVariables[analogVar]) / 10.0)
	}
}
