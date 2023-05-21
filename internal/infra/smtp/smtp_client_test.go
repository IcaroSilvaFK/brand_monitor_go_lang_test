package smtp_test

import (
	"testing"

	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/smtp"
	"github.com/stretchr/testify/assert"
)

func TestShouldSendEmail(t *testing.T) {
	client := smtp.NewSMTPClient()

	err := client.CreateAccountSendEmail("icaro@icaro.com")

	assert.Nil(t, err)
}
