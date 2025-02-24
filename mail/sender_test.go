package mail

import (
	"testing"

	"github.com/gufir/money-management/utils"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	config, err := utils.LoadConfig("..")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	subject := "Test Email"
	content := `
	<h1>Hello World</h1>
	<p>This is a test email from Money Wise</p>
	`
	to := []string{config.EmailSenderAddress}
	attachfile := []string{"../README.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachfile)
	require.NoError(t, err)
}
