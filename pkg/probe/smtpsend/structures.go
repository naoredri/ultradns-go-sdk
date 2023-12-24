package smtpsend

import "github.com/naoredri/ultradns-go-sdk/pkg/probe/helper"

type Details struct {
	Message string             `json:"message,omitempty"`
	From    string             `json:"from,omitempty"`
	To      string             `json:"to,omitempty"`
	Port    int                `json:"port,omitempty"`
	Limits  *helper.LimitsInfo `json:"limits,omitempty"`
}
