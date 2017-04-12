//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2017] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package types

type NodeList []Node

type Node struct {
	// Node metadata
	Meta NodeMeta `json:"meta"`
	// Node spec info
	Spec NodeSpec `json:"spec"`
	// Node state info
	State NodeState `json:"state"`
}

type NodeMeta struct {
	Meta

	Hostname     string `json:"hostname"`
	OSName       string `json:"os_name"`
	OSType       string `json:"os_type"`
	Architecture string `json:"architecture"`

	CRI     PodCRIMeta `json:"cri"`
	Network PodNetwork `json:"network"`
}

type PodCRIMeta struct {
	Type    string `json:"type"`
	Version string `json:"version"`
}

type PodNetwork struct {
	Interface string   `json:"interface,omitempty"`
	IP        []string `json:"ip,omitempty"`
}

type NodeState struct {
	// Node Capacity
	Capacity NodeResources `json:"capacity"`
	// Node Allocated
	Allocated NodeResources `json:"allocated"`
}

type NodeResources struct {
	// Node total containers
	Containers int `json:"containers"`
	// Node total pods
	Pods int `json:"pods"`
	// Node total memory
	Memory string `json:"memory"`
	// Node total cpu
	Cpu int `json:"cpu"`
	// Node storage
	Storage int `json:"storage"`
}

type NodeSpec struct {
	// Pod spec for node
	Pods []*PodNodeSpec `json:"pods"`
}
