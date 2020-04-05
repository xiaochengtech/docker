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

// 网络配置
type Network struct {
	// TODO driver
	// TODO driver_opts
	// TODO attachable
	// TODO enable_ipv6
	// TODO ipam
	// TODO internal
	// TODO labels
	External bool `yaml:"external,omitempty"` // 是否是外部创建
	// TODO name
}
