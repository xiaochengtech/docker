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
)

// 密钥(Short Syntax)
type SecretSimple struct {
	Source string // 名称
}

// 新建一个密钥
func NewSecretSimple(source string) SecretSimple {
	return SecretSimple{
		Source: source,
	}
}

// 实现公共接口
func (SecretSimple) IsSecret() bool {
	return true
}

func (m SecretSimple) MarshalYAML() (result interface{}, err error) {
	if len(m.Source) == 0 {
		err = errors.New("docker: simple-secret source can not be empty")
		return
	}
	result = m.Source
	return
}

func (m *SecretSimple) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var origin string
	if err = unmarshal(&origin); err != nil {
		return
	}
	m.Source = origin
	// 校验
	if len(m.Source) == 0 {
		err = errors.New("docker: simple-secret format error")
		return
	}
	return
}
