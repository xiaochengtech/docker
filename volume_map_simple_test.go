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
	"fmt"
	"strings"
	"testing"
)

func TestVolumeMapSimple(t *testing.T) {
	tests := []struct {
		item          string
		wantHost      string
		wantContainer string
		wantMode      string
		wantErr       bool
	}{
		{item: "/var/lib/mysql", wantHost: "/var/lib/mysql", wantErr: false},
		{item: "/opt/data:/var/lib/mysql", wantHost: "/opt/data", wantContainer: "/var/lib/mysql", wantErr: false},
		{item: "./cache:/tmp/cache", wantHost: "./cache", wantContainer: "/tmp/cache", wantErr: false},
		{item: "~/configs:/udp", wantHost: "~/configs", wantContainer: "/udp", wantErr: false},
		{item: "~/configs:/etc/configs/:ro", wantHost: "~/configs", wantContainer: "/etc/configs/", wantMode: "ro", wantErr: false},
		{item: "datavolume:/var/lib/mysql", wantHost: "datavolume", wantContainer: "/var/lib/mysql", wantErr: false},
		{item: "datavolume::ro", wantHost: "datavolume", wantContainer: "", wantMode: "ro", wantErr: true},
		{item: ":/var/lib/mysql", wantHost: "", wantContainer: "/var/lib/mysql", wantMode: "", wantErr: true},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			// MarshalYaml
			if !tt.wantErr {
				item := VolumeMapSimple{Host: tt.wantHost, Container: tt.wantContainer, Mode: tt.wantMode}
				content := MarshalYaml(item)
				content = strings.TrimRight(content, "\n")
				if content != tt.item {
					t.Logf("%d %d", len(content), len(tt.item))
					t.Errorf("VolumeMapSimple.MarshalYAML() content = %v, wantContent %v", content, tt.item)
					return
				}
			}
			// UnmarshalYaml
			var item VolumeMapSimple
			err := UnmarshalYaml(tt.item, &item)
			if (err != nil) != tt.wantErr {
				t.Errorf("VolumeMapSimple.UnarshalYAML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if item.Host != tt.wantHost {
				t.Errorf("VolumeMapSimple.UnarshalYAML() host = %v, wantHost %v", item.Host, tt.wantHost)
				return
			}
			if item.Container != tt.wantContainer {
				t.Errorf("Image.UnarshalYAML() container = %v, wantContainer %v", item.Container, tt.wantContainer)
				return
			}
			if item.Mode != tt.wantMode {
				t.Errorf("VolumeMapSimple.UnarshalYAML() mode = %v, wantMode %v", item.Mode, tt.wantMode)
				return
			}
		})
	}
}
