// Copyright 2016 zxfonline@sina.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package smtp

import (
	"bytes"
	"net/smtp"
	"strings"
)

type Smtp struct {
	username string   // smtp账号
	password string   // smtp密码
	host     string   // smtp主机，需要带上端口
	sendTo   []string // 接收者
	subject  string   // 邮件主题

	// 邮件内容的缓存
	cache *bytes.Buffer
	// 邮件头部分的长度
	headerLen int

	auth smtp.Auth
}

func NewSmtp(username, password, subject, host string, sendTo []string) *Smtp {
	ret := &Smtp{
		username: username,
		password: password,
		subject:  subject,
		host:     host,
		sendTo:   sendTo,
	}
	ret.Init()
	return ret
}

func (s *Smtp) Init() {
	s.cache = bytes.NewBufferString("")
	s.cache.Grow(1024)

	// to
	s.cache.WriteString("To: ")
	s.cache.WriteString(strings.Join(s.sendTo, ";"))
	s.cache.WriteString("\r\n")

	// from
	s.cache.WriteString("From: ")
	s.cache.WriteString(s.username)
	s.cache.WriteString("\r\n")

	// subject
	s.cache.WriteString("Subject: ")
	s.cache.WriteString(s.subject)
	s.cache.WriteString("\r\n")

	// mime-version
	s.cache.WriteString("MIME-Version: ")
	s.cache.WriteString("1.0\r\n")

	// contentType
	s.cache.WriteString(`Content-Type: text/plain; charset="utf-8"`)
	s.cache.WriteString("\r\n\r\n")

	s.headerLen = s.cache.Len()

	// 去掉端口部分
	h := strings.Split(s.host, ":")[0]
	s.auth = smtp.PlainAuth("", s.username, s.password, h)
}

func (s *Smtp) Write(msg []byte) (int, error) {
	s.cache.Write(msg)

	err := smtp.SendMail(
		s.host,
		s.auth,
		s.username,
		s.sendTo,
		s.cache.Bytes(),
	)
	l := s.cache.Len()

	s.cache.Truncate(s.headerLen)

	return l, err
}
