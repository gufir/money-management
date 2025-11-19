package worker

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	db "github.com/gufir/money-management/db/sqlc"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

func (processor *RedisTaskProcessor) ProcessTaskSendForgotPassword(
	ctx context.Context,
	task *asynq.Task,
) error {

	var payload PayloadSendForgotPassword
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	user, err := processor.store.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return fmt.Errorf("user not found: %w", asynq.SkipRetry)
		}
		return fmt.Errorf("failed to get user: %w", err)
	}

	resetUrl := fmt.Sprintf(
		"http://localhost:3000/reset-password?id=%s&user_id=%s&token=%s",
		payload.ResetID,
		user.UserUuid,
		payload.Token,
	)

	subject := "Reset Your Password - Money Wise"

	content := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="UTF-8">
		<title>Reset Password</title>
	</head>
	<body style="font-family: Arial; background-color:#f4f4f4; padding:20px;">
		<table width="100%%" style="max-width:600px; margin:auto; background:white; padding:20px;">
			<tr>
				<td>
					<h2 style="text-align:center;">Reset Your Password</h2>
					<p>Hello %s,</p>
					<p>You requested to reset your password.</p>
					<p>Click the button below to set a new password:</p>

					<a href="%s"
						style="background:#007bff; color:white; padding:12px 20px; text-decoration:none; border-radius:5px;">
						Reset Password
					</a>

					<p style="margin-top:20px;">If the button doesn’t work, click this link:</p>
					<a href="%s">%s</a>

					<p style="color:#888; margin-top:20px;">If you didn't request this, please ignore this email.</p>
				</td>
			</tr>
		</table>
	</body>
	</html>
	`, user.FullName, resetUrl, resetUrl, resetUrl)

	err = processor.mailer.SendEmail(subject, content, []string{user.Email}, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to send reset email: %w", err)
	}

	log.Info().
		Str("type", task.Type()).
		Bytes("payload", task.Payload()).
		Str("email", user.Email).
		Msg("processed reset password task")

	return nil
}
