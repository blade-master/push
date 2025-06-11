package service

import (
	"context"
	pushcore "github.com/blade-master/push/rpc_gen/kitex_gen/pushcore"
)

type SendGameEventService struct {
	ctx context.Context
} // NewSendGameEventService new SendGameEventService
func NewSendGameEventService(ctx context.Context) *SendGameEventService {
	return &SendGameEventService{ctx: ctx}
}

// Run create note info
func (s *SendGameEventService) Run(req *pushcore.GameEvent) (resp *pushcore.PushAck, err error) {
	// Finish your business logic.

	return
}
