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

package main

import "github.com/jjdekker/ponder/cmd"

func main() {
	cmd.Execute()
}

// TODO: Make a command to check that settings are legal
// TODO: Add a clean command
// TODO: Add support for Ly files to the add command
// TODO: Allow an latex input file for styling
// TODO: Sort output into categories
