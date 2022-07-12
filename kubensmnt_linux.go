//go:build linux && amd64 && !_test_nonlinux
// +build linux,amd64,!_test_nonlinux

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
	"os"
	"syscall"
)

const (
	CLONE_NEWNS = 0x00020000
	SYS_SETNS   = 308
)

func enter(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	return setNs(file.Fd(), CLONE_NEWNS)
}

func setNs(fd uintptr, nstype uintptr) syscall.Errno {
	_, _, err := syscall.Syscall(SYS_SETNS, uintptr(fd), uintptr(nstype), 0)
	return err
}

func setAllNs(fd uintptr, nstype uintptr) syscall.Errno {
	_, _, err := syscall.AllThreadsSyscall(SYS_SETNS, uintptr(fd), uintptr(nstype), 0)
	return err
}
