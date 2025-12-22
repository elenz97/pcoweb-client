package metrics

type AnalogVariables map[GlenDimplexAnalogVariableName]int16

type Measurement struct {
	AnalogVariables AnalogVariables
}

type GlenDimplexAnalogVariables struct {
	Length      uint16
	Names, Bits map[uint16]GlenDimplexAnalogVariableName
}
