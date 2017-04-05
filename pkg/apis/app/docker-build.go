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

package app

import (
	//"encoding/json"
	"net/http"
	//"strconv"

	"apiserver/pkg/api/docker-build"
	//"apiserver/pkg/resource"
	//"apiserver/pkg/resource/sync"
	r "apiserver/pkg/router"
	"apiserver/pkg/util/log"
	//"apiserver/pkg/util/parseUtil"

	//res "k8s.io/apimachinery/pkg/api/resource"
	//"k8s.io/client-go/pkg/api/v1"

	"github.com/gorilla/mux"
)

func DockerRegister(rout *mux.Router) {
	r.RegisterHttpHandler(rout, "/build", "POST", OfflineDockerBuild)
	r.RegisterHttpHandler(rout, "/build", "PUT", OnlineDockerBuild)
}

func OfflineDockerBuild(request *http.Request) (string, interface{}) {
	db:=docker_build.DockerBuildImg{}
	log.Info(db)
	return "", "offline app successed"
}

func OnlineDockerBuild(request *http.Request) (string, interface{}) {

	return "", "online app successed"
}