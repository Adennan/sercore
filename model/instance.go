package model

import (
	"encoding/json"
	"os"
	"time"

	"github.com/go-kit/kit/log"
)

type Instance struct {
	Region   string            `json:"region"`
	Env      string            `json:"env"`
	UID      string            `json:"uid"`
	Hostname string            `json:"hostname"`
	Addrs    []string          `json:"addrs"`
	Version  string            `json:"version"`
	Metadata map[string]string `json:"metadata"`
	Status   int32             `json:"status"`

	RegisterTimestamp int64 `json:"register_timestamp"`
	RunTimestamp      int64 `json:"run_timestamp"` // time of become up
	RenewTimestamp    int64 `json:"renew_timestamp"`
	RemoveTimestamp   int64 `json:"remove_timestamp"`
	LatestTimestamp   int64 `json:"latest_timestamp"`
}

func NewInstance(arg *ArgRegister) (ins *Instance) {
	t := time.Now().Local().UnixNano()
	ins = &Instance{
		Region:            arg.Region,
		Env:               arg.Env,
		UID:               arg.UID,
		Hostname:          arg.Hostname,
		Addrs:             arg.Addrs,
		Version:           arg.Version,
		RegisterTimestamp: t,
		RunTimestamp:      t,
		RenewTimestamp:    t,
		RemoveTimestamp:   t,
		LatestTimestamp:   t,
	}
	if arg.Metadata != "" {
		if err := json.Unmarshal([]byte(arg.Metadata), &ins.Metadata); err != nil {
			log.NewLogfmtLogger(os.Stderr)
		}
	}
	return
}

// InstanceInfo the info provided to consumers
type InstanceInfo struct {
	Instances map[string][]*Instance `json:"instances"`

	LatestTimestamp int64 `json:"latest_timestamp"`
}


