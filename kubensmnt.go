/*
 * Copyright 2022 Jim Ramsay <jramsay@redhat.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package kubensmnt

import (
	"errors"
	"os"
)

// EnvName is the name of the environment variable where we check for a mount namespace bindmount.
// If unset, no action is taken.  If set and points at a mount namespace
// bindmount, enter that mount namespace before executing this Go program.  If
// set and an error occurs, the Status function will report an error.
const EnvName = "KUBENSMNT"

// DefaultKubensmnt is the default location where kubensmnt.service places its bound mount namespace
const DefaultKubensmnt = "/run/kubensmnt/mnt"

// Enter enters the namespace given by the path
func Enter(path string) error {
	return enter(path)
}

// Autodetect tries to enter the kubensmnt.service mount namespace using well-known defaults.
// First $KUBENSMNT in the environment, then the DefaultKubensmnt defined above.
func Autodetect() (string, error) {
	path := os.Getenv(EnvName)
	if path == "" {
		path = DefaultKubensmnt
		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			return "", nil
		}
	}
	return path, Enter(path)
}
