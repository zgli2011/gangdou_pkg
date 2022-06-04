package gangdoupkg

import (
	"context"
	"net/smtp"
	"strings"
)

type Mail struct {
	mailAuth smtp.Auth
	toList   []string
	subject  string
	msg      string
	host     string
	username string
}

func (m *Mail) Auth(username string, password string, host string) *Mail {
	m.host = host
	hosts := strings.Split(host, ":")
	m.mailAuth = smtp.PlainAuth("", username, password, hosts[0])
	return m
}

func (m *Mail) To(mailList []string) *Mail {
	m.toList = mailList
	return m
}

func (m *Mail) Subject(subject string) *Mail {
	m.subject = subject
	return m
}

func (m *Mail) Html(msg string) *Mail {
	m.msg = msg
	return m
}
func (m *Mail) Send(ctx context.Context) error {
	smtp.SendMail(m.host, m.mailAuth, m.username, m.toList, []byte(m.msg))
}
