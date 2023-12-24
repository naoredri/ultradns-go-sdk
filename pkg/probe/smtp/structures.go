package smtp

import "github.com/naoredri/ultradns-go-sdk/pkg/probe/helper"

type Details struct {
	Port   int                `json:"port,omitempty"`
	Limits *helper.LimitsInfo `json:"limits,omitempty"`
}
