package opcda

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/diyliv/opc"

	"github.com/diyliv/healthcheck/internal/models"
	opcdacheck "github.com/diyliv/healthcheck/pkg/errors"
)

type healthcheck struct {
	logger *zap.Logger
	conn   opc.Connection
	cache  models.Cache
}

func NewHealthCheck(ctx context.Context, logger *zap.Logger, opcda models.OPCDAHealthCheck, cache models.Cache) (*healthcheck, error) {
	// no checking for tags field because in some cases
	// they are needed to be specified after
	// initialization part
	if opcda.Server == "" {
		return nil, opcdacheck.ErrNoServerSpecified
	} else if len(opcda.Nodes) == 0 {
		return nil, opcdacheck.ErrNoNodeSpecified
	}

	conn, err := connectToOpc(ctx, opcda)
	if err != nil {
		return nil, err
	}
	healthcheck := &healthcheck{
		logger: logger,
		conn:   conn,
		cache:  cache,
	}
	if err := healthcheck.cache.Put(ctx, opcda.Server, conn); err != nil {
		return nil, err
	}

	go healthcheck.check(ctx, opcda.HeartBeat)

	return healthcheck, nil
}

func (h *healthcheck) check(ctx context.Context, heartBeat time.Duration) error {
	ticker := time.NewTicker(heartBeat)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			_, err := h.conn.Read()
			if err != nil {
				h.logger.Error("Error while checking OPC-DA server: " + err.Error())
			}
		}
	}
}

func (h *healthcheck) Add(ctx context.Context, opcda models.OPCDAHealthCheck) error {
	conn, err := connectToOpc(ctx, opcda)
	if err != nil {
		return err
	}
	if err := h.cache.Put(ctx, opcda.Server, conn); err != nil {
		return err
	}
	go h.check(ctx, opcda.HeartBeat)
	return nil
}

func (h *healthcheck) Stop(ctx context.Context, cancel context.CancelFunc, server string) error {
	val, err := h.cache.Get(ctx, server)
	if err != nil {
		return err
	}
	conn, ok := val.(opc.Connection)
	if !ok {
		return opcdacheck.ErrOPCDAConvertion
	}
	conn.Close()
	if err := h.cache.Remove(ctx, server); err != nil {
		return err
	}
	cancel()
	return nil
}

func connectToOpc(ctx context.Context, opcda models.OPCDAHealthCheck) (opc.Connection, error) {
	conn, err := opc.NewConnection(opcda.Server, opcda.Nodes, opcda.Tags)
	if err != nil {
		return nil, err
	}
	return conn, err
}
