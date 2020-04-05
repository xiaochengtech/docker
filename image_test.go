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

func TestImage(t *testing.T) {
	tests := []struct {
		item     string
		wantName string
		wantTag  string
		wantErr  bool
	}{
		{"testimage:mainline", "testimage", "mainline", false},
		{"testimage", "testimage", "", false},
		{"testimage:mainline:tag", "", "", true},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			// MarshalYaml
			if !tt.wantErr {
				item := Image{Name: tt.wantName, Tag: tt.wantTag}
				content := MarshalYaml(item)
				content = strings.TrimRight(content, "\n")
				if content != tt.item {
					t.Errorf("Image.MarshalYAML() content = %v, wantContent %v", content, tt.item)
					return
				}
			}
			// UnmarshalYaml
			var item Image
			err := UnmarshalYaml(tt.item, &item)
			if (err != nil) != tt.wantErr {
				t.Errorf("Image.UnarshalYAML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if item.Name != tt.wantName {
				t.Errorf("Image.UnarshalYAML() name = %v, wantName %v", item.Name, tt.wantName)
				return
			}
			if item.Tag != tt.wantTag {
				t.Errorf("Image.UnarshalYAML() tag = %v, wantTag %v", item.Tag, tt.wantTag)
				return
			}
		})
	}
}
