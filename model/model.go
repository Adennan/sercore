package model

// Define register renew cancel fetch poll set action

type ArgRegister struct {
	Region      string   `json:"region"`
	Env         string   `json:"env"`
	UID         string   `json:"uid"`
	Hostname    string   `json:"hostname"`
	Status      int32    `json:"status"`
	Addrs       []string `json:"addrs"`
	Version     string   `json:"version"`
	Metadata    string   `json:"metadata"`
	Replication bool     `json:"replication"`
}
