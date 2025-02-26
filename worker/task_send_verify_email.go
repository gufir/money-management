package worker

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
	db "github.com/gufir/money-management/db/sqlc"
	"github.com/gufir/money-management/utils"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const TaskSendVerifyEmail = "send_verify_email"

type PayloadSendVerifyEmail struct {
	UserID string `json:"user_id"`
}

func (distributor *RedisTaskDistributor) DistributeTaskVerifyEmail(
	ctx context.Context,
	payload *PayloadSendVerifyEmail,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %v", err)
	}

	task := asynq.NewTask(TaskSendVerifyEmail, jsonPayload, opts...)
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %v", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("task enqueued")
	return nil
}

func (processor *RedisTaskProcessor) ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error {
	var payload PayloadSendVerifyEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	userID, err := uuid.Parse(payload.UserID)
	if err != nil {
		return fmt.Errorf("failed to parse user id: %w", asynq.SkipRetry)
	}

	user, err := processor.store.GetUserByUserId(ctx, userID)
	if err != nil {

		if errors.Is(err, db.ErrRecordNotFound) {
			return fmt.Errorf("fuser doesn't exist: %w", asynq.SkipRetry)
		}

		return fmt.Errorf("failed to get user: %w", err)
	}

	verifyemail, err := processor.store.CreateVerifyEmails(ctx, db.CreateVerifyEmailsParams{
		UserID:     user.UserUuid,
		Email:      user.Email,
		SecretCode: utils.RandomString(32),
	})

	if err != nil {
		return fmt.Errorf("failed to create verify email: %w", err)
	}

	subject := "Welcome to Money Wise! Verify your email"
	frontendUrl := "http://localhost:3000/verify-email"
	verifyUrl := fmt.Sprintf("%s?email_id=%d&secret_code=%s", frontendUrl, verifyemail.ID, verifyemail.SecretCode)
	content := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="UTF-8">
		<title>Verify Your Email</title>
	</head>
	<body style="font-family: Arial, sans-serif; background-color: #f4f4f4; margin: 0; padding: 20px;">
		<table width="100%%" border="0" cellspacing="0" cellpadding="0" style="max-width: 600px; background: white; padding: 20px; margin: auto; border-radius: 8px; box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);">
			<tr>
				<td style="text-align: center;">
					<h2 style="color: #333;">Welcome to <span style="color: #007bff;">Money Wise</span>!</h2>
					<p style="color: #666; font-size: 16px;">Hello %s,</p>
					<p style="color: #666; font-size: 16px;">Thank you for registering with us!</p>
					<p style="color: #666; font-size: 16px;">Please click the button below to verify your email address:</p>
					<a href="%s" style="display: inline-block; padding: 12px 20px; background: #007bff; color: white; text-decoration: none; border-radius: 5px; font-size: 16px; margin-top: 10px;">
						Verify Email
					</a>
					<p style="color: #666; font-size: 14px; margin-top: 20px;">If the button doesn't work, you can also use the following link:</p>
					<p style="word-wrap: break-word;"><a href="%s" style="color: #007bff;">%s</a></p>
					<p style="color: #888; font-size: 12px; margin-top: 20px;">If you didn't create an account, please ignore this email.</p>
				</td>
			</tr>
		</table>
	</body>
	</html>
	`, user.FullName, verifyUrl, verifyUrl, verifyUrl)

	to := []string{user.Email}

	err = processor.mailer.SendEmail(subject, content, to, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("email", user.Email).Msg("processed task")

	return nil

}
