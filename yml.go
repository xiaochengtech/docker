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

// 完整的配置文件
type Yml struct {
	Version  string             `yaml:"version"`            // 版本号
	Services map[string]Service `yaml:"services"`           // 服务配置
	Volumes  map[string]Volume  `yaml:"volumes,omitempty"`  // 挂载卷配置
	Networks map[string]Network `yaml:"networks,omitempty"` // 网络配置
	// TODO configs
	// TODO secrets
	// TODO Variable substitution
	// TODO Extension fields
}
