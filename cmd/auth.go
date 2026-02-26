package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/the20100/writesonic-cli/internal/config"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Manage Writesonic API authentication",
}

var authSetKeyCmd = &cobra.Command{
	Use:   "set-key <api-key>",
	Short: "Save your Writesonic API key",
	Args:  cobra.ExactArgs(1),
	RunE:  runAuthSetKey,
}

var authStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show current authentication status",
	RunE:  runAuthStatus,
}

var authLogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Remove stored API key",
	RunE:  runAuthLogout,
}

var authConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Set default engine, language, and copies",
	RunE:  runAuthConfig,
}

var (
	configEngine   string
	configLanguage string
	configCopies   int
)

func init() {
	authConfigCmd.Flags().StringVar(&configEngine, "engine", "", "Default engine (economy, average, good, premium)")
	authConfigCmd.Flags().StringVar(&configLanguage, "language", "", "Default language code (e.g. en, fr, de)")
	authConfigCmd.Flags().IntVar(&configCopies, "copies", 0, "Default number of copies (1-5)")

	authCmd.AddCommand(authSetKeyCmd, authStatusCmd, authLogoutCmd, authConfigCmd)
	rootCmd.AddCommand(authCmd)
}

func runAuthSetKey(cmd *cobra.Command, args []string) error {
	key := args[0]

	cfg, err := config.Load()
	if err != nil {
		cfg = &config.Config{}
	}
	cfg.APIKey = key
	if err := cfg.Save(); err != nil {
		return fmt.Errorf("save config: %w", err)
	}

	path, _ := config.ConfigPath()
	fmt.Printf("API key saved to %s\n", path)
	fmt.Println("You can now run: writesonic blog-ideas --topic \"AI in 2025\"")
	return nil
}

func runAuthStatus(cmd *cobra.Command, args []string) error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	path, _ := config.ConfigPath()
	fmt.Printf("Config file:      %s\n", path)

	key := resolveAPIKey()
	if key == "" {
		fmt.Println("API key:          not set")
		fmt.Println("\nRun: writesonic auth set-key <your-key>")
		fmt.Println("Or:  export WRITESONIC_API_KEY=<your-key>")
		return nil
	}

	masked := maskKey(key)
	source := "environment variable"
	if cfg.APIKey != "" && key == cfg.APIKey {
		source = "config file"
	}
	fmt.Printf("API key:          %s (%s)\n", masked, source)

	engine := cfg.DefaultEngine
	if engine == "" {
		engine = "good (default)"
	}
	language := cfg.DefaultLanguage
	if language == "" {
		language = "en (default)"
	}
	copies := cfg.DefaultCopies
	fmt.Printf("Default engine:   %s\n", engine)
	fmt.Printf("Default language: %s\n", language)
	if copies == 0 {
		fmt.Printf("Default copies:   1 (default)\n")
	} else {
		fmt.Printf("Default copies:   %d\n", copies)
	}
	return nil
}

func runAuthLogout(cmd *cobra.Command, args []string) error {
	if err := config.Clear(); err != nil {
		return fmt.Errorf("clear config: %w", err)
	}
	fmt.Println("API key removed. Set a new key with: writesonic auth set-key <key>")
	return nil
}

func runAuthConfig(cmd *cobra.Command, args []string) error {
	cfg, err := config.Load()
	if err != nil {
		cfg = &config.Config{}
	}

	changed := false
	if configEngine != "" {
		cfg.DefaultEngine = configEngine
		changed = true
	}
	if configLanguage != "" {
		cfg.DefaultLanguage = configLanguage
		changed = true
	}
	if configCopies > 0 {
		cfg.DefaultCopies = configCopies
		changed = true
	}

	if !changed {
		fmt.Println("No changes. Use --engine, --language, or --copies flags.")
		fmt.Println("Example: writesonic auth config --engine premium --language fr --copies 3")
		return nil
	}

	if err := cfg.Save(); err != nil {
		return fmt.Errorf("save config: %w", err)
	}
	fmt.Println("Defaults updated.")
	return nil
}

func maskKey(key string) string {
	if len(key) <= 8 {
		return "****"
	}
	return key[:4] + "..." + key[len(key)-4:]
}
