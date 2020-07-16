/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package neo4j

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"gopkg.in/redis.v5"
	"time"
)

type Neo4j struct {
	dbc neo4j.Session
	tm  *redis.Client
}

type NeoConf struct {
	//MaxOpenConns   uint64
	//MaxIdleConns   uint64
	Address    string
	port       string
	URI        string
	UserName   string
	Password   string
	Timeout    time.Duration
	graphName  string
	accessMode neo4j.AccessMode
}

// NewMgo returns new RDB
func NewNeo4j(config NeoConf, timeout time.Duration) (*Neo4j, error) {
	if config.URI == "" {
		config.URI = fmt.Sprintf("bolt://%s:%s", config.Address, config.port)
	}
	configForNeo4j40 := func(conf *neo4j.Config) { conf.Encrypted = false }
	driver, err := neo4j.NewDriver(config.URI, neo4j.BasicAuth(config.UserName, config.Password, ""), configForNeo4j40)
	if err != nil {
		return nil, err
	}
	//defer driver.Close()

	sessionConfig := neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: config.graphName,
	}
	session, err := driver.NewSession(sessionConfig)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	return &Neo4j{
		dbc: session,
		tm:  &redis.Client{},
	}, nil
}

// ParseConfigFromKV returns a new config
func ParseConfigFromKV(prefix string, configmap map[string]string) NeoConf {
	c := NeoConf{
		Address:   configmap[prefix+".host"],
		port:      configmap[prefix+".port"],
		UserName:  configmap[prefix+".usr"],
		Password:  configmap[prefix+".pwd"],
		graphName: configmap[prefix+".database"],
	}
	return c
}
