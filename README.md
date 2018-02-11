# ds18b20

Get sensor data from ds18b20 connected to the Raspberry (GPIO w1 pin).

## Usage

### Connect DS18B20
On the Raspberry Pi, you will need to add `dtoverlay=w1-gpio"` (for regular connection) or `dtoverlay=w1-gpio,pullup="y"` (for parasitic connection) to your /boot/config.txt. The default data pin is GPIO4 (RaspPi connector pin 7), but that can be changed from 4 to `x` with `dtoverlay=w1-gpio,gpiopin=x`.

Here's what I did:
```
sudo echo dtoverlay=w1-gpio-pullup,gpiopin=4 >> /boot/config.txt
sudo modprobe w1_gpio && sudo modprobe w1_therm
```

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
