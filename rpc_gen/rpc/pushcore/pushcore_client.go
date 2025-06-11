package pushcore

import (
	"context"
	pushcore "github.com/blade-master/push/rpc_gen/kitex_gen/pushcore"

	"github.com/blade-master/push/rpc_gen/kitex_gen/pushcore/pushservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() pushservice.Client
	Service() string
	SendGameEvent(ctx context.Context, Req *pushcore.GameEvent, callOptions ...callopt.Option) (r *pushcore.PushAck, err error)
	BroadcastGlobalEvent(ctx context.Context, Req *pushcore.GlobalEvent, callOptions ...callopt.Option) (r *pushcore.PushAck, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := pushservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient pushservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() pushservice.Client {
	return c.kitexClient
}

func (c *clientImpl) SendGameEvent(ctx context.Context, Req *pushcore.GameEvent, callOptions ...callopt.Option) (r *pushcore.PushAck, err error) {
	return c.kitexClient.SendGameEvent(ctx, Req, callOptions...)
}

func (c *clientImpl) BroadcastGlobalEvent(ctx context.Context, Req *pushcore.GlobalEvent, callOptions ...callopt.Option) (r *pushcore.PushAck, err error) {
	return c.kitexClient.BroadcastGlobalEvent(ctx, Req, callOptions...)
}
