# pcoweb-client

Prometheus exporter for Glen Dimplex heat pumps

## Usage

```shell
Usage of pcoweb-client:
      --host string           modbus host to connect to
      --metrics-port string   metrics address to listen on (default ":9112")
```

## Metrics

Below values are (roughly mapped, errors not to be ruled out) taken from `https://pcoweb-client-address/config/adminpage.html` and `https://pcoweb-client-address/http/`

### Analog Variables

| Analogue Bit | Name                              | Description |
|--------------|-----------------------------------|-------------|
| 1            | Outside Temperature               |             |
| 2            | House Temperature                 |             |
| 3            | Actual Hot Water                  |             |
| 5            | Flow (in)                         |             |
| 8            | High-Pressure Sensor (bar)        |             |
| 29           | Heating Setpoint                  |             |
| 53           | Heating Goal Temperature          |             |
| 58           | Hot Water Setpoint                |             |
| 96           | Heating Power Level (unsure)      |             |
| 100          | Low-Pressure Sensor (bar)         |             |
| 71           | Additional Pump (Operating Hours) |             |
| 72           | Compressor 1 (Operating Hours)    |             |
| 73           | Compressor 2 (Operating Hours)    |             |
| 74           | Fan (Operating Hours)             |             |
| 76           | Heating Pump (Operating Hours)    |             |
| 77           | Hot Water Pump (Operating Hours)  |             |

### Integer Variables

| Integer Bit | Name                             | Description |
|-------------|----------------------------------|-------------|
| 95          | Generated Heat (kWh)             |             |
| 1660        | Total Generated Heat (kWh)       |             |
| 1663        | Total Heating (kWh)              |             |
| 1675        | Heating (kWh)                    |             |
| 1681        | Hot Water (kWh)                  |             |
| 1669        | Total Hot Water (kWh)	           |             |
| 1647        | Environmental Energy (kWh)       |             |
| 1644        | Total Environmental Energy (kWh) |             |
