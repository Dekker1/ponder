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

package compiler

import (
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/jjdekker/ponder/helpers"
	"github.com/jjdekker/ponder/settings"
)

// compilePath calls the given function on all files that aren't in any
// of the ignored directories
func compilePath(root string, opts *settings.Settings,
	f func(string, os.FileInfo) error) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		// Handle walking error
		if err != nil {
			log.withFields(log.Fields{
				"error": err,
				"path":  path,
			}).Warning("error occurred transversing project path")
			return nil
		}

		if info.IsDir() {
			// Skip directories that are ignored
			relPath, err = filepath.Rel(root, path)
			helpers.Check(err, "Unable to create relative Path")
			for dir := range append(append(opts.IgnoreDirs, opts.LilypondIncludes), []string{opts.OutputDir}) {
				if relPath == dir || (filepath.IsAbs(dir) && path == dir) {
					log.WithFields(log.Fields{"path": path}).Info("Ignoring directory")
					return filepath.SkipDir
				}
			}
		} else {
			// Call function on non-directory
			err = f(path, info)
			if err != nil {
				log.WithFields(log.Fields{
					"error": err,
					"path":  path,
				}).Error("error occured processing files")
			}
		}
		return nil
	}
}
