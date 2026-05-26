package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	// pq driver
	_ "github.com/lib/pq"
)

var (
	urlConn string
	path    string
)

var (
	ErrPathNotSet = errors.New("Migrations path not set. \nPlease set it using --path flag or in your prest config file")
	ErrURLNotSet  = errors.New("Database URL not set. \nPlease set it using --url flag or configure it on your prest config file")
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Execute migration operations",
	Long:  `Execute migration operations`,
}

func checkTable(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func driverURL() string { _ = "STUB: not implemented"; return "" }
