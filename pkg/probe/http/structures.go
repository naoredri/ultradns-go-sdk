package http

import (
	"github.com/naoredri/ultradns-go-sdk/pkg/probe/helper"
)

type Details struct {
	Transactions []*Transaction `json:"transactions,omitempty"`
	TotalLimits  *helper.Limit  `json:"totalLimits,omitempty"`
}

type Transaction struct {
	Method           string             `json:"method,omitempty"`
	ProtocolVersion  string             `json:"protocolVersion,omitempty"`
	URL              string             `json:"url,omitempty"`
	TransmittedData  string             `json:"transmittedData,omitempty"`
	ExpectedResponse string             `json:"expectedResponse,omitempty"`
	FollowRedirects  bool               `json:"followRedirects,omitempty"`
	Limits           *helper.LimitsInfo `json:"limits,omitempty"`
}
