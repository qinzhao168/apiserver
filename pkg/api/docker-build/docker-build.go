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

package docker_build

import (
	"errors"
	"apiserver/pkg/storage/mysqld"
	"apiserver/pkg/util/jsonx"
	"apiserver/pkg/util/log"
)

type DockerBuildStatus int32

const (
	Building  DockerBuildStatus = 0
	Successed DockerBuildStatus = 1
	BuildFailed    DockerBuildStatus = 2
)

//DockerBuildImg is struct of application
type DockerBuildImg struct {
	Id          string            `json:"name" xorm:"pk not null int"`
	ImageName          string            `json:"imageName" xorm:"varchar(256)"`
	Version        string            `json:"version" xorm:"varchar(11)"`
	BaseImage        string            `json:"baseImage" xorm:"varchar(256)"`
	Registry       string            `json:"registry" xorm:"varchar(1024)"`
	Image         string            `json:"image" xorm:"varchar(1024)"`
	Repository     string          `json:"repository" xorm:"varchar(1024)"`
	Status        DockerBuildStatus         `json:"status" xorm:"int(1) default(0)"` //构建中 0 成功 1 失败 2 运行中 3 停止 4 删除 5
	UserName      string            `json:"userName" xorm:"varchar(256)"`
	Remark        string            `json:"remark" xorm:"varchar(1024)"`
	Branch        string            `json:"branch" xorm:"varchar(1024)"`
}
var (
	engine = mysqld.GetEngine()
	Status = map[DockerBuildStatus]string{
		Building:  "Building",
		Successed: "Successed",
		BuildFailed:    "BuildFailed",
	}
)

func init() {
	engine.ShowSQL(true)
	if err := engine.Sync(new(DockerBuildImg)); err != nil {
		log.Fatalf("Sync DockerBuildImg fail :%s", err.Error())
	}
}

func (db *DockerBuildImg) String() string {
	dockerBuildStr := jsonx.ToJson(db)
	return dockerBuildStr
}

func (db *DockerBuildImg) Insert() error {
	_, err := engine.Insert(db)
	if err != nil {
		return err
	}
	return nil
}

func (db *DockerBuildImg) Delete() error {
	_, err := engine.Id(db.Id).Delete(db)
	if err != nil {
		return err
	}
	return nil
}

func (db *DockerBuildImg) Update() error {
	_, err := engine.Id(db.Id).Update(db)
	if err != nil {
		return err
	}
	return nil
}

func (db *DockerBuildImg) QueryOne() (*DockerBuildImg, error) {
	has, err := engine.Id(db.Id).Get(db)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("current docker build info not exsit")
	}
	return db, nil
}

func (db *DockerBuildImg) QuerySet() ([]*DockerBuildImg, error) {
	DBSet := []*DockerBuildImg{}
	err := engine.Where("1 and 1 order by id desc").Find(&DBSet)
	if err != nil {
		return nil, err
	}
	return DBSet, nil
}

