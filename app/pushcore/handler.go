package main

import (
	"context"
	"github.com/blade-master/push/app/pushcore/biz/service"
)

// PushServiceImpl implements the last service interface defined in the IDL.
type PushServiceImpl struct{}

// SendGameEvent implements the PushServiceImpl interface.
func (s *PushServiceImpl) SendGameEvent(ctx context.Context, req *pushcore.GameEvent) (resp *pushcore.PushAck, err error) {
	resp, err = service.NewSendGameEventService(ctx).Run(req)

	return resp, err
}

// BroadcastGlobalEvent implements the PushServiceImpl interface.
func (s *PushServiceImpl) BroadcastGlobalEvent(ctx context.Context, req *pushcore.GlobalEvent) (resp *pushcore.PushAck, err error) {
	resp, err = service.NewBroadcastGlobalEventService(ctx).Run(req)

	return resp, err
}
