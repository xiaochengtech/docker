/*
	Copyright (c) 2020 XiaochengTech
	gitee.com/xiaochengtech/docker is licensed under the Mulan PSL v1.
	You can use this software according to the terms and conditions of the Mulan PSL v1.
	You may obtain a copy of Mulan PSL v1 at:
		http://license.coscl.org.cn/MulanPSL
	THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR
	PURPOSE.
	See the Mulan PSL v1 for more details.
*/

package docker

// 挂载卷
type VolumeMap struct {
	Type     string `yaml:"type"`                // 挂载类型
	Source   string `yaml:"source"`              // 外部的源地址
	Target   string `yaml:"target"`              // 容器内的目标地址
	ReadOnly string `yaml:"read_only,omitempty"` // 只读标志
	// TODO bind
	// TODO volume
	// TODO tmpfs
	// TODO consistency
}
