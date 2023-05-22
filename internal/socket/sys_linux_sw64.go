// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build sw64
// +build sw64

package socket

const (
	sysRECVMMSG = 0x1df
	sysSENDMMSG = 0x1f7
)
