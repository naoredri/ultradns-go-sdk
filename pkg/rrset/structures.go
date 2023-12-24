package rrset

import (
	"fmt"
	"net/url"

	"github.com/naoredri/ultradns-go-sdk/pkg/helper"
)

type RRSet struct {
	OwnerName string     `json:"ownerName,omitempty"`
	RRType    string     `json:"rrtype,omitempty"`
	TTL       int        `json:"ttl,omitempty"`
	RData     []string   `json:"rdata,omitempty"`
	Profile   RawProfile `json:"-"`
}

type RRSetKey struct {
	ID         string
	Owner      string
	Zone       string
	RecordType string
	PType      string
	Limit      int
	Offset     int
}

type RawProfile interface {
	SetContext()
	GetContext() string
}

type ResponseList struct {
	ZoneName   string             `json:"zoneName,omitempty"`
	QueryInfo  *helper.QueryInfo  `json:"queryInfo,omitempty"`
	ResultInfo *helper.ResultInfo `json:"resultInfo,omitempty"`
	RRSets     []*RRSet           `json:"rrSets,omitempty"`
}

func (r RRSetKey) RecordURI() string {
	r.Owner = url.PathEscape(r.Owner)
	r.Zone = url.PathEscape(r.Zone)

	if r.RecordType == "" {
		r.RecordType = "ANY"
	}

	if r.Limit == 0 {
		r.Limit = 100
	}

	return fmt.Sprintf("zones/%s/rrsets/%s/%s?offset=%v&limit=%v", r.Zone, r.RecordType, r.Owner, r.Offset, r.Limit)
}

func (r RRSetKey) ProbeURI() string {
	return fmt.Sprintf("%s/probes/%s", r.RecordURI(), r.ID)
}

func (r RRSetKey) ProbeListURI(query string) string {
	return fmt.Sprintf("%s/probes?q=%s", r.RecordURI(), query)
}

func (r RRSetKey) RecordID() string {
	r.Owner = helper.GetOwnerFQDN(r.Owner, r.Zone)
	r.Zone = helper.GetZoneFQDN(r.Zone)
	r.RecordType = helper.GetRecordTypeFullString(r.RecordType)

	return fmt.Sprintf("%s:%s:%s", r.Owner, r.Zone, r.RecordType)
}

func (r RRSetKey) PID() string {
	return fmt.Sprintf("%s:%s", r.RecordID(), r.ID)
}
