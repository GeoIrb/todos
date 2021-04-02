package smtp

import (
	"context"
	"time"

	mail "github.com/xhit/go-simple-mail"
)

// SMTP for send from smt protocol
type SMTP struct {
	srcAddress string
	server     *mail.SMTPServer
}

// New ...
func New(
	address string,
	password string,
	smtpHost string,
	smtpPort int,

	connectTimeout time.Duration,
	sendTimeout time.Duration,
) *SMTP {
	return &SMTP{
		srcAddress: address,
		server: &mail.SMTPServer{
			Username:       address,
			Password:       password,
			Host:           smtpHost,
			Port:           smtpPort,
			ConnectTimeout: connectTimeout,
			SendTimeout:    sendTimeout,
			Encryption:     mail.EncryptionSSL,
		},
	}
}

// Send message to dst.
func (s *SMTP) Send(ctx context.Context, dst, message string) error {
	email := mail.NewMSG()
	email.SetFrom(s.srcAddress).
		AddTo(dst).
		SetBody(mail.TextPlain, message)

	client, err := s.server.Connect()
	if err != nil {
		return err
	}
	return email.Send(client)
}
