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

import log "github.com/Sirupsen/logrus"

// Crash outputs it's arguments to the log and stops the program
func Crash(err error, msg string) {
	log.WithFields(log.Fields{
		"error": err,
	}).Fatal(msg)
}

// Check calls Crash if the error is not nil
func Check(err error, msg string) {
	if err != nil {
		Crash(err, msg)
	}
}