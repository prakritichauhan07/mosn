/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package tcpproxy

import (
	"context"

	"mosn.io/api"
	"mosn.io/mosn/pkg/api/v2"
	"mosn.io/mosn/pkg/config"
)

func init() {
	api.RegisterNetwork(v2.TCP_PROXY, CreateTCPProxyFactory)
}

type tcpProxyFilterConfigFactory struct {
	Proxy *v2.TCPProxy
}

func (f *tcpProxyFilterConfigFactory) CreateFilterChain(context context.Context, callbacks api.NetWorkFilterChainFactoryCallbacks) {
	rf := NewProxy(context, f.Proxy)
	callbacks.AddReadFilter(rf)
}

func CreateTCPProxyFactory(conf map[string]interface{}) (api.NetworkFilterChainFactory, error) {
	p, err := config.ParseTCPProxy(conf)
	if err != nil {
		return nil, err
	}
	return &tcpProxyFilterConfigFactory{
		Proxy: p,
	}, nil
}
