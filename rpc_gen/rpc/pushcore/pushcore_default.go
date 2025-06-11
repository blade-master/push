package pushcore

import (
	"context"
	pushcore "github.com/blade-master/push/rpc_gen/kitex_gen/pushcore"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func SendGameEvent(ctx context.Context, req *pushcore.GameEvent, callOptions ...callopt.Option) (resp *pushcore.PushAck, err error) {
	resp, err = defaultClient.SendGameEvent(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SendGameEvent call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func BroadcastGlobalEvent(ctx context.Context, req *pushcore.GlobalEvent, callOptions ...callopt.Option) (resp *pushcore.PushAck, err error) {
	resp, err = defaultClient.BroadcastGlobalEvent(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "BroadcastGlobalEvent call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
