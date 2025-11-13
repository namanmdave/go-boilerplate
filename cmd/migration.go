package cmd

import (
	"fmt"
	"strings"

	"go-boilerplate/config"
	"go-boilerplate/store"
	"go-boilerplate/store/schema"

	"github.com/spf13/cobra"
)

// migrationCmd represents the migration command
var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "Manage database migrations",
	Long:  `Manage database migrations (up, down, status, fresh, etc.)`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.Init()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			fmt.Println("Please specify a migration action: up, down, status, fresh")
			return nil
		}

		action := args[0]
		fmt.Printf("Running migration: %s\n", action)

		switch action {
		case "up":
			if err := migrateUp(); err != nil {
				return err
			}
		default:
			fmt.Printf("Unknown migration action: %s\n", action)
		}

		return nil
	},
}

func init() {
	AddCommand(migrationCmd)
}

// migrateUp executes statements from the embedded schema.Schema string.
func migrateUp() error {
	raw := schema.Schema
	parts := strings.Split(raw, ";")

	db, _ := store.InitPostgres(config.GetDBConfig())
	if db == nil {
		return fmt.Errorf("config.DB is nil; please expose *sql.DB in config package as DB (set by config.Init)")
	}

	for _, p := range parts {
		stmt := strings.TrimSpace(p)
		if stmt == "" {
			continue
		}
		if _, err := db.Exec(stmt); err != nil {
			return fmt.Errorf("exec statement failed: %w\nstatement: %s", err, stmt)
		}
	}

	fmt.Println("Migrations applied from embedded schema")
	return nil
}
