package utils

import (
    "fmt"
    "net/smtp"
)

// Mailer represents an email sender.
type Mailer struct {
    senderEmail     string
    senderPassword  string
    smtpServer      string
    smtpPort        string
}

// NewMailer creates a new instance of Mailer.
func NewMailer(senderEmail, senderPassword, smtpServer, smtpPort string) *Mailer {
    return &Mailer{
        senderEmail:    senderEmail,
        senderPassword: senderPassword,
        smtpServer:     smtpServer,
        smtpPort:       smtpPort,
    }
}

// SendEmail sends an email using the configured SMTP server.
func (m *Mailer) SendEmail(to, subject, body string) error {
    auth := smtp.PlainAuth("", m.senderEmail, m.senderPassword, m.smtpServer)
    msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body))
    addr := fmt.Sprintf("%s:%s", m.smtpServer, m.smtpPort)

    err := smtp.SendMail(addr, auth, m.senderEmail, []string{to}, msg)
    if err != nil {
        return err
    }

    return nil
}
