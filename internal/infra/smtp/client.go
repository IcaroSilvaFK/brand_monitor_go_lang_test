package smtp

import (
	"bufio"
	"log"
	"os"
	"path/filepath"

	"github.com/go-mail/mail"
)

type SMTClient struct {
	dialer  *mail.Dialer
	from    string
	address string
}

type SMTPClientInterface interface {
	CreateAccountSendEmail(to string) error
	SendNewProductAdded(to string) error
}

var (
	SMTP_USER = "SMTP_USER"
	SMTP_PASS = "SMTP_PASS"
	SMTP_HOST = "SMTP_HOST"
	SMTP_PORT = "SMTP_PORT"
	SMTP_FROM = "SMTP_FROM"
)

// Subiscrible
func NewSMTPClient() SMTPClientInterface {

	// smtpUser := os.Getenv(SMTP_USER)
	// smtpPass := os.Getenv(SMTP_PASS)
	// smtpHost := os.Getenv(SMTP_HOST)
	// smtpFrom := os.Getenv(SMTP_FROM)

	dialer := mail.NewDialer("sandbox.smtp.mailtrap.io", 2525, "2bf0fb4c61ecd4", "5ad2006a67d6bd")

	return &SMTClient{
		dialer:  dialer,
		from:    "teste@test.com",
		address: "smtp.mailtrap.io",
	}
}

func (client *SMTClient) CreateAccountSendEmail(to string) error {
	abs, err := filepath.Abs("internal/infra/smtp/templates/subscrible.html")

	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Open(abs)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	body := ""

	for scanner.Scan() {
		body += scanner.Text() + "\n"
	}

	m := mail.NewMessage()

	m.SetHeader("From", client.from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Thanks from brand monitor")
	m.SetBody("text/html", body)

	if err := client.dialer.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func (client *SMTClient) SendNewProductAdded(to string) error {
	abs, err := filepath.Abs("internal/infra/smtp/templates/new_product_added.html")

	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(abs)

	if err != nil {
		log.Fatal(err)
	}
	body := ""
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		body += scanner.Text() + "\n"
	}

	m := mail.NewMessage()

	m.SetHeader("From", client.from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "New product added")
	m.SetBody("text/html", body)

	if err := client.dialer.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
