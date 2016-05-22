// Copyright 2016 yryz Author. All Rights Reserved.

package ds18b20

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

// Sensors get all connected sensor IDs as array
func Sensors() ([]string, error) {
	data, err := ioutil.ReadFile("/sys/bus/w1/devices/w1_bus_master1/w1_master_slaves")
	if err != nil {
		return nil, err
	}

	sensors := strings.Split(string(data), "\n")
	if len(sensors) > 0 {
		sensors = sensors[:len(sensors)-1]
	}

	return sensors, nil
}

// Temperature get the temperature of a given sensor
func Temperature(sensor string) (float64, error) {
	data, err := ioutil.ReadFile("/sys/bus/w1/devices/" + sensor + "/w1_slave")
	if err != nil {
		return 0.0, nil
	}

	if strings.Contains(string(data), "YES") {
		arr := strings.SplitN(string(data), " ", 3)

		switch arr[1][0] {
		case 'f': //-0.5 ~ -55°C
			x, err := strconv.ParseInt(arr[1]+arr[0], 16, 32)
			if err != nil {
				return 0.0, err
			}
			return float64(^x+1) * 0.0625, nil

		case '0': //0~125°C
			x, err := strconv.ParseInt(arr[1]+arr[0], 16, 32)
			if err != nil {
				return 0.0, err
			}
			return float64(x) * 0.0625, nil
		}
	}

	return 0.0, errors.New("can not read temperature for sensor " + sensor)
}
