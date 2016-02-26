package rest

import (
	"testing"

	"github.com/pilatuz/go-devicehive"
	"github.com/stretchr/testify/assert"
)

// Test GetNetworkList and GetNetwork methods
func TestNetworkListAndGet(t *testing.T) {
	service := testNewRest(t)
	if service == nil {
		return // nothing to test
	}

	networks, err := service.GetNetworkList(0, 0, testWaitTimeout)
	assert.NoError(t, err, "Failed to get list of networks")
	assert.NotEmpty(t, networks, "No any network avaialble")
	//	for i, n := range networks {
	//		t.Logf("network-%d: %s", i, n)
	//	}

	for i, a := range networks {
		b, err := service.GetNetwork(a.ID, testWaitTimeout)
		assert.NoError(t, err, "Failed to get network")
		assert.NotNil(t, b, "No network avaialble")
		t.Logf("network-%d/A: %s", i, a)
		t.Logf("network-%d/B: %s", i, b)
		assert.JSONEq(t, toJsonStr(a), toJsonStr(b), "Networks are not the same")
	}
}

// Test UpdateNetwork method
func TestNetworkUpdate(t *testing.T) {
	return // IGNORED, DOESN'T WORK with playground

	service := testNewRest(t)
	if service == nil {
		return // nothing to test
	}

	networks, err := service.GetNetworkList(0, 0, testWaitTimeout)
	assert.NoError(t, err, "Failed to get list of networks")
	assert.NotEmpty(t, networks, "No any network avaialble")
	// t.Logf("networks: %s", networks)

	for _, a := range networks {
		a.Description += "-updated"
		err := service.UpdateNetwork(a, testWaitTimeout)
		assert.NoError(t, err, "Failed to update network")
	}
}

// Test InsertNetwork DeleteNetwork methods
func TestNetworkInsertAndDelete(t *testing.T) {
	return // IGNORED, DOESN'T WORK with playground

	service := testNewRest(t)
	if service == nil {
		return // nothing to test
	}

	network := devicehive.NewNetwork("test-net", "no-secure-key")
	err := service.InsertNetwork(network, testWaitTimeout)
	assert.NoError(t, err, "Failed to insert network")
	assert.NotEmpty(t, network.ID, "No network identifier provided")

	err = service.DeleteNetwork(network, testWaitTimeout)
	assert.NoError(t, err, "Failed to delete network")
}