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
	"fmt"
	"net/http"

	a "apiserver/pkg/apis/app"
	"apiserver/pkg/componentconfig"
	"apiserver/pkg/configz"
	"apiserver/pkg/util/log"

	"github.com/gorilla/mux"
)

type DockerBuild struct {
	*componentconfig.DockerBuildConfig
}

func NewDockerBuild() *DockerBuild {
	return &DockerBuild{
		DockerBuildConfig: &componentconfig.DockerBuildConfig{
			HttpAddr: configz.GetString("docker", "dockerHttpAddr", "0.0.0.0"),
			HttpPort: configz.MustInt("docker", "dockerPort", 9191),
			RpcAddr:  configz.GetString("docker", "dockerRpcAddr", "0.0.0.0"),
			RpcPort:  configz.MustInt("docker", "dockerRpcPort", 7171),
		},
	}
}

func Run(d *DockerBuild) error {
	root := mux.NewRouter()
	dockerbuild := root.PathPrefix("/docker").Subrouter()
	installDockerApiGroup(dockerbuild)
	http.Handle("/", root)
	log.Infof("starting dockerbuild and listen on : %v", fmt.Sprintf("%v:%v", d.HttpAddr, d.HttpPort))
	//go sync.Sync()
	return http.ListenAndServe(fmt.Sprintf("%v:%v", d.HttpAddr, d.HttpPort), nil)
}

func installDockerApiGroup(router *mux.Router) {
	a.DockerRegister(router)
}