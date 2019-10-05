// Copyright 2019 LinkinStar
// license that can be found in the LICENSE file.

package gu

import (
	"os"
)

// create dir recursion
func CreateDirIfNotExist(dir string) error {
	if CheckPathIfNotExist(dir) {
		return nil
	}
	return os.MkdirAll(dir, os.ModePerm)
}

// if file exist return true
func CheckPathIfNotExist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
