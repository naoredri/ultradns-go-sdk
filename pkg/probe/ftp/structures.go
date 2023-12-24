package ftp

import "github.com/naoredri/ultradns-go-sdk/pkg/probe/helper"

type Details struct {
	Username    string             `json:"username,omitempty"`
	Password    string             `json:"password,omitempty"`
	Path        string             `json:"path,omitempty"`
	Port        int                `json:"port,omitempty"`
	PassiveMode bool               `json:"passiveMode,omitempty"`
	Limits      *helper.LimitsInfo `json:"limits,omitempty"`
}
