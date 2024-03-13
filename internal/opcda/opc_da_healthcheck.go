package opcda

import (
	"go.uber.org/zap"

	"github.com/diyliv/opc"

	"github.com/diyliv/healthcheck/internal/models"
	opcdacheck "github.com/diyliv/healthcheck/pkg/errors"
)

type healthcheck struct {
	logger *zap.Logger
	conn   opc.Connection
}

func NewHealthCheck(logger *zap.Logger, opcda models.OPCDAHealthCheck) (*healthcheck, error) {
	if opcda.Server == "" {
		return nil, opcdacheck.ErrNoServerSpecified
	} else if len(opcda.Nodes) == 0 {
		return nil, opcdacheck.ErrNoNodeSpecified
	}
	// no checking for tags field because in some cases
	// they are needed to be specified after
	// initialization part

	conn, err := opc.NewConnection(opcda.Server, opcda.Nodes, opcda.Tags)
	if err != nil {
		return nil, err
	}

	return &healthcheck{
		logger: logger,
		conn:   conn,
	}, nil
}
