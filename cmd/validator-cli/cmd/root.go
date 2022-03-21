package cmd

import (
	"os"
	"time"

	"github.com/arttet/validator-service/internal/app/validator-service/entrypoint"

	"github.com/spf13/cobra"

	_ "github.com/go-sql-driver/mysql"
)

var (
	// DSN (Data Source Name).
	dsn string
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "validator-cli",
	Short: "The command-line tool, validator-cli, allows you to perform service availability checks",
	Long:  "",
	Run:   run,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&dsn, "dsn", "", "", "the data source name for connecting to the MySQL database (required)")
	rootCmd.MarkFlagRequired("dsn") // nolint:errcheck
}

func run(cmd *cobra.Command, args []string) {
	entrypoint.EntryPoint(dsn, 30*time.Second)
}
