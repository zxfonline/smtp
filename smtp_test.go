// Copyright 2016 zxfonline@sina.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package smtp

import (
	"io"
	"testing"
	"time"
)

var _ io.Writer = &Smtp{}

func testSmtp(t *testing.T) {
	smtp := NewSmtp("zxfonline@sina.com", "password", "subject", "smtp.sina.com:25", []string{"359168950@qq.com"})

	size, err := smtp.Write([]byte("content1"))
	if err != nil {
		t.Error(err)
	}
	if size <= 0 {
		t.Error("head content error")
	}

	time.Sleep(10 * time.Second)

	size, err = smtp.Write([]byte("content2"))
	if err != nil {
		t.Error(err)
	}
	if size <= 0 {
		t.Error("head content error")
	}
}
