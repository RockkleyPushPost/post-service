package dto

import (
	"errors"
	"github.com/google/uuid"
)

type CreatePostDto struct {
	UserUUID uuid.UUID `json:"userUUID"`
	Type     string    `json:"type"`
	Content  string    `json:"content"`
}

func (dto *CreatePostDto) Validate() error {
	if dto.UserUUID == uuid.Nil {
		return errors.New("userUUID is required")
	}
	if dto.Type == "" {
		return errors.New("type is required")
	}
	if dto.Content == "" {
		return errors.New("content is required")
	}
	return nil
}
