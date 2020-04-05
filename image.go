/*
   Copyright (c) 2020 XiaochengTech
   gitee.com/xiaochengtech/docker is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
       http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package docker

import (
	"errors"
	"fmt"
	"strings"
)

// 镜像
type Image struct {
	Name string
	Tag  string
}

func (m Image) MarshalYAML() (result interface{}, err error) {
	if len(m.Tag) > 0 {
		result = fmt.Sprintf("%s:%s", m.Name, m.Tag)
	} else {
		result = m.Name
	}
	return
}

func (m *Image) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var origin string
	if err = unmarshal(&origin); err != nil {
		return
	}
	parts := strings.Split(origin, ":")
	if len(parts) > 2 {
		err = errors.New("docker: image format error")
		return
	}
	m.Name = parts[0]
	if len(parts) > 1 {
		m.Tag = parts[1]
	}
	return
}
