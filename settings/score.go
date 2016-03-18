// Copyright © 2016 Jip J. Dekker <jip@dekker.li>
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

package settings

import (
	"encoding/json"
	"io/ioutil"
)

// Score represents the settings for a specific score file
type Score struct {
	Name       string   // The name of the score in the songbook
	Categories []string // Categories to which the scores belong
	Path       string   // The path to the scores (uncompiled) file
}

// FromJSON reads the settings of a score from a JSON file
func FromJSON(path string) (*Score, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var s Score
	err = json.Unmarshal(data, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
