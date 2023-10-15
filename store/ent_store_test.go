package store

import (
	"context"
	"fmt"
	"github.com/bedrocksquirrel/obsmon/ent"
	"log"
	"testing"
	"time"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
)

func TestEntStore(t *testing.T) {
	// Create an ent.Client with in-memory SQLite database.
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	env1, err := client.Environment.
		Create().
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating network: %v", err)
	}
	fmt.Println(env1)

	network1, err := client.Network.
		Create().
		SetEnvironment(env1).
		SetDeployTime(time.Now()).
		SetGithubBuildNum(1).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating network: %v", err)
	}
	fmt.Println(network1)

	node1, err := client.Node.
		Create().
		SetNetwork(network1).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating network: %v", err)
	}
	fmt.Println(node1)
	// Output:
}
