// Copyright 2019 LinkinStar
// license that can be found in the LICENSE file.

package gu

import (
	"bytes"
	"math/rand"
	"strconv"
	"time"
)

// generate random number string
func GenerateRandomNumber(maxLength int) string {
	rand.Seed(time.Now().UnixNano())

	var buffer bytes.Buffer
	for i := 0; i < maxLength; i++ {
		str := strconv.Itoa(rand.Intn(10))
		buffer.WriteString(str)
	}

	return buffer.String()
}
