package main

import (
	"context"
	"log/slog"
	"net"
	"os"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/rpc"
	"github.com/tablelandnetwork/basin-cli/pkg/basinprovider"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:"+os.Getenv("PORT"))
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	slog.Info("Listening", "port", os.Getenv("PORT"))

	server := basinprovider.NewBasinServerMock()
	for {
		conn, err := listener.Accept()
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
		defer conn.Close()

		client := basinprovider.Publications_ServerToClient(server)
		rpcConn := rpc.NewConn(rpc.NewStreamTransport(conn), &rpc.Options{
			BootstrapClient: capnp.Client(client),
		})
		defer rpcConn.Close()

		ctx := context.Background()

		// Block until the connection terminates.
		select {
		case <-rpcConn.Done():
			slog.Info("connection closed")
		case <-ctx.Done():
			conn.Close()
		}

	}
}
