package cmd

import (
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
	"github.com/the20100/writesonic-cli/internal/output"
)

// copy.go contains PAS, AIDA, CTA, and bullet-point-answers commands.

var (
	pasProductName        string
	pasProductDescription string
	aidaProductName       string
	aidaProductDescription string
	ctaProductName        string
	bulletQuestion        string
)

var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Generate marketing copy (PAS, AIDA, CTA, bullets)",
}

var copyPASCmd = &cobra.Command{
	Use:   "pas",
	Short: "Pain-Agitate-Solution framework copy",
	Example: `  writesonic copy pas --name "Acme" --desc "Project management software for remote teams"
  writesonic copy pas --name "FitTrack" --desc "Fitness tracking app" --copies 3`,
	RunE: runCopyPAS,
}

var copyAIDACmd = &cobra.Command{
	Use:   "aida",
	Short: "Attention-Interest-Desire-Action framework copy",
	Example: `  writesonic copy aida --name "CloudStore" --desc "Cloud storage for businesses"`,
	RunE: runCopyAIDA,
}

var copyCTACmd = &cobra.Command{
	Use:   "cta",
	Short: "Generate eye-catching calls to action",
	Example: `  writesonic copy cta --name "Writesonic"
  writesonic copy cta --name "My SaaS" --copies 5`,
	RunE: runCopyCTA,
}

var copyBulletsCmd = &cobra.Command{
	Use:   "bullets",
	Short: "Generate bullet-point answers",
	Example: `  writesonic copy bullets --question "What are the benefits of remote work?"`,
	RunE: runCopyBullets,
}

func init() {
	copyPASCmd.Flags().StringVar(&pasProductName, "name", "", "Product or service name (required)")
	copyPASCmd.Flags().StringVar(&pasProductDescription, "desc", "", "Product description (required)")
	copyPASCmd.MarkFlagRequired("name")
	copyPASCmd.MarkFlagRequired("desc")

	copyAIDACmd.Flags().StringVar(&aidaProductName, "name", "", "Product or service name (required)")
	copyAIDACmd.Flags().StringVar(&aidaProductDescription, "desc", "", "Product description (required)")
	copyAIDACmd.MarkFlagRequired("name")
	copyAIDACmd.MarkFlagRequired("desc")

	copyCTACmd.Flags().StringVar(&ctaProductName, "name", "", "Product or service name (required)")
	copyCTACmd.MarkFlagRequired("name")

	copyBulletsCmd.Flags().StringVar(&bulletQuestion, "question", "", "Question or topic to answer (required)")
	copyBulletsCmd.MarkFlagRequired("question")

	copyCmd.AddCommand(copyPASCmd, copyAIDACmd, copyCTACmd, copyBulletsCmd)
	rootCmd.AddCommand(copyCmd)
}

func runCopyPAS(cmd *cobra.Command, args []string) error {
	return postAndPrint("/pas", map[string]interface{}{
		"product_name":        pasProductName,
		"product_description": pasProductDescription,
	})
}

func runCopyAIDA(cmd *cobra.Command, args []string) error {
	return postAndPrint("/aida", map[string]interface{}{
		"product_name":        aidaProductName,
		"product_description": aidaProductDescription,
	})
}

func runCopyCTA(cmd *cobra.Command, args []string) error {
	return postAndPrint("/call-to-action", map[string]interface{}{
		"product_name": ctaProductName,
	})
}

func runCopyBullets(cmd *cobra.Command, args []string) error {
	return postAndPrint("/bulletpoint-answers", map[string]interface{}{
		"question": bulletQuestion,
	})
}

// postAndPrint is a shared helper for simple text-result endpoints.
func postAndPrint(path string, body map[string]interface{}) error {
	params := url.Values{}
	params.Set("engine", engineFlag)
	params.Set("language", langFlag)
	params.Set("num_copies", fmt.Sprintf("%d", copiesFlag))

	results, err := client.PostResults(path, params, body)
	if err != nil {
		return err
	}

	if output.IsJSON(jsonFlag, prettyFlag) {
		return output.PrintJSON(results, prettyFlag)
	}

	texts := make([]string, len(results))
	for i, r := range results {
		texts[i] = r.Text
	}
	output.PrintText(texts)
	return nil
}
