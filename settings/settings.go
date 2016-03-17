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

// Settings provides a structure to interact with the settings
// of a Ponder library
type Settings struct {
	Name             string   // Name of the Ponder library
	IgnoreDirs       []string // Directories to be ignored on search
	LilypondIncludes []string // Directories to be included when running the lilypond compiler
	OutputDir        string   // Directory in which all complete file are stored
}
