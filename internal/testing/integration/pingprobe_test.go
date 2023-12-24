package integration_test

import (
	"testing"

	"github.com/naoredri/ultradns-go-sdk/internal/testing/integration"
	"github.com/naoredri/ultradns-go-sdk/pkg/probe"
	"github.com/naoredri/ultradns-go-sdk/pkg/probe/helper"
	"github.com/naoredri/ultradns-go-sdk/pkg/probe/ping"
)

func (t *IntegrationTest) TestPINGProbeResources(zoneName, ownerName string) {
	it := IntegrationTest{}

	t.Test.Run("TestCreateProbeResourceTypePING",
		func(st *testing.T) {
			it.Test = st
			it.CreateProbeTypePING(ownerName, zoneName)
		})
	t.Test.Run("TestListProbeResourceTypePING",
		func(st *testing.T) {
			it.Test = st
			it.ListProbe(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.PING))
		})
	t.Test.Run("TestUpdateProbeResourceTypePING",
		func(st *testing.T) {
			it.Test = st
			it.UpdateProbeTypePING(ownerName, zoneName)
		})
	t.Test.Run("TestPartialUpdateProbeResourceTypePING",
		func(st *testing.T) {
			it.Test = st
			it.PartialUpdateProbeTypePING(ownerName, zoneName)
		})
	t.Test.Run("TestReadProbeResourceTypePING",
		func(st *testing.T) {
			it.Test = st
			it.ReadProbe(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.PING))
		})
	t.Test.Run("TestDeleteProbeResourceTypePING",
		func(st *testing.T) {
			it.Test = st
			it.DeleteProbe(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.PING))
		})
}

func (t *IntegrationTest) CreateProbeTypePING(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.PING)
	probedata := getProbeTypePING()
	t.CreateProbe(rrSetKey, probedata)
}

func (t *IntegrationTest) UpdateProbeTypePING(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.PING)
	probedata := getProbeTypePING()
	probedata.Interval = testProbeInterval
	t.UpdateProbe(rrSetKey, probedata)
}

func (t *IntegrationTest) PartialUpdateProbeTypePING(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.PING)
	probedata := getProbeTypePING()
	limit := &helper.Limit{
		Fail: 20,
	}
	limitInfo := &helper.LimitsInfo{
		LossPercent: limit,
		Total:       limit,
	}
	details := &ping.Details{
		Packets:    5,
		PacketSize: 56,
		Limits:     limitInfo,
	}
	probedata.Details = details
	t.PartialUpdateProbe(rrSetKey, probedata)
}

func getProbeTypePING() *probe.Probe {
	limit := &helper.Limit{
		Fail: 10,
	}
	limitInfo := &helper.LimitsInfo{
		LossPercent: limit,
		Total:       limit,
	}
	details := &ping.Details{
		Packets:    3,
		PacketSize: 56,
		Limits:     limitInfo,
	}
	probedata := &probe.Probe{
		Type:      probe.PING,
		Interval:  "ONE_MINUTE",
		Agents:    []string{"NEW_YORK", "DALLAS"},
		Threshold: 2,
		Details:   details,
	}

	return probedata
}
