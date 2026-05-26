package cmd

import (
	"github.com/prest/prest/v2/adapters/postgres"
	"github.com/prest/prest/v2/config"

	"log/slog"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "prestd",
	Short: "Serve a RESTful API from any PostgreSQL database",
	Long:  `prestd (PostgreSQL REST), simplify and accelerate development, ⚡ instant, realtime, high-performance on any Postgres application, existing or new`,
	Run: func(cmd *cobra.Command, args []string) {
		if config.PrestConf.Adapter == nil {
			slog.Warn("adapter is not set. Using the default (postgres)")
			postgres.Load()
		}
		startServer()
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() { _ = "STUB: not implemented"; return }

// startServer starts the server
func startServer() {
	_ = "STUB: not implemented"
	// Fail fast when JWT enforcement is enabled but no verification material
	// was provided — otherwise the middleware would validate bearer tokens
	// against an empty HMAC key. Subcommands like `migrate` don't reach here,
	// so they keep working without JWT material configured.
	// See GHSA-fj7v-859r-2fm4.
	return
}
