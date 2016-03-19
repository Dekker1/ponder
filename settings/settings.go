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

// Settings provides a structure to interact with the settings
// of a Ponder library
type Settings struct {
	Name             string   // Name of the Ponder library
	IgnoreDirs       []string // Directories to be ignored on search
	LilypondIncludes []string // Directories to be included when running the lilypond compiler
	OutputDir        string   // Directory in which all complete file are stored
}

// FromFile reads a settings file in json format and returns the Settings struct
func FromFile(path string) (*Settings, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var s Settings
	err = json.Unmarshal(data, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

// AbsolutePaths makes all paths in settings absolute using the given
// absolute root
func (s *settings.Settings) AbsolutePaths(root string) {
	for i := range s.IgnoreDirs {
		s.IgnoreDirs[i] := helpers.AbsolutePath(s.IgnoreDirs[i], root)
	}
	for i := range s.LilypondIncludes {
		s.LilypondIncludes[i] := helpers.AbsolutePath(s.LilypondIncludes[i], root)
	}
	s.OutputDir := helpers.AbsolutePath(s.OutputDir, root)
}
