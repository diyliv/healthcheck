package opcda

import (
	"errors"

	"go.uber.org/zap"

	"github.com/diyliv/opc"

	"github.com/diyliv/healthcheck/internal/models"
	opcdacheck "github.com/diyliv/healthcheck/pkg/errors"
)

type healthcheck struct {
	logger *zap.Logger
	conn   opc.Connection
}

func NewHealthCheck(opcda models.OPCDAHealthCheck) (*healthcheck, error) {
	if opcda.Server == "" {
		return nil, opcdacheck.ErrNoServerSpecified
	} else if len(opcda.Nodes) == 0 {
		return nil, errors.New("node must be specified")
	}

	return &healthcheck{}, nil
}
