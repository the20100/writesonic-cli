package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/the20100/writesonic-cli/internal/api"
	"github.com/the20100/writesonic-cli/internal/config"
)

var (
	jsonFlag   bool
	prettyFlag bool
	engineFlag string
	langFlag   string
	copiesFlag int

	client *api.Client
	cfg    *config.Config
)

var rootCmd = &cobra.Command{
	Use:   "writesonic",
	Short: "Writesonic CLI — AI content generation at your fingertips",
	Long: `writesonic is a command-line interface for the Writesonic API.
Generate blog ideas, articles, landing pages, and more from your terminal.

Authentication:
  Set your API key with:  writesonic auth set-key <your-key>
  Or via environment var: WRITESONIC_API_KEY=<your-key> (or aliases: WRITESONIC_KEY, WRITESONIC_API, ...)`,
	SilenceUsage: true,
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&jsonFlag, "json", false, "Output as JSON")
	rootCmd.PersistentFlags().BoolVar(&prettyFlag, "pretty", false, "Output as pretty-printed JSON")
	rootCmd.PersistentFlags().StringVar(&engineFlag, "engine", "", "AI engine: economy, average, good, premium (default from config)")
	rootCmd.PersistentFlags().StringVar(&langFlag, "lang", "", "Language code (e.g. en, fr, de) (default from config)")
	rootCmd.PersistentFlags().IntVar(&copiesFlag, "copies", 0, "Number of copies to generate (1-5, default from config)")

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if isAuthCommand(cmd) {
			return nil
		}
		var err error
		cfg, err = config.Load()
		if err != nil {
			return fmt.Errorf("load config: %w", err)
		}

		key := resolveAPIKey()
		if key == "" {
			return fmt.Errorf("no API key found — run: writesonic auth set-key <your-key>\n" +
				"Or set the WRITESONIC_API_KEY environment variable")
		}
		client = api.NewClient(key)

		// Apply config defaults if flags not set
		if engineFlag == "" {
			if cfg.DefaultEngine != "" {
				engineFlag = cfg.DefaultEngine
			} else {
				engineFlag = "good"
			}
		}
		if langFlag == "" {
			if cfg.DefaultLanguage != "" {
				langFlag = cfg.DefaultLanguage
			} else {
				langFlag = "en"
			}
		}
		if copiesFlag == 0 {
			if cfg.DefaultCopies > 0 {
				copiesFlag = cfg.DefaultCopies
			} else {
				copiesFlag = 1
			}
		}
		return nil
	}
}

// resolveEnv returns the value of the first non-empty environment variable from the given names.
func resolveEnv(names ...string) string {
	for _, name := range names {
		if v := os.Getenv(name); v != "" {
			return v
		}
	}
	return ""
}

func resolveAPIKey() string {
	if k := resolveEnv(
		"WRITESONIC_API_KEY", "WRITESONIC_KEY", "WRITESONIC_API", "API_KEY_WRITESONIC", "API_WRITESONIC", "WRITESONIC_PK", "WRITESONIC_PUBLIC",
		"WRITESONIC_API_SECRET", "WRITESONIC_SECRET_KEY", "WRITESONIC_API_SECRET_KEY", "WRITESONIC_SECRET", "SECRET_WRITESONIC", "API_SECRET_WRITESONIC", "SK_WRITESONIC", "WRITESONIC_SK",
	); k != "" {
		return k
	}
	if cfg != nil && cfg.APIKey != "" {
		return cfg.APIKey
	}
	return ""
}

func isAuthCommand(cmd *cobra.Command) bool {
	c := cmd
	for c != nil {
		if strings.HasPrefix(c.Use, "auth") {
			return true
		}
		c = c.Parent()
	}
	return false
}
