/*
 *  \
 *  \\,
 *   \\\,^,.,,.                     Zero to Hero
 *   ,;7~((\))`;;,,               <zerotohero.dev>
 *   ,(@') ;)`))\;;',    stay up to date, be curious: learn
 *    )  . ),((  ))\;,
 *   /;`,,/7),)) )) )\,,
 *  (& )`   (,((,((;( ))\,
 */

package main

import (
	"context"
	"flag"
	"github.com/spiffe/go-spiffe/v2/spiffegrpc/grpccredentials"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/peer"
	"log"
	"os"
	"time"
)

func sendRequest(ctx context.Context, client helloworld.GreeterClient) {
	peer := new(peer.Peer)
	res, err := client.SayHello(ctx, &helloworld.HelloRequest{
		Name: "Volkan Özçelik",
	}, grpc.Peer(peer))
	if err != nil {
		log.Printf("Failed to say hello: %v\n", err)
		return
	}

	// #region
	// Learn server’s identity from the peer SVID.
	serverId := "UnknownServer"
	if peerId, ok := grpccredentials.PeerIDFromPeer(peer); ok {
		serverId = peerId.String()
	}
	// #endregion

	log.Printf("%s said %q\n", serverId, res.Message)
}

func main() {
	var addr string
	flag.StringVar(&addr, "addr", "", "host:port of the server")
	flag.Parse()

	if addr == "" {
		addr = os.Getenv("GREETER_SERVER_ADDR")
	}
	if addr == "" {
		addr = "localhost:8123"
	}

	log.Println("Client starting up… (%s)", addr)

	ctx := context.Background()

	// create an id struct by parsing the spiffe id string.
	serverId := spiffeid.RequireFromString(
		"spiffe://cluster3.demo/ns/default/sa/default/app/greeter-server",
	)

	// #region
	// Obtain secure credentials
	source, err := workloadapi.NewX509Source(ctx)
	if err != nil {
		log.Fatal(err)
	}
	creds := grpc.WithTransportCredentials(
		grpccredentials.MTLSClientCredentials(
			source, source,
			// Only allow the server that has a `serverId` SPIFFE ID.
			// If the server has a different id, the connection
			// will be rejected.
			tlsconfig.AuthorizeID(serverId),
		),
	)
	// #endregion

	client, err := grpc.DialContext(ctx, addr, creds)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	greeterClient := helloworld.NewGreeterClient(client)

	const interval = time.Second * 10
	log.Printf("Will send request every %s…\n", interval)
	for {
		sendRequest(ctx, greeterClient)
		time.Sleep(interval)
	}
}
