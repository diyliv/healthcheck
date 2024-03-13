package opcda

import (
	"context"

	"go.uber.org/zap"

	"github.com/diyliv/healthcheck/internal/models"
)

type healthcheck struct {
	logger *zap.Logger
}

func NewHealthCheck(opcda models.OPCDAHealthCheck) *healthcheck {

	return &healthcheck{}
}

func (h *healthcheck) Init(ctx context.Context) error {
	return nil
}
