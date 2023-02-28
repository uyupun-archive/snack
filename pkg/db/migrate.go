package db

import (
	"context"
	"os"
)

func Up() error {
	conn, err := Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	ctx := context.Background()
	if err := conn.Schema.Create(ctx); err != nil {
		return err
	}

	return nil
}

func Down() error {
	file, err := os.Create("database.sqlite3")
	if err != nil {
		return err
	}

	empty_data := []byte("")
	file.Write(empty_data)

	return nil
}

func Refresh() {
	Down()
	Up()
}
