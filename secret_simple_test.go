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
	"testing"
)

func TestSecretSimple(t *testing.T) {
	tests := []struct {
		item       string
		wantSource string
		wantErr    bool
	}{
		{item: "my_secret", wantSource: "my_secret", wantErr: false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			// MarshalYaml
			if !tt.wantErr {
				item := SecretSimple{Source: tt.wantSource}
				result, _ := item.MarshalYAML()
				content := fmt.Sprintf("%s", result)
				if content != tt.item {
					t.Logf("%d %d", len(content), len(tt.item))
					t.Errorf("SecretSimple.MarshalYAML() content = %v, wantContent %v", content, tt.item)
					return
				}
			}
			// UnmarshalYaml
			var item SecretSimple
			err := UnmarshalYaml(tt.item, &item)
			if (err != nil) != tt.wantErr {
				t.Errorf("SecretSimple.UnarshalYAML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if item.Source != tt.wantSource {
				t.Errorf("SecretSimple.UnarshalYAML() source = %v, wantSouce %v", item.Source, tt.wantSource)
				return
			}
		})
	}
}
