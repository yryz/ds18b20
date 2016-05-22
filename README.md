# ds18b20

Get sensor data from ds18b20 connected to the Raspberry (GPIO w1 pin).

## Usage

### Drivers

1-Wire drivers need to be loaded in order to create the connection between the physical sensor and the rPI.
You can load them from the terminal (or from the bin/modules.sh script).

    sudo modprobe wire
    sudo modprobe w1-gpio
    sudo modprobe w1-therm

### Install
    go get github.com/yryz/ds18b20

### Code
```go
package main

import (
    "fmt"

    "github.com/yryz/ds18b20"
)

func main() {
    sensors, err := ds18b20.Sensors()
    if err != nil {
        panic(err)
    }

    fmt.Printf("sensor IDs: %v\n", sensors)

    for _, sensor := range sensors {
        t, err := ds18b20.Temperature(sensor)
        if err == nil {
            fmt.Printf("sensor: %s temperature: %.2f°C\n", sensor, t)
        }
    }
}
```
