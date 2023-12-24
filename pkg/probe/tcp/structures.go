package tcp

import "github.com/naoredri/ultradns-go-sdk/pkg/probe/helper"

type Details struct {
	ControlIP string             `json:"controlIP,omitempty"`
	Port      int                `json:"port,omitempty"`
	Limits    *helper.LimitsInfo `json:"limits,omitempty"`
}
