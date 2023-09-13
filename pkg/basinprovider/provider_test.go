package basinprovider

import (
	"bytes"
	"context"
	"testing"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/rpc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/stretchr/testify/require"
	"github.com/tablelandnetwork/basin-cli/internal/app"
	basincapnp "github.com/tablelandnetwork/basin-cli/pkg/capnp"
	"github.com/tablelandnetwork/basin-cli/pkg/pgrepl"
	"google.golang.org/grpc/test/bufconn"
)

func TestBasinProvider_Create(t *testing.T) {
	// in this test we create a fake tx,
	// send to the server, the server deserialize it and send the value back

	bp := newServer()
	err := bp.Create(context.Background(), "n", "t", basincapnp.Schema{}, common.HexToAddress(""))
	require.NoError(t, err)
}

func TestBasinProvider_Push(t *testing.T) {
	// in this test we create a fake tx,
	// send to the server, the server deserialize it and send the value back

	bp := newServer()
	tx := newTx(t, &pgrepl.Tx{
		CommitLSN: 333,
		Records: []pgrepl.Record{
			{
				Action: "I",
			},
		},
	})

	err := bp.Push(context.Background(), "n", "t", tx, []byte{})
	require.NoError(t, err)
}

func TestBasinProvider_Upload(t *testing.T) {
	bp := newServer()

	data := bytes.NewReader([]byte{})
	privateKey, _ := crypto.GenerateKey()
	err := bp.Upload(context.Background(), "", "", data, app.NewSigner(privateKey), bytes.NewBuffer([]byte{}))
	require.NoError(t, err)
}

func newTx(t *testing.T, tx *pgrepl.Tx) basincapnp.Tx {
	capnpTx, err := basincapnp.FromPgReplTx(tx)
	require.NoError(t, err)

	return capnpTx
}

func newServer() *BasinProvider {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	mock := NewBasinServerMock()
	p := Publications_ServerToClient(mock)
	bootstrapClient := capnp.Client(p)

	go func() {
		_ = rpc.Serve(lis, bootstrapClient)
	}()

	return &BasinProvider{
		p:        p,
		provider: "mock",
		ctx:      context.Background(),
		cancel: func() {
			close(make(chan struct{}))
		},
	}
}
