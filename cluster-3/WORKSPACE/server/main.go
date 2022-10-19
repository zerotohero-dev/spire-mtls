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
	"fmt"
	"github.com/spiffe/go-spiffe/v2/spiffegrpc/grpccredentials"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/examples/helloworld/helloworld"
	"log"
	"net"
)

type greeter struct {
	helloworld.UnimplementedGreeterServer
}

func (greeter) SayHello(ctx context.Context, req *helloworld.HelloRequest) (
	*helloworld.HelloReply, error,
) {
	clientId := "UnknownClient"

	// #region
	// Learn client’s id from its SPIRE SVID.
	if peerId, ok := grpccredentials.PeerIDFromContext(ctx); ok {
		clientId = peerId.String()
	}
	// #endregion

	log.Printf("%s has requested that I say hello to %q…", clientId, req.Name)

	return &helloworld.HelloReply{
		Message: fmt.Sprintf("On behalf of %s, hello %s.", clientId, req.Name),
	}, nil
}

func main() {
	var addr string
	flag.StringVar(&addr, "addr", "localhost:8123", "host:port of the server")
	flag.Parse()

	log.Printf("Server (%s) starting up…", addr)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	// #region
	// Get the source from the workload API.
	// Source contains both the workload’s SVID and also trust bundles.
	source, err := workloadapi.NewX509Source(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer source.Close()
	// #endregion

	// SVID cryptographically identifies the server workload.
	svid, _ := source.GetX509SVID()
	log.Println("Got the source, Luke:", svid.ID)

	// #region
	// Create secure credentials.
	//
	// We are using the X.509 certs from the source
	// to create the credentials.
	//
	// The first `source` in the arguments is the source
	// of the SVID.
	// The second `source` in the arguments is the source
	// of the trust bundle.
	// In our case, the two sources are the same.
	//
	// Note that, we are using `tlsconfig.AuthorizeAny()`,
	// which means, as long as the certificates in the
	// creds are valid and trusted, we’ll allow the connection.
	//
	// Instead of “authorize any”, we can use different
	// configuration that, say, leverages a policy engine
	// such as OPA.
	creds := grpc.Creds(
		grpccredentials.MTLSServerCredentials(
			source, source, tlsconfig.AuthorizeAny(),
		),
	)
	// #enregion

	server := grpc.NewServer(creds)
	helloworld.RegisterGreeterServer(server, greeter{})

	log.Println("Serving on", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
