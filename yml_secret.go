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

// 暴露的映射的公共接口
type YmlSecret struct {
	File     string `yaml:"file,omitempty"`     // 文件路径
	External string `yaml:"external,omitempty"` // 是否已存在，存在不需要再创建。
	Name     string `yaml:"name,omitempty"`     // (v3.5+) 名称
}
