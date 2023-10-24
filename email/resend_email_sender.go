package email

import (
	"fmt"
	"os"

	"github.com/resendlabs/resend-go"
)

func ResendEmailSender(toAddres, serverName, serverAddress, timestamp string) error {
	apiKey := os.Getenv("RESEND_API_KEY")

	client := resend.NewClient(apiKey)

	html := generateServerDownEmail(serverName, serverAddress, timestamp)

	params := &resend.SendEmailRequest{
		From:    "vigilate@resend.dev",
		To:      []string{toAddres},
		Html:    html,
		Subject: "Server Down Alert",
	}

	sent, err := client.Emails.Send(params)
	if err != nil {
		return err
	}
	fmt.Println(sent.Id)

	return nil
}

func generateServerDownEmail(serverName, serverAddress, timestamp string) string {
	htmlTemplate := `
			<!DOCTYPE html>
			<html>
			<head>
					<meta charset="UTF-8">
					<title>Server Down Alert</title>
			</head>
			<body>
					<h1>Server Down Alert</h1>
					<p>Dear User,</p>
					<p>This is to inform you that the server named <strong>%s</strong> at the address <strong>%s</strong> is currently down.</p>
					<p>We apologize for any inconvenience this may cause and are actively working to resolve the issue.</p>
					<p>Timestamp: <strong>%s</strong></p>
					<p>If you have any questions or require further assistance, please do not hesitate to contact our support team.</p>
					<p>Thank you for your understanding.</p>
					<p>Sincerely,<br>Your Support Team</p>
			</body>
			</html>
	`

	return fmt.Sprintf(htmlTemplate, serverName, serverAddress, timestamp)
}
