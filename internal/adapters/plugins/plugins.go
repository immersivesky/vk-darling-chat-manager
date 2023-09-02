package plugins

import (
	// "gitlab.com/immersivesky/affinity-plugins/go"
	"google.golang.org/grpc"
)

type PluginsClient struct{}

func NewPluginsClient() *PluginsClient {
	conn, err := grpc.Dial("localhost:4200")
	if err != nil {
		panic(err)
	}

	client := pb.NewPluginsClient(conn)

	return &PluginsClient{}
}
