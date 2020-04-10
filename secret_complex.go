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

// 密钥(Long Syntax)
type SecretComplex struct {
	Source string `yaml:"source"`           // 名称
	Target string `yaml:"target,omitempty"` // 文件名
	Uid    string `yaml:"uid,omitempty"`    // 文件UID
	Gid    string `yaml:"gid,omitempty"`    // 文件GID
	Mode   string `yaml:"mode,omitempty"`   // 文件权限
}

// 实现公共接口
func (SecretComplex) IsSecret() bool {
	return true
}
