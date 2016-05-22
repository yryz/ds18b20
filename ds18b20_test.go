package ds18b20

import (
	"testing"
)

func TestSensors(t *testing.T) {
	sensors, err := Sensors()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("sensors: %v", sensors)
}

func TestTemperature(t *testing.T) {
	temperature, err := Temperature("28-0000010ae26f")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("temperature: %.4f°C", temperature)
}
