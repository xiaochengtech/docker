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

// 端口
type Port struct {
	Target    string `yaml:"target"`             // 容器内部端口号
	Published string `yaml:"published"`          // 暴漏的端口号
	Protocol  string `yaml:"protocol,omitempty"` // 传输协议
	// TODO mode
}
