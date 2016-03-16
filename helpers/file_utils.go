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

package helpers

import (
	"os"
	"path/filepath"
)

// CleanPath returns a cleaned path from a given string relative
// to the working directory unless explicitly absolute
func CleanPath(path string) (string, error) {
	if path == "" {
		// No given path, use working directory
		return os.Getwd()
	} else if filepath.IsAbs(path) {
		// Given path is absolute
		return filepath.Clean(path), nil
	} else {
		// Given path is a relative path
		dir, err := os.Getwd()
		if err == nil {
			return filepath.Join(dir, path), nil
		}
		return "", err
	}
}

// Exists checks if a file or directory exists.
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}