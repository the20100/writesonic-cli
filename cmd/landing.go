package cmd

import (
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
	"github.com/vincentmaurin/writesonic-cli/internal/output"
)

var (
	landingProductName        string
	landingProductDescription string
	landingFeature1           string
	landingFeature2           string
	landingFeature3           string
	headlineProductName       string
	headlineProductDescription string
)

var landingCmd = &cobra.Command{
	Use:   "landing",
	Short: "Generate landing page copy",
}

var landingPageCmd = &cobra.Command{
	Use:   "page",
	Short: "Generate full landing page copy with features and CTAs",
	Example: `  writesonic landing page --name "Acme SaaS" --desc "Project management tool" --f1 "Task tracking" --f2 "Team collaboration" --f3 "Analytics"`,
	RunE: runLandingPage,
}

var landingHeadlineCmd = &cobra.Command{
	Use:   "headline",
	Short: "Generate catchy landing page headlines",
	Example: `  writesonic landing headline --name "Acme SaaS" --desc "Project management made simple"
  writesonic landing headline --name "ShopEasy" --desc "E-commerce platform" --copies 5`,
	RunE: runLandingHeadline,
}

func init() {
	landingPageCmd.Flags().StringVar(&landingProductName, "name", "", "Product or service name (required)")
	landingPageCmd.Flags().StringVar(&landingProductDescription, "desc", "", "Product description (required)")
	landingPageCmd.Flags().StringVar(&landingFeature1, "f1", "", "Feature 1 (required)")
	landingPageCmd.Flags().StringVar(&landingFeature2, "f2", "", "Feature 2 (required)")
	landingPageCmd.Flags().StringVar(&landingFeature3, "f3", "", "Feature 3 (required)")
	landingPageCmd.MarkFlagRequired("name")
	landingPageCmd.MarkFlagRequired("desc")
	landingPageCmd.MarkFlagRequired("f1")
	landingPageCmd.MarkFlagRequired("f2")
	landingPageCmd.MarkFlagRequired("f3")

	landingHeadlineCmd.Flags().StringVar(&headlineProductName, "name", "", "Product or service name (required)")
	landingHeadlineCmd.Flags().StringVar(&headlineProductDescription, "desc", "", "Product description (required)")
	landingHeadlineCmd.MarkFlagRequired("name")
	landingHeadlineCmd.MarkFlagRequired("desc")

	landingCmd.AddCommand(landingPageCmd, landingHeadlineCmd)
	rootCmd.AddCommand(landingCmd)
}

func runLandingPage(cmd *cobra.Command, args []string) error {
	params := url.Values{}
	params.Set("engine", engineFlag)
	params.Set("language", langFlag)
	params.Set("num_copies", fmt.Sprintf("%d", copiesFlag))

	body := map[string]interface{}{
		"product_name":        landingProductName,
		"product_description": landingProductDescription,
		"feature_1":           landingFeature1,
		"feature_2":           landingFeature2,
		"feature_3":           landingFeature3,
	}

	results, err := client.PostLandingPages(params, body)
	if err != nil {
		return err
	}

	if output.IsJSON(jsonFlag, prettyFlag) {
		return output.PrintJSON(results, prettyFlag)
	}

	for i, r := range results {
		if len(results) > 1 {
			fmt.Printf("--- Result %d ---\n\n", i+1)
		}
		output.PrintKeyValue([][]string{
			{"Title", r.Title},
			{"Subtitle", r.Subtitle},
			{"Main Feature Title", r.MainFeatureTitle},
			{"Main Feature Subtitle", r.MainFeatureSubtitle},
			{"Feature 1 Title", r.Feature1Title},
			{"Feature 1 Subtitle", r.Feature1Subtitle},
			{"Feature 2 Title", r.Feature2Title},
			{"Feature 2 Subtitle", r.Feature2Subtitle},
			{"Feature 3 Title", r.Feature3Title},
			{"Feature 3 Subtitle", r.Feature3Subtitle},
			{"CTA", r.CTA},
			{"Button", r.Button},
		})
		if i < len(results)-1 {
			fmt.Println()
		}
	}
	return nil
}

func runLandingHeadline(cmd *cobra.Command, args []string) error {
	params := url.Values{}
	params.Set("engine", engineFlag)
	params.Set("language", langFlag)
	params.Set("num_copies", fmt.Sprintf("%d", copiesFlag))

	body := map[string]interface{}{
		"product_name":        headlineProductName,
		"product_description": headlineProductDescription,
	}

	results, err := client.PostResults("/landing-page-headlines", params, body)
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
