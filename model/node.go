package model

type NodeStatus int

const (
	NodeStatusUp NodeStatus = iota
	NodeStatusDown
)

// Node node
type Node struct {
	Addr   string     `json:"addr"`
	Status NodeStatus `json:"status"`
	Region string     `json:"region"`
}

