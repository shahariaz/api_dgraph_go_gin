// model/schema_manager.go
package model

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
)

type SchemaManager struct {
	client *dgo.Dgraph
}

type EntitySchema struct {
	Name   string
	Schema string
}

func NewSchemaManager(client *dgo.Dgraph) *SchemaManager {
	return &SchemaManager{
		client: client,
	}
}

func (sm *SchemaManager) GetAllEntitySchemas() []EntitySchema {
	return []EntitySchema{
		{Name: "User", Schema: UserSchema},
		{Name: "Post", Schema: PostSchema},
		{Name: "Comment", Schema: CommentSchema},
	}
}

// SetupAllSchemas sets up all schemas at once
func (sm *SchemaManager) SetupAllSchemas(ctx context.Context) error {
	schemas := sm.GetAllEntitySchemas()
	var allSchemas []string

	for _, schema := range schemas {
		allSchemas = append(allSchemas, schema.Schema)
	}

	combinedSchema := strings.Join(allSchemas, "\n")

	if err := sm.client.Alter(ctx, &api.Operation{Schema: combinedSchema}); err != nil {
		log.Printf("Error setting up all schemas: %v", err)
		return fmt.Errorf("failed to setup schemas: %w", err)
	}

	log.Println("All Dgraph schemas setup completed successfully.")
	return nil
}
