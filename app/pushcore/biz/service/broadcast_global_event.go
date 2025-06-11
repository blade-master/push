package service

import (
	"context"
	pushcore "github.com/blade-master/push/rpc_gen/kitex_gen/pushcore"
)

type BroadcastGlobalEventService struct {
	ctx context.Context
} // NewBroadcastGlobalEventService new BroadcastGlobalEventService
func NewBroadcastGlobalEventService(ctx context.Context) *BroadcastGlobalEventService {
	return &BroadcastGlobalEventService{ctx: ctx}
}

// Run create note info
func (s *BroadcastGlobalEventService) Run(req *pushcore.GlobalEvent) (resp *pushcore.PushAck, err error) {
	// Finish your business logic.

	return
}
