package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"dagger.io/dagger"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		return err
	}
	defer client.Close()
	projectName := "test"
	database := client.Container().From("mysql:8.3").
		WithEnvVariable("MYSQL_DATABASE", projectName).
		WithEnvVariable("MYSQL_USER", projectName).
		WithEnvVariable("MYSQL_PASSWORD", projectName).
		WithEnvVariable("MYSQL_ALLOW_EMPTY_PASSWORD", "true").
		WithExposedPort(3306).
		AsService()
	database, err = database.Start(ctx)
	if err != nil {
		return err
	}
	time.Sleep(30 * time.Second)
	database, err = database.Stop(ctx)
	return err
}
