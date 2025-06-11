package service

import (
	"context"
	pushcore "github.com/blade-master/push/rpc_gen/kitex_gen/pushcore"
	"testing"
)

func TestSendGameEvent_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSendGameEventService(ctx)
	// init req and assert value

	req := &pushcore.GameEvent{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
