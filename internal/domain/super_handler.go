package domain

import (
	"context"

	"github.com/ai-core-tools/hsu-echo/go/domain"
	"github.com/ai-core-tools/hsu-echo/go/logging"
)

func NewSuperHandler(logger logging.Logger) domain.Contract {
	return &superHandler{
		logger: logger,
	}
}

type superHandler struct {
	logger logging.Logger
}

func (h *superHandler) Echo(ctx context.Context, message string) (string, error) {
	return "go-super-echo: " + message, nil
}
