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

const (
	// UserAgent client user agent
	UserAgent = "go-axios/0.1.0"

	headerUserAgent       = "User-Agent"
	headerAcceptEncoding  = "Accept-Encoding"
	headerContentEncoding = "Content-Encoding"
	headerContentLength   = "Content-Length"
	headerContentType     = "Content-Type"

	defaultAcceptEncoding = "gzip, br"
	brEncoding            = "br"
	gzipEncoding          = "gzip"

	contentTypeWWWFormUrlencoded = "application/x-www-form-urlencoded;charset=utf-8"
	contentTypeJSON              = "application/json;charset=utf-8"
)
