package healtcheck

import "context"

type healtcheck struct{}

func NewHealthCheck() *healtcheck {
	return &healtcheck{}
}

func (h *healtcheck) Init(ctx context.Context) error {
	return nil
}
