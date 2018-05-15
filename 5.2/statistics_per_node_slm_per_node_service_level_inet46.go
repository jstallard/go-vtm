// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 5.2.
package vtm

import (
	"encoding/json"
)

type PerNodeSlmPerNodeServiceLevelInet46Statistics struct {
	Statistics struct {
		ResponseMean *int `json:"response_mean"`
		NodePort     *int `json:"node_port"`
		TotalNonConf *int `json:"total_non_conf"`
		ResponseMin  *int `json:"response_min"`
		ResponseMax  *int `json:"response_max"`
		TotalConn    *int `json:"total_conn"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetPerNodeSlmPerNodeServiceLevelInet46Statistics(name string) (*PerNodeSlmPerNodeServiceLevelInet46Statistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/5.2/status/local_tm/statistics/per_node_slm/per_node_service_level_inet46/" + name)
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(PerNodeSlmPerNodeServiceLevelInet46Statistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}