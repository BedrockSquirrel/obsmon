package model

import "time"

// Environment is a struct that represents one of the environments we run testnets in (e.g. dev-, uat-, sepolia-)
type Environment struct {
	ID string `json:"id"`
}

// Testnet is a struct that represents an instance of a testnet
type Testnet struct {
	ID             string      `json:"id"`
	Environment    Environment `json:"environment"`
	DeployTime     time.Time   `json:"deployTime"`
	GithubBuildNum int         `json:"githubBuildNum"`
	Nodes          []Node      `json:"nodes"`
}

type Node struct {
	ID        string `json:"id"` // this is the short name of the host that we see in datadog, e.g. T-1-354
	TestnetID string `json:"testnetId"`
	Hostname  string `json:"hostname"`
}
