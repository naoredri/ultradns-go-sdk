package integration_test

import (
	"testing"

	"github.com/naoredri/ultradns-go-sdk/internal/testing/integration"
	"github.com/naoredri/ultradns-go-sdk/pkg/probe"
	"github.com/naoredri/ultradns-go-sdk/pkg/probe/ftp"
	"github.com/naoredri/ultradns-go-sdk/pkg/probe/helper"
)

func (t *IntegrationTest) TestFTPProbeResources(zoneName, ownerName string) {
	it := IntegrationTest{}

	t.Test.Run("TestCreateProbeResourceTypeFTP",
		func(st *testing.T) {
			it.Test = st
			it.CreateProbeTypeFTP(ownerName, zoneName)
		})
	t.Test.Run("TestListProbeResourceTypeFTP",
		func(st *testing.T) {
			it.Test = st
			it.ListProbe(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.FTP))
		})
	t.Test.Run("TestUpdateProbeResourceTypeFTP",
		func(st *testing.T) {
			it.Test = st
			it.UpdateProbeTypeFTP(ownerName, zoneName)
		})
	t.Test.Run("TestPartialUpdateProbeResourceTypeFTP",
		func(st *testing.T) {
			it.Test = st
			it.PartialUpdateProbeTypeFTP(ownerName, zoneName)
		})
	t.Test.Run("TestReadProbeResourceTypeFTP",
		func(st *testing.T) {
			it.Test = st
			it.ReadProbe(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.FTP))
		})
	t.Test.Run("TestDeleteProbeResourceTypeFTP",
		func(st *testing.T) {
			it.Test = st
			it.DeleteProbe(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.FTP))
		})
}

func (t *IntegrationTest) CreateProbeTypeFTP(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.FTP)
	probedata := getProbeTypeFTP()
	t.CreateProbe(rrSetKey, probedata)
}

func (t *IntegrationTest) UpdateProbeTypeFTP(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.FTP)
	probedata := getProbeTypeFTP()
	probedata.Interval = testProbeInterval
	t.UpdateProbe(rrSetKey, probedata)
}

func (t *IntegrationTest) PartialUpdateProbeTypeFTP(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeA, probe.FTP)
	probedata := getProbeTypeFTP()
	limit := &helper.Limit{
		Fail: 20,
	}
	limitInfo := &helper.LimitsInfo{
		Run:     limit,
		Connect: limit,
	}
	details := &ftp.Details{
		Limits: limitInfo,
	}
	probedata.Details = details
	t.PartialUpdateProbe(rrSetKey, probedata)
}

func getProbeTypeFTP() *probe.Probe {
	limit := &helper.Limit{
		Fail: 10,
	}
	limitInfo := &helper.LimitsInfo{
		Run:     limit,
		Connect: limit,
	}
	details := &ftp.Details{
		Port:   21,
		Limits: limitInfo,
		Path:   "test",
	}
	probedata := &probe.Probe{
		Type:      probe.FTP,
		Interval:  "ONE_MINUTE",
		Agents:    []string{"NEW_YORK", "DALLAS"},
		Threshold: 2,
		Details:   details,
	}

	return probedata
}
