package main

import (
	"fmt"
	"log"

	storageaccrual "github.com/DeneesK/go-mart/internal/storage-accrual"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	config := parseFlags()
	_, err := storageaccrual.NewDBStorage(config.dsn)
	if err != nil {
		return fmt.Errorf("failed to start server - %w", err)
	}

	return nil
}
