package ping

import "github.com/naoredri/ultradns-go-sdk/pkg/probe/helper"

type Details struct {
	Packets    int                `json:"packets,omitempty"`
	PacketSize int                `json:"packetSize,omitempty"`
	Limits     *helper.LimitsInfo `json:"limits,omitempty"`
}
