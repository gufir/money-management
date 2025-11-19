package worker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const TaskSendForgotPassword = "send_forgot_password"

type PayloadSendForgotPassword struct {
	Email   string    `json:"email"`
	Token   string    `json:"token"`
	ResetID uuid.UUID `json:"id"`
}

func (distributor *RedisTaskDistributor) DistributeTaskResetPassword(
	ctx context.Context,
	payload *PayloadSendForgotPassword,
	opts ...asynq.Option,
) error {

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %v", err)
	}

	task := asynq.NewTask(TaskSendForgotPassword, jsonPayload, opts...)

	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to enqueue reset password task: %v", err)
	}

	log.Info().
		Str("type", task.Type()).
		Bytes("payload", task.Payload()).
		Str("queue", info.Queue).
		Int("max_retry", info.MaxRetry).
		Msg("reset password task enqueued")

	return nil
}
