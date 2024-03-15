package opcda

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/diyliv/opc"

	"github.com/diyliv/healthcheck/internal/models"
	opcdacheck "github.com/diyliv/healthcheck/pkg/errors"
)

type healthcheck struct {
	logger  *zap.Logger
	conn    opc.Connection
	mu      sync.RWMutex
	storage map[string]opc.Connection
}

func NewHealthCheck(ctx context.Context, logger *zap.Logger, opcda models.OPCDAHealthCheck) (*healthcheck, error) {
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
		logger:  logger,
		conn:    conn,
		storage: make(map[string]opc.Connection),
	}

	healthcheck.mu.Lock()
	healthcheck.storage[opcda.Server] = conn
	healthcheck.mu.Unlock()

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
	h.mu.Lock()
	defer h.mu.Unlock()
	_, ok := h.storage[opcda.Server]
	if ok {
		return opcdacheck.ErrOpcServerExists
	}
	conn, err := connectToOpc(ctx, opcda)
	if err != nil {
		return err
	}
	h.storage[opcda.Server] = conn
	go h.check(ctx, opcda.HeartBeat)
	return nil
}

func (h *healthcheck) Stop(cancel context.CancelFunc, server string) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	conn, ok := h.storage[server]
	if !ok {
		return opcdacheck.ErrNoOpcConnection
	}
	conn.Close()
	cancel()
	delete(h.storage, server)
	return nil
}

func connectToOpc(ctx context.Context, opcda models.OPCDAHealthCheck) (opc.Connection, error) {
	conn, err := opc.NewConnection(opcda.Server, opcda.Nodes, opcda.Tags)
	if err != nil {
		return nil, err
	}
	return conn, err
}
