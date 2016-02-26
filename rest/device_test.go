package rest

import (
	"testing"

	"github.com/pilatuz/go-devicehive"
	"github.com/stretchr/testify/assert"
)

// Test GetDeviceList and GetDevice methods
func TestDeviceListAndGet(t *testing.T) {
	service := testNewRest(t)
	if service == nil {
		return // nothing to test
	}

	devices, err := service.GetDeviceList(0, 0, testWaitTimeout)
	assert.NoError(t, err, "Failed to get list of devices")
	// assert.NotEmpty(t, devices, "No any device avaialble")
	//	for i, d := range devices {
	//		t.Logf("device-%d: %s", i, d)
	//	}

	for i, a := range devices {
		b, err := service.GetDevice(a.ID, a.Key, testWaitTimeout)
		assert.NoError(t, err, "Failed to get device")
		assert.NotNil(t, b, "No device avaialble")
		t.Logf("device-%d/A: %s", i, a)
		t.Logf("device-%d/B: %s", i, b)
		assert.JSONEq(t, toJsonStr(a), toJsonStr(b), "Devices are not the same")
	}
}

// Test RegisterDevice and DeleteDevice methods
func TestDeviceRegisterAndDelete(t *testing.T) {
	service := testNewRest(t)
	if service == nil {
		return // nothing to test
	}

	device := devicehive.NewDevice("go-unit-test-device", "go test device",
		devicehive.NewDeviceClass("go-test-deviceclass", "0.0.1"))
	err := service.RegisterDevice(device, testWaitTimeout)
	if assert.NoError(t, err, "Failed to register device") {
		t.Logf("device registered: %s", device)

		devices, err := service.GetDeviceList(0, 0, testWaitTimeout)
		assert.NoError(t, err, "Failed to get list of devices")
		for _, d := range devices {
			if d.ID == device.ID {
				err = service.DeleteDevice(device, testWaitTimeout)
				assert.NoError(t, err, "Failed to delete device")
				return // OK
			}
		}

		assert.Fail(t, "No new device found in the device list")
	}
}