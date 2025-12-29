# pcoweb-client

Prometheus exporter for Glen Dimplex heat pumps (heavily modified fork of https://github.com/tgulacsi/pcoweb-client)

## Installation

```shell
$ make help

Usage:
  make <target>

Help
  help             Display this help.

Build
  build-linux      Build binary for Linux
  build-armv7      Build binary for ARMv7 architecture
  build-arm64      Build binary for ARM64 architecture
```

## Usage

```shell
Usage of pcoweb-client:
      --host string           modbus host to connect to
      --metrics-port string   metrics address to listen on (default ":9112")
```

## Metrics

Below values are (roughly mapped, errors not to be ruled out) taken from `https://pcoweb-client-address/config/adminpage.html` and `https://pcoweb-client-address/http/`.
For now, only analog variables are read from the client.

### Analog Variables

<details>

<summary>Analog Variables</summary>

```
| Analogue Bit | Name                              | Description |
|--------------|-----------------------------------|-------------|
| 1            | Outside Temperature               |             |
| 2            | Return / House Temperature        |             |
| 3            | Actual Hot Water                  |             |
| 5            | Flow (in)                         |             |
| 8            | High-Pressure Sensor (bar)        |             |
| 29           | Heating Setpoint                  |             |
| 53           | Heating Goal Temperature          |             |
| 58           | Hot Water Setpoint                |             |
| 71           | Additional Pump (Operating Hours) |             |
| 72           | Compressor 1 (Operating Hours)    |             |
| 73           | Compressor 2 (Operating Hours)    |             |
| 74           | Fan (Operating Hours)             |             |
| 76           | Heating Pump (Operating Hours)    |             |
| 77           | Hot Water Pump (Operating Hours)  |             |
| 96           | Heating Power Level (unsure)      |             |
| 101          | Low-Pressure Sensor (bar)         |             |
```

</details>

### Integer Variables

<details>

<summary>Integer Variables</summary>

```
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
```

</details>

### Grafana Dashboard

In the `./grafana` directory, you can find a `dashboard.json` file showing the heat pump's metrics.

![Grafana](./static/assets/grafana.png "a title")
