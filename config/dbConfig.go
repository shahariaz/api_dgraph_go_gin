package configDB

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"

	"github.com/shahariaz/api_dgraph/model"
)

func InitDgraph() {
	//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	conn, err := grpc.NewClient(
		"localhost:9080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal("Failed to connect to Dgraph:", err)
	}

	dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	ctx := context.Background()
	schemaManager := model.NewSchemaManager(dgraphClient)
	if err := schemaManager.SetupAllSchemas(ctx); err != nil {
		log.Printf("Failed to setup schemas: %v", err)

	}

}
