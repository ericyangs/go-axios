// Copyright 2019 tree xie
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package axios

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	t.Run("json", func(t *testing.T) {
		assert := assert.New(t)
		m := make(map[string]string)
		resp := &Response{
			Data: []byte(`{"a": "b"}`),
		}
		err := resp.JSON(&m)
		assert.Nil(err)
		assert.Equal("b", m["a"])

		result := make(map[string]int)
		err = resp.JSON(&result)
		assert.NotNil(err)
	})
}
