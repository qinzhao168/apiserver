// Copyright © 2017 huang jia <449264675@qq.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"apiserver/pkg/configz"
	"apiserver/pkg/util/log"
	"github.com/docker/docker/client"
	"net/http"
)

var (
	DockerClient *client.Client
	err_docker      error
)

//init create client of k8s's apiserver
func init() {

	var http_client *http.Client
	dockerHost:=configz.GetString("docker","dockerHost","127.0.0.1:5555")
	dockerVersion:=configz.GetString("docker","dockerVersion","1.12.3")
	httpHeader:=make(map[string]string)
	DockerClient,err_docker=client.NewClient(dockerHost,dockerVersion,http_client,httpHeader)
	if err_docker!=nil{
		log.Fatalf("init docker client err: %v", err_docker)
	}
}
