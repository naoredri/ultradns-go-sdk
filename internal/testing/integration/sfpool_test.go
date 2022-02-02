package integration_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/pool"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
	"github.com/ultradns/ultradns-go-sdk/pkg/sfpool"
)

func TestSFPoolResources(t *testing.T) {
	zoneName := integration.GetRandomZoneName()

	t.Parallel()

	it := IntegrationTest{}
	ownerName := integration.GetRandomString()

	t.Run("TestCreateSFPoolResourceZone",
		func(st *testing.T) {
			it.Test = st
			it.CreatePrimaryZone(zoneName)
		})
	t.Run("TestCreateSFPoolResourceTypeAAAA",
		func(st *testing.T) {
			it.Test = st
			it.CreateSFPoolTypeAAAA(ownerName, zoneName)
		})
	t.Run("TestUpdateSFPoolResourceTypeAAAA",
		func(st *testing.T) {
			it.Test = st
			it.UpdateSFPoolTypeAAAA(ownerName, zoneName)
		})
	t.Run("TestPartialUpdateSFPoolResourceTypeAAAA",
		func(st *testing.T) {
			it.Test = st
			it.PartialUpdateSFPoolTypeAAAA(ownerName, zoneName)
		})
	t.Run("TestReadSFPoolResourceTypeAAAA",
		func(st *testing.T) {
			it.Test = st
			it.ReadSFPool(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeAAAA))
		})
	t.Run("TestDeleteSFPoolResourceTypeAAAA",
		func(st *testing.T) {
			it.Test = st
			it.DeleteSFPool(integration.GetRRSetKey(ownerName, zoneName, testRecordTypeAAAA))
		})
	t.Run("TestDeleteSFPoolResourceZone",
		func(st *testing.T) {
			it.Test = st
			it.DeleteZone(zoneName)
		})
}

func (it *IntegrationTest) CreateSFPoolTypeAAAA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeAAAA)
	rrSet := getSFPoolTypeAAAA(ownerName)
	it.CreateSFPool(rrSetKey, rrSet)
}

func (it *IntegrationTest) UpdateSFPoolTypeAAAA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeAAAA)
	rrSet := getSFPoolTypeAAAA(ownerName)
	rrSet.RData = []string{"AAAA:BBBB:CCCC:DDDD:EEEE:FFFF:1:11"}
	it.UpdateSFPool(rrSetKey, rrSet)
}

func (it *IntegrationTest) PartialUpdateSFPoolTypeAAAA(ownerName, zoneName string) {
	rrSetKey := integration.GetRRSetKey(ownerName, zoneName, testRecordTypeAAAA)
	rrSet := getSFPoolTypeAAAA(ownerName)
	rrSet.RData = []string{"AAAA:BBBB:CCCC:DDDD:EEEE:FFFF:1:12"}
	it.PartialUpdateSFPool(rrSetKey, rrSet)
}

func (it *IntegrationTest) CreateSFPool(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) {
	sfPoolService, err := sfpool.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	if _, er := sfPoolService.CreateSFPool(rrSetKey, rrSet); er != nil {
		it.Test.Fatal(er)
	}
}

func (it *IntegrationTest) UpdateSFPool(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) {
	sfPoolService, err := sfpool.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	if _, er := sfPoolService.UpdateSFPool(rrSetKey, rrSet); er != nil {
		it.Test.Fatal(er)
	}
}

func (it *IntegrationTest) PartialUpdateSFPool(rrSetKey *rrset.RRSetKey, rrSet *rrset.RRSet) {
	sfPoolService, err := sfpool.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	if _, er := sfPoolService.PartialUpdateSFPool(rrSetKey, rrSet); er != nil {
		it.Test.Fatal(er)
	}
}

func (it *IntegrationTest) ReadSFPool(rrSetKey *rrset.RRSetKey) {
	sfPoolService, err := sfpool.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	if _, _, er := sfPoolService.ReadSFPool(rrSetKey); er != nil {
		it.Test.Fatal(er)
	}
}

func (it *IntegrationTest) DeleteSFPool(rrSetKey *rrset.RRSetKey) {
	sfPoolService, err := sfpool.Get(integration.TestClient)

	if err != nil {
		it.Test.Fatal(err)
	}

	if _, er := sfPoolService.DeleteSFPool(rrSetKey); er != nil {
		it.Test.Fatal(er)
	}
}

func getSFPoolTypeAAAA(ownerName string) *rrset.RRSet {
	backupRecord := &pool.BackupRecord{
		RData: "AAAA:BBBB:CCCC:DDDD:EEEE:FFFF:1:2",
	}
	monitor := &pool.Monitor{
		Method: "GET",
		URL:    integration.TestHost,
	}
	profile := &sfpool.Profile{
		BackupRecord:             backupRecord,
		Monitor:                  monitor,
		RegionFailureSensitivity: "HIGH",
	}

	return &rrset.RRSet{
		OwnerName: ownerName,
		RRType:    testRecordTypeAAAA,
		RData:     []string{"AAAA:BBBB:CCCC:DDDD:EEEE:FFFF:1:1"},
		Profile:   profile,
	}
}
