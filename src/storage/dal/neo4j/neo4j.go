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
	"configcenter/src/common/mapstr"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"gopkg.in/redis.v5"
	"strings"
	"time"
)

type Neo4j struct {
	Dbc neo4j.Session
	Tm  *redis.Client
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

	sessionConfig := neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}
	session, err := driver.NewSession(sessionConfig)
	if err != nil {
		return nil, err
	}

	return &Neo4j{
		Dbc: session,
		Tm:  &redis.Client{},
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

// Insert 插入数据, docs 可以为 单个数据 或者 多个数据
func (c *Neo4j) NeoInsert(dataType string, inputParam mapstr.MapStr) error {
	fmt.Println(inputParam)
	//组装一个属性字段
	param := mapToJson(inputParam)
	//去掉key的引号

	fmt.Println(param)
	//组装一个cypher
	cypher := fmt.Sprintf("CREATE (n:%s { %s } ) RETURN n", dataType, param)
	fmt.Println(cypher)
	result, err := c.Dbc.Run(cypher, nil)
	if err != nil {
		return err
	}
	for result.Next() {
		fmt.Printf("insert one data into neo4j.")
	}
	return nil
}

func mapToJson(param map[string]interface{}) string {
	dataString := ""
	for k, v := range param {
		switch t := v.(type) {
		case string:
			_ = t
			v = fmt.Sprintf(`"%s"`, v)
		case uint64:
			//do nothing
		case bool:
			v = fmt.Sprintf(`"%t"`, v)
		case time.Time:
			v = fmt.Sprintf(`"%v"`, v)
		default:
			continue
		}
		dataString = fmt.Sprintf("%s%s:%v,", dataString, k, v)
	}
	dataString = strings.TrimRight(dataString, ",")

	fmt.Println(dataString)
	//dataType , _ := json.Marshal(param)
	//dataString := string(dataType)

	return dataString
}
