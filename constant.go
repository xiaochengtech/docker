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

const (
	TagDefault = "latest" // 镜像默认Tag

	// 重启策略
	RestartNo            = "no"             // 默认重启策略，不重启容器
	RestartAlways        = "always"         // 总是重启容器
	RestartOnFailure     = "on-failure"     // 如果退出时的错误码是错误的话，重启容器
	RestartUnlessStopped = "unless-stopped" // 只有停止了要重启容器

	// 挂载卷的读写模式
	VolumeReadOnly  = "ro" // 只读模式
	VolumeReadWrite = "rw" // 读写模式
)
