package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrExpiredToken = errors.New("token is expired")
	ErrInvalidToken = errors.New("token is invalid")
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	UserID    uuid.UUID `json:"user_id"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(userID uuid.UUID, username string, role string, duration time.Duration) (*Payload, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        id,
		Username:  username,
		UserID:    userID,
		Role:      role,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}

func (p *Payload) GetIssuer() (string, error) {
	return "", nil
}

func (p *Payload) GetAudience() (jwt.ClaimStrings, error) {
	return nil, nil
}

func (p *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	expirationTime := jwt.NewNumericDate(p.ExpiredAt)
	return expirationTime, nil
}

func (p *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	return nil, nil
}

func (p *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	issueTime := jwt.NewNumericDate(p.CreatedAt)
	return issueTime, nil
}

func (p *Payload) GetSubject() (string, error) {
	return p.Username, nil
}
